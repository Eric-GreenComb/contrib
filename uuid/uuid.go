package uuid

import (
	"strings"

	googleuuid "github.com/pborman/uuid"
)

// UUID Google UUID
func UUID() string {
	return strings.Replace(googleuuid.NewRandom().String(), "-", "", -1)
}
