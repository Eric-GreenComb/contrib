package crypto

import (
	"fmt"
	"testing"
)

func TestSHA256(t *testing.T) {

	_encodeBase64 := SHA256Base64("go语言中，判断两个字符串是否相等")
	fmt.Println("SHA256Base64: " + _encodeBase64)

	_encodeHex := SHA256Hex("go语言中，判断两个字符串是否相等")
	fmt.Println("SHA256Hex: " + _encodeHex)
}
