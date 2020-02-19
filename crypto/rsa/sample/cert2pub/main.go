package main

import (
	"crypto/rsa"
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

var cert, key []byte

func init() {
	var err error
	cert, err = ioutil.ReadFile("cert.pem")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
