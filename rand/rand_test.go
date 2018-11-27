package rand

import (
	"testing"
)

func TestGen6Number(t *testing.T) {
	_rand := Gen6Number()
	t.Logf("%v\n", _rand)
}

func TestGetRandomNumber(t *testing.T) {
	_rand := GetRandomNumber(6)
	t.Logf("%v\n", _rand)

	_rand = GetRandomNumber(5)
	t.Logf("%v\n", _rand)

	_rand = GetRandomNumber(4)
	t.Logf("%v\n", _rand)
}
