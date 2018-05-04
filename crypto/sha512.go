package crypto

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

// SHA512Byte sha512 byte in/out
func SHA512Byte(in []byte) []byte {
	hash := sha512.New()
	hash.Write(in)
	md := hash.Sum(nil)
	return md
}

// SHA512Base64 sha512 in string/out base64
func SHA512Base64(in string) string {
	hash := sha512.New()
	hash.Write([]byte(in))
	md := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(md)
}

// SHA512Hex sha512 in string /out hex
func SHA512Hex(in string) string {
	hash := sha512.New()
	hash.Write([]byte(in))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

// CompareSHA512Hex compare sha512 value
func CompareSHA512Hex(in, sha512Hex string) bool {
	if SHA512Hex(in) == sha512Hex {
		return true
	}
	return false
}
