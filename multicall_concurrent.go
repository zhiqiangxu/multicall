package multicall

import (
	"context"
	"reflect"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DoSliceConcurrent(
	ctx context.Context,
	clients []*ethclient.Client,
	ab *abi.ABI,
	total, unit int,
	invokeFunc func(i int) []Invoke,
	beforeDo func(from, to int),
	onErr func(subInvokes []Invoke, err error, client *ethclient.Client),
	result interface{},
	fromAddr ...common.Address) (height uint64, err error) {
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

	nextFrom := 0
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

		invokeUnit := len(invokes)/len(clients) + 1
		for i := 0; i*invokeUnit < len(invokes); i++ {
			to := (i + 1) * invokeUnit
			if to > len(invokes) {
				to = len(invokes)
			}
			subInvokes := invokes[i*invokeUnit : to]
			client := clients[i]
			sliceFrom := nextFrom + i*invokeUnit
			sliceTo := nextFrom + to
			goFunc(&wg, func() {
				unitHeight, err := Do(ctx, client, ab, subInvokes, s.Slice(sliceFrom, sliceTo).Interface(), fromAddr...)
				if err != nil {
					if onErr != nil {
						onErr(subInvokes, err, client)
					}

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
		nextFrom += len(invokes)

		if onErr != nil {
			continue
		}
		select {
		case err = <-errCh:
			return
		default:
		}
	}

	if onErr != nil {
		select {
		case err = <-errCh:
			return
		default:
		}
	}
	height = <-heightCh
	return
}

func DoSliceCvtConcurrent[T any](
	ctx context.Context,
	clients []*ethclient.Client,
	ab *abi.ABI,
	total, unit int,
	invokeFunc func(i int) []Invoke,
	cvtFunc func(from, to int, result []T) error,
	beforeDo func(from, to int),
	onErr func(subInvokes []Invoke, err error, client *ethclient.Client),
	fromAddr ...common.Address) (height uint64, err error) {
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

		if len(buffer) < len(invokes) {
			buffer = make([]T, len(invokes))
		}
		buffer = buffer[0:len(invokes)]

		invokeUnit := len(invokes)/len(clients) + 1
		for i := 0; i*invokeUnit < len(invokes); i++ {
			to := (i + 1) * invokeUnit
			if to > len(invokes) {
				to = len(invokes)
			}
			subInvokes := invokes[i*invokeUnit : to]
			client := clients[i]
			sliceFrom := i * invokeUnit
			sliceTo := to
			goFunc(&wg, func() {
				unitHeight, err := Do(ctx, client, ab, subInvokes, buffer[sliceFrom:sliceTo], fromAddr...)
				if err != nil {
					if onErr != nil {
						onErr(subInvokes, err, client)
					}
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

		if onErr != nil {
			continue
		}

		select {
		case err = <-errCh:
			return
		default:
		}
	}

	if onErr != nil {
		select {
		case err = <-errCh:
			return
		default:
		}
	}
	height = <-heightCh
	return
}

// goFunc runs a goroutine under WaitGroup
func goFunc(routinesGroup *sync.WaitGroup, f func()) {
	routinesGroup.Add(1)
	go func() {
		defer routinesGroup.Done()
		f()
	}()
}
