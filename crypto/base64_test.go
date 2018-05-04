package crypto

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {

	_encodeBase64 := EncodeBase64("go语言中，判断两个字符串是否相等")
	fmt.Println(_encodeBase64)

	_decodeBase64, _ := DecodeBase64(_encodeBase64)
	fmt.Println(_decodeBase64)

	if _decodeBase64 != "go语言中，判断两个字符串是否相等" {
		t.Errorf("Base64 error")
	}
}
