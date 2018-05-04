package crypto

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// SHA256Byte sha256 byte in/out
func SHA256Byte(in []byte) []byte {
	hash := sha256.New()
	hash.Write(in)
	md := hash.Sum(nil)
	return md
}

// SHA256Base64 sha256 in string/out base64
func SHA256Base64(in string) string {
	hash := sha256.New()
	hash.Write([]byte(in))
	md := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(md)
}

// SHA256Hex sha256 in string /out hex
func SHA256Hex(in string) string {
	hash := sha256.New()
	hash.Write([]byte(in))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

// CompareSHA256Hex compare sha256 value
func CompareSHA256Hex(in, sha256Hex string) bool {
	if SHA256Hex(in) == sha256Hex {
		return true
	}
	return false
}
