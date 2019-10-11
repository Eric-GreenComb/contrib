package pbkdf2

import (
	"crypto/sha256"
	"encoding/base32"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// Hash hash
func Hash(password, salt string, iter, keyLen int) string {
	enc := pbkdf2.Key([]byte(password), []byte(salt), iter, keyLen, sha256.New)
	return strings.Trim(base32.StdEncoding.EncodeToString(enc), "=")
}
