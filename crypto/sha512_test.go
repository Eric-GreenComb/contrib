package crypto

import (
	"fmt"
	"testing"
)

func TestSHA512(t *testing.T) {

	_encodeBase64 := SHA512Base64("go语言中，判断两个字符串是否相等")
	fmt.Println("SHA256Base64: " + _encodeBase64)

	_encodeHex := SHA512Hex("go语言中，判断两个字符串是否相等")
	fmt.Println("SHA256Hex: " + _encodeHex)
}
