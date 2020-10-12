package crypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestDesEncrypt(t *testing.T) {

	_des, err := DesEncrypt([]byte("go语言中，判断两个字符串是否相等"), []byte("11111111"))
	if err != nil {
		fmt.Println(err.Error())
	}

	_encodeBase64 := base64.StdEncoding.EncodeToString(_des)
	fmt.Println(_encodeBase64)

}

func TestTripleDesEncrypt(t *testing.T) {

	_des3, err := TripleDesEncrypt([]byte("go语言中，判断两个字符串是否相等"), []byte("111111111111111111111111"))
	if err != nil {
		fmt.Println(err.Error())
	}

	_encodeBase64 := base64.StdEncoding.EncodeToString(_des3)
	fmt.Println(_encodeBase64)

}
