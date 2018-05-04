package archive

import (
	"testing"
)

func TestCompress(t *testing.T) {

	_compress := Compress("ministor@126.com")
	_decompress := Decompress(_compress)

	if _decompress != "ministor@126.com" {
		t.Errorf("Compress error")
	}
}
