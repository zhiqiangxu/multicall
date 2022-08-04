package multicall

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"reflect"
	"strings"

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

	resultBytes, err := client.CallContract(ctx, ethereum.CallMsg{Data: append(common.FromHex(mcabi.MultiCallBin), packed...)}, nil)
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

func DoSlice(ctx context.Context, client *ethclient.Client, ab *abi.ABI, total, unit int, invokeFunc func(i int) []Invoke, result interface{}) (height uint64, err error) {
	if total <= 0 {
		return
	}
	if unit <= 0 {
		panic("unit <= 0")
	}
	s := reflect.ValueOf(result)
	if total != s.Len() {
		err = fmt.Errorf("total != #results")
		return
	}

	height = uint64(math.MaxUint64)
	for from := 0; from < total; from += unit {
		to := from + unit
		if to > total {
			to = total
		}
		invokes := make([]Invoke, 0, to-from)
		for k := from; k < to; k++ {
			invokes = append(invokes, invokeFunc(k)...)
		}
		var unitHeight uint64
		unitHeight, err = Do(ctx, client, ab, invokes, s.Slice(from, to).Interface())
		if err != nil {
			return
		}
		if height == math.MaxUint64 {
			height = unitHeight
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
