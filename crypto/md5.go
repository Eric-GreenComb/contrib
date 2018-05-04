package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 md5 encrypt 2 string
func Md5(source string) string {
	h := md5.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}

// CompareMd5 compare md5 value
func CompareMd5(source, md5 string) bool {
	if Md5(source) == md5 {
		return true
	}
	return false
}

// DoubleMd5 double md5 encrypt 2 string
func DoubleMd5(source string) string {
	return Md5(Md5(source))
}

// CompareDoubleMd5 compare double md5 value
func CompareDoubleMd5(source, md5 string) bool {
	_temp := Md5(source)
	if Md5(_temp) == md5 {
		return true
	}
	return false
}
