package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	myrsa "github.com/Eric-GreenComb/contrib/crypto/rsa"
)

func main() {

	initData := "abcdefghij_真是klmnopq_"
	init := []byte(initData)

	data, err := myrsa.Encrypt(init, publicKey)
	if err != nil {
		panic(err)
	}
	pre := time.Now()
	origData, err := myrsa.Decrypt(data, privateKey)
	if err != nil {
		panic(err)
	}
	now := time.Now()
	fmt.Println(now.Sub(pre))
	fmt.Println(string(origData))
}

// RsaEncrypt 公钥加密
func RsaEncrypt(origData, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RsaDecrypt 私钥解密
func RsaDecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

var privateKey, publicKey []byte

func init() {
	var err error
	// flag.StringVar(&decrypted, "d", "", "加密过的数据")
	// flag.Parse()
	publicKey, err = ioutil.ReadFile("public.pem")
	if err != nil {
		os.Exit(-1)
	}
	privateKey, err = ioutil.ReadFile("private.pem")
	if err != nil {
		os.Exit(-1)
	}
}
