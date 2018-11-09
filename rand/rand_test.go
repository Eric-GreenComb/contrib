package rand

import (
	"testing"
)

func TestGen6Number(t *testing.T) {
	_rand := Gen6Number()
	t.Logf("%v\n", _rand)
}
