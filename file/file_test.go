package file

import (
	"testing"
)

func TestIsFileExist(t *testing.T) {

	_is := IsFileExist("/home/eric/go/src/key.pem")

	if !_is {
		t.Errorf("IsEmail %b", _is)
	}

	_is = IsFileExist("/home/eric/go/src/key1.pem")

	if _is {
		t.Errorf("IsEmail %b", _is)
	}
}
