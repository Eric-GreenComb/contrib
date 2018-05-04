package main

import (
	"crypto"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// RsaGenKey(1024)

	initData := "用于签名的原始数据"
	hashed := sha256.Sum256([]byte(initData))
	_sign, err := RsaSign(crypto.SHA256, hashed[:], private)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))

	err = RsaVerify(crypto.SHA256, hashed[:], public, _sign)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sign is right")
	}

}

var public, private []byte

func init() {
	var err error

	public, err = ioutil.ReadFile("public.pem")
	if err != nil {
		os.Exit(-1)
	}
	private, err = ioutil.ReadFile("private.pem")
	if err != nil {
		os.Exit(-1)
	}
}
