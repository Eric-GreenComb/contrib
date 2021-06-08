package uuid

import (
	"fmt"
	"time"

	"testing"
)

func TestUUID(t *testing.T) {
	begin := time.Now()

	_uuid := UUID()
	fmt.Println(_uuid)

	fmt.Println(time.Since(begin))
}
