package multicall

import (
	"context"
	"reflect"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zhiqiangxu/util"
)

func DoSliceConcurrent(ctx context.Context, clients []*ethclient.Client, ab *abi.ABI, total, unit int, invokeFunc func(i int) []Invoke, beforeDo func(from, to int), result interface{}) (height uint64, err error) {
	if total <= 0 {
		return
	}
	if unit <= 0 {
		panic("unit <= 0")
	}
	s := reflect.ValueOf(result)

	concurrentUnit := len(clients) * unit
	invokes := make([]Invoke, 0, concurrentUnit)

	heightCh := make(chan uint64, 1)
	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	for from := 0; from < total; from += concurrentUnit {
		to := from + concurrentUnit
		if to > total {
			to = total
		}
		if beforeDo != nil {
			beforeDo(from, to)
		}

		invokes = invokes[:0]
		for k := from; k < to; k++ {
			invokes = append(invokes, invokeFunc(k)...)
		}

		for i := 0; i*unit < len(invokes); i++ {
			to := (i + 1) * unit
			if to > len(invokes) {
				to = len(invokes)
			}
			subInvokes := invokes[i*unit : to]
			client := clients[i]
			sliceFrom := from + i*unit
			sliceTo := from + to
			util.GoFunc(&wg, func() {
				unitHeight, err := Do(ctx, client, ab, subInvokes, s.Slice(sliceFrom, sliceTo).Interface())
				if err != nil {
					select {
					case errCh <- err:
					default:
					}
					return
				}
				select {
				case heightCh <- unitHeight:
				default:
				}
			})
		}
		wg.Wait()

		select {
		case err = <-errCh:
			return
		default:
		}
	}

	height = <-heightCh
	return
}

func DoSliceCvtConcurrent[T any](ctx context.Context, clients []*ethclient.Client, ab *abi.ABI, total, unit int, invokeFunc func(i int) []Invoke, cvtFunc func(from, to int, result []T) error, beforeDo func(from, to int)) (height uint64, err error) {
	if total <= 0 {
		return
	}
	if unit <= 0 {
		panic("unit <= 0")
	}

	concurrentUnit := len(clients) * unit
	invokes := make([]Invoke, 0, concurrentUnit)
	buffer := make([]T, concurrentUnit)

	heightCh := make(chan uint64, 1)
	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	for from := 0; from < total; from += unit {
		to := from + unit
		if to > total {
			to = total
		}
		if beforeDo != nil {
			beforeDo(from, to)
		}

		invokes = invokes[:0]
		for k := from; k < to; k++ {
			invokes = append(invokes, invokeFunc(k)...)
		}

		if len(buffer) < len(invokes) {
			buffer = make([]T, len(invokes))
		}
		buffer = buffer[0:len(invokes)]

		for i := 0; i*unit < len(invokes); i++ {
			to := (i + 1) * unit
			if to > len(invokes) {
				to = len(invokes)
			}
			subInvokes := invokes[i*unit : to]
			client := clients[i]
			sliceFrom := i * unit
			sliceTo := to
			util.GoFunc(&wg, func() {
				unitHeight, err := Do(ctx, client, ab, subInvokes, buffer[sliceFrom:sliceTo])
				if err != nil {
					select {
					case errCh <- err:
					default:
					}
					return
				}
				select {
				case heightCh <- unitHeight:
				default:
				}
			})
		}
		wg.Wait()

		err = cvtFunc(from, to, buffer[0:len(invokes)])
		if err != nil {
			return
		}

		select {
		case err = <-errCh:
			return
		default:
		}
	}

	height = <-heightCh
	return
}
