package parse

import (
	"errors"
	"math/big"
)

// ERC20Transfer ParseTokenTransfer
// 0xa9059cbb0000000000000000000000000cf0698955123303a9a36ce470552c8d10ee6198000000000000000000000000000000000000000000000001158e460913d00000
// 0xa9059cbb000000000000000000000000af6bcac3d8cc3d93560ce8a4d5ab1a2c3bf0436100000000000000000000000000000000000000000000000006f05b59d3b20000
func ERC20Transfer(inputData string) (string, int64, error) {

	if len(inputData) != 138 {
		return "", 0, errors.New("len is error")
	}

	if inputData[0:10] != "0xa9059cbb" {
		return "", 0, errors.New("MethodID is error")
	}

	_addr := "0x" + inputData[34:74]

	n := new(big.Int)
	m := new(big.Int)
	k := big.NewInt(1000000000000)
	n, ok := n.SetString(inputData[74:138], 16)
	if !ok {
		return "", 0, errors.New("SetString is error")
	}

	return _addr, m.Div(n, k).Int64(), nil
}
