package crypto

import (
	"fmt"
	"testing"
)

func TestHS512(t *testing.T) {

	_encodeBase64 := HS512Base64("go语言中，判断两个字符串是否相等", "1")
	fmt.Println("HS512Base64: " + _encodeBase64)

	_encodeHex := HS512Hex("go语言中，判断两个字符串是否相等", "1")
	fmt.Println("HS512Hex: " + _encodeHex)
}
