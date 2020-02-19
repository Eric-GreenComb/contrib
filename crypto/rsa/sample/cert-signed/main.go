package main

import (
	"crypto"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	myrsa "github.com/Eric-GreenComb/contrib/crypto/rsa"
)

func main() {

	hexPublicKey := hex.EncodeToString([]byte("gopper.in"))
	fmt.Println(hexPublicKey)

	initData := "pub" + hexPublicKey
	init := []byte(initData)
	hashed := sha256.Sum256([]byte(init))
	// hashed := md5.Sum([]byte(init))

	_sign, err := myrsa.Sign(crypto.SHA256, hashed[:], key)
	// _sign, err := RsaSign(crypto.MD5, hashed[:], key)

	fmt.Println(hex.EncodeToString(_sign))

	err = myrsa.CertVerify(crypto.SHA256, hashed[:], cert, _sign)
	// err = RsaVerify(crypto.MD5, hashed[:], cert, _sign)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Cert Sign is right")
	}

}

var cert, key []byte

func init() {
	var err error
	cert, err = ioutil.ReadFile("cert.pem")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	key, err = ioutil.ReadFile("private.pem")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
