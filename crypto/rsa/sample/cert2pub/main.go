package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	block, _ := pem.Decode(cert)
	var cert *x509.Certificate
	cert, _ = x509.ParseCertificate(block.Bytes)

	fmt.Println(cert.Subject)
	fmt.Println(cert.SerialNumber)
	fmt.Println(cert.SignatureAlgorithm)

	// h := sha256.New()
	// h.Write(cert.RawTBSCertificate)
	// _hash := h.Sum(nil)
	_hashed := sha256.Sum256(cert.RawTBSCertificate)

	_block, _ := pem.Decode(ca)
	if _block == nil {
		fmt.Println("public key error")
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(_block.Bytes)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_pub := pubInterface.(*rsa.PublicKey)

	err = rsa.VerifyPKCS1v15(_pub, crypto.SHA256, _hashed[:], cert.Signature)
	if err != nil {
		println("Signature does not match")
	} else {
		println("Signature matched")
	}

	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)
	derPkix, err := x509.MarshalPKIXPublicKey(rsaPublicKey)
	if err != nil {
		return
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err := os.Create("public.pem")
	if err != nil {
		return
	}
	err = pem.Encode(file, block)
	if err != nil {
		return
	}
}

var cert, ca []byte

func init() {
	var err error
	cert, err = ioutil.ReadFile("cert.pem")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	ca, err = ioutil.ReadFile("ca_public.pem")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

}
