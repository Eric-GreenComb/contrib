package crypto

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

// HS512Byte hs512 byte in/out
func HS512Byte(in, secret []byte) []byte {
	h := hmac.New(sha512.New, secret)
	h.Write(in)
	return h.Sum(nil)
}

// HS512Base64 hs512 in string/out base64
func HS512Base64(in, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(in))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// HS512Hex hs512 in string /out hex
func HS512Hex(in, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(in))
	return hex.EncodeToString(h.Sum(nil))
}

// CompareHS512Hex compare HS512Hex value
func CompareHS512Hex(in, secret, hs512Hex string) bool {
	if HS512Hex(in, secret) == hs512Hex {
		return true
	}
	return false
}
