package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(GenAddr("工作计划"))
}

func GenAddr(plainText string) string {
	fHash := sha256.Sum256([]byte(plainText))
	lHash := sha256.Sum256(fHash[:])
	strHash := hex.EncodeToString(lHash[:])
	return strHash
}
