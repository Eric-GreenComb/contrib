package uuid

import (
	"fmt"
	"time"

	"testing"
)

func TestNewGuid(t *testing.T) {
	begin := time.Now()

	_uuid := NewGUID()
	fmt.Println(_uuid)

	fmt.Println(time.Since(begin))
}

func TestNewGuidWith(t *testing.T) {
	begin := time.Now()

	_uuid := NewGUIDWith("ministor@126.com")
	fmt.Println(_uuid)

	fmt.Println(time.Since(begin))
}

func TestUUID(t *testing.T) {
	begin := time.Now()

	_uuid := UUID()
	fmt.Println(_uuid)

	fmt.Println(time.Since(begin))
}
