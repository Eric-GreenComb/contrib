package uuid

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"

	googleuuid "github.com/pborman/uuid"

	"github.com/banerwai/gommon/crypto"
)

// UUID Google UUID
func UUID() string {
	return strings.Replace(googleuuid.NewRandom().String(), "-", "", -1)
}

// NewGUID md5 Guid
func NewGUID() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return crypto.Md5(base64.URLEncoding.EncodeToString(b))
}

// NewGUIDWith md5 Guid 后面加个str生成之, 更有保障, 确保唯一
func NewGUIDWith(str string) string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return crypto.Md5(base64.URLEncoding.EncodeToString([]byte(string(b) + str)))
}
