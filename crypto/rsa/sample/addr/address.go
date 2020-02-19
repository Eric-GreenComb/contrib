package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// GenJavaAddr()

	_publicKey := hex.EncodeToString(publicKey)

	fmt.Println(GenAddr("cebbank", _publicKey))
}

// GenJavaAddr GenJavaAddr
func GenJavaAddr() {
	fHash := sha256.Sum256([]byte("工作计划"))
	lHash := sha256.Sum256(fHash[:])
	strHash := hex.EncodeToString(lHash[:])
	fmt.Println(strHash)
}

// GenAddr GenAddr
func GenAddr(ID, publicKey string) string {
	fHash := sha256.Sum256([]byte(publicKey))
	lHash := sha256.Sum256(fHash[:])
	strHash := hex.EncodeToString(lHash[:])
	return fmt.Sprintf("%s:%s", ID, strHash)
}

var publicKey []byte

func init() {

	var err error
	publicKey, err = ioutil.ReadFile("public.pem")
	if err != nil {
		os.Exit(-1)
	}
}
