package pbkdf2

import (
	"crypto/sha256"
	"encoding/base32"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// Unique generates unique id based on the current id with a prefix and salt.
func Unique(password, salt string) string {
	enc := pbkdf2.Key([]byte(password), []byte(salt), 4096, 16, sha256.New)
	return strings.Trim(base32.StdEncoding.EncodeToString(enc), "=")
}
