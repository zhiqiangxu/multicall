package multicall

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"reflect"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	mcabi "github.com/zhiqiangxu/multicall/contracts/abi"
)

type Invoke struct {
	Contract common.Address
	Name     string
	Args     []interface{}
	AB       *abi.ABI
}

var (
	uint256Type    abi.Type
	bytesSliceType abi.Type
	parsedAbi      abi.ABI
)

func init() {
	var err error
	uint256Type, err = abi.NewType("uint256", "", nil)
	if err != nil {
		log.Fatalf("addressSliceType failed:%v", err)
	}
	bytesSliceType, err = abi.NewType("bytes[]", "", nil)
	if err != nil {
		log.Fatalf("bytesSliceType failed:%v", err)
	}
	parsedAbi, _ = abi.JSON(strings.NewReader(mcabi.MultiCallABI))
}

var Retry = 3
var BackoffInterval = time.Second * 2

func Do(ctx context.Context, client *ethclient.Client, ab *abi.ABI, invokes []Invoke, result interface{}) (height uint64, err error) {

	results := InterfaceSlice(result)
	if len(invokes) != len(results) {
		err = fmt.Errorf("#invokes != #results")
		return
	}

	var (
		targets []common.Address
		inputs  [][]byte
	)
	for _, invoke := range invokes {
		invokeAB := invoke.AB
		if invokeAB == nil {
			invokeAB = ab
		}
		method, exist := invokeAB.Methods[invoke.Name]
		if !exist {
			err = fmt.Errorf("method '%s' not found", invoke.Name)
			return
		}

		var arguments []byte
		arguments, err = method.Inputs.Pack(invoke.Args...)
		if err != nil {
			return
		}

		targets = append(targets, invoke.Contract)
		inputs = append(inputs, append(method.ID, arguments...))
	}

	packed, err := parsedAbi.Pack("", targets, inputs)
	if err != nil {
		return
	}

	var resultBytes []byte
	for i := 0; i < Retry; i++ {
		resultBytes, err = client.CallContract(ctx, ethereum.CallMsg{Data: append(common.FromHex(mcabi.MultiCallBin), packed...)}, nil)
		if err != nil {
			time.Sleep(BackoffInterval)
			continue
		}
		break
	}
	if err != nil {
		return
	}

	arguments := abi.Arguments{
		{Type: uint256Type, Name: "Height"},
		{Type: bytesSliceType, Name: "ReturnDatas"},
	}

	var output struct {
		Height      *big.Int
		ReturnDatas [][]byte
	}
	resultInterface, err := arguments.Unpack(resultBytes)
	if err != nil {
		return
	}
	err = arguments.Copy(&output, resultInterface)
	if err != nil {
		return
	}

	if len(output.ReturnDatas) != len(invokes) {
		err = fmt.Errorf("#ReturnDatas != #invokes")
		return
	}
	for i, returnData := range output.ReturnDatas {
		if len(returnData) == 0 {
			continue
		}

		invoke := invokes[i]
		invokeAB := invoke.AB
		if invokeAB == nil {
			invokeAB = ab
		}

		method := invokeAB.Methods[invoke.Name]
		var returnValue []interface{}
		returnValue, err = method.Outputs.Unpack(returnData)
		if err != nil {
			return
		}

		err = method.Outputs.Copy(results[i], returnValue)
		if err != nil {
			return
		}
	}

	height = output.Height.Uint64()
	return
}

func DoSlice(ctx context.Context, client *ethclient.Client, ab *abi.ABI, total, unit int, invokeFunc func(i int) []Invoke, beforeDo func(from, to int), result interface{}) (height uint64, err error) {
	if total <= 0 {
		return
	}
	if unit <= 0 {
		panic("unit <= 0")
	}
	s := reflect.ValueOf(result)

	height = uint64(math.MaxUint64)
	invokes := make([]Invoke, 0, unit)
	nextFrom := 0
	for from := 0; from < total; from += unit {
		to := from + unit
		if to > total {
			to = total
		}
		invokes = invokes[:0]
		for k := from; k < to; k++ {
			invokes = append(invokes, invokeFunc(k)...)
		}
		var unitHeight uint64
		if beforeDo != nil {
			beforeDo(from, to)
		}
		unitHeight, err = Do(ctx, client, ab, invokes, s.Slice(nextFrom, nextFrom+len(invokes)).Interface())
		if err != nil {
			return
		}
		nextFrom += len(invokes)

		if height == math.MaxUint64 {
			height = unitHeight
		}
	}
	return
}

func DoSliceCvt[T any](ctx context.Context, client *ethclient.Client, ab *abi.ABI, total, unit int, invokeFunc func(i int) []Invoke, cvtFunc func(from, to int, result []T) error, beforeDo func(from, to int)) (height uint64, err error) {
	if total <= 0 {
		return
	}
	if unit <= 0 {
		panic("unit <= 0")
	}
	height = uint64(math.MaxUint64)

	buffer := make([]T, unit)
	invokes := make([]Invoke, 0, unit)
	for from := 0; from < total; from += unit {
		to := from + unit
		if to > total {
			to = total
		}
		invokes = invokes[:0]
		for k := from; k < to; k++ {
			invokes = append(invokes, invokeFunc(k)...)
		}
		if len(invokes) > len(buffer) {
			buffer = make([]T, len(invokes))
		}
		var unitHeight uint64
		if beforeDo != nil {
			beforeDo(from, to)
		}
		unitHeight, err = Do(ctx, client, ab, invokes, buffer[0:len(invokes)])
		if err != nil {
			return
		}
		if height == math.MaxUint64 {
			height = unitHeight
		}

		err = cvtFunc(from, to, buffer[0:len(invokes)])
		if err != nil {
			return
		}
	}
	return
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Addr().Interface()
	}

	return ret
}
