# multicall, inspired by [@indexed-finance/multicall](https://github.com/indexed-finance/multicall)

## Usage

```golang
Do(ctx context.Context, client *ethclient.Client, ab *abi.ABI, invokes []Invoke, result interface{}) (height uint64, err error)
```

`invokes` is the solidity calls you want to make;

`result` is the result slice, its size should be the same as `invokes`, and the type of the ith result should match the output type of the ith invoke. 

The `result` is updated in-place, and the `height` of the call is also returned.

Follow example [here](https://github.com/zhiqiangxu/multicall/blob/master/multicall_test.go#L37).