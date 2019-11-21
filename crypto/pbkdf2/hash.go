package pbkdf2

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base32"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// HashCustom HashCustom
// salt 盐 iter 迭代次数,keylen 长度
func HashCustom(source, salt string, iter, keyLen int) string {
	enc := pbkdf2.Key([]byte(source), []byte(salt), iter, keyLen, sha256.New)
	return strings.Trim(base32.StdEncoding.EncodeToString(enc), "=")
}

// Hash Hash
// iter 迭代次数,keylen 长度
func Hash(source string, iter, keylen int) string {
	h := md5.New()
	h.Write([]byte(source))
	_md5 := h.Sum(nil)
	enc := pbkdf2.Key([]byte(source), _md5, iter, keylen, sha256.New)
	return strings.Trim(base32.StdEncoding.EncodeToString(enc), "=")
}
