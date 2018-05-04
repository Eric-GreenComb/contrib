package crypto

import (
	"fmt"
	"testing"
)

func TestHS256(t *testing.T) {

	_encodeBase64 := HS256Base64("go语言中，判断两个字符串是否相等", "相等")
	fmt.Println("HS256Base64: " + _encodeBase64)

	_encodeHex := HS256Hex("go语言中，判断两个字符串是否相等", "相等")
	fmt.Println("HS256Hex: " + _encodeHex)
}
