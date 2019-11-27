package uuid

import (
	"fmt"
	"time"

	"testing"
)

func TestGetFlakeID(t *testing.T) {
	begin := time.Now()

	_uuid := GetFlakeID()
	fmt.Println(_uuid)

	fmt.Println(time.Since(begin))
}
