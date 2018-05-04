package crypto

import (
	"encoding/base64"
)

// Base64Encode base64 encode 2 []byte
func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

// Base64Decode base64 decode 2 []byte
func Base64Decode(dec []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(dec))
}

// EncodeBase64 base64 encode 2 string
func EncodeBase64(src string) string {
	if len(src) == 0 {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(src))
}

// DecodeBase64 base64 decode 2 string
func DecodeBase64(dec string) (string, error) {
	if len(dec) == 0 {
		return "", nil
	}
	bytes, err := base64.StdEncoding.DecodeString(dec)
	return string(bytes), err
}
