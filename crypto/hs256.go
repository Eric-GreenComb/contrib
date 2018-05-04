package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// HS256Byte hs256 byte in/out
func HS256Byte(in, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write(in)
	return h.Sum(nil)
}

// HS256Base64 hs256 in string/out base64
func HS256Base64(in, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(in))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// HS256Hex hs256 in string /out hex
func HS256Hex(in, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(in))
	return hex.EncodeToString(h.Sum(nil))
}

// CompareHS256Hex compare HS256Hex value
func CompareHS256Hex(in, secret, hs256Hex string) bool {
	if HS256Hex(in, secret) == hs256Hex {
		return true
	}
	return false
}
