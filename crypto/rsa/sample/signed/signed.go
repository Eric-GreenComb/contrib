package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
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

	err = myrsa.Verify(crypto.SHA256, hashed[:], pub, _sign)
	// err = RsaVerify(crypto.MD5, hashed[:], pub, _sign)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sign is right")
	}

}

// RsaSign rsa sign
// hashed := md5.Sum(src)
// hashed := sha256.Sum256(src)
func RsaSign(hash crypto.Hash, hashed, privateKey []byte) ([]byte, error) {
	// Only small messages can be signed directly; thus the hash of a
	// message, rather than the message itself, is signed. This requires
	// that the hash function be collision resistant. SHA-256 is the
	// least-strong hash function that should be used for this at the time
	// of writing (2016).

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.SignPKCS1v15(rand.Reader, priv, hash, hashed)
	// return rsa.SignPKCS1v15(rand.Reader, priv, crypto.MD5, hashed[:])
	// return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed[:])
}

// RsaVerify rsa verify
// hashed := md5.Sum(src)
// hashed := sha256.Sum256(src)
// signature is a valid signature of message from the public key.
func RsaVerify(hash crypto.Hash, hashed, publicKey, signature []byte) error {

	// Only small messages can be signed directly; thus the hash of a
	// message, rather than the message itself, is signed. This requires
	// that the hash function be collision resistant. SHA-256 is the
	// least-strong hash function that should be used for this at the time
	// of writing (2016).

	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)

	return rsa.VerifyPKCS1v15(pub, hash, hashed, signature)
	// return rsa.VerifyPKCS1v15(pub, crypto.MD5, hashed[:], signature)
	// return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
}

var pub, key []byte

func init() {
	var err error
	pub, err = ioutil.ReadFile("public.pem")
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
