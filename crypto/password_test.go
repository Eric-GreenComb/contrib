package crypto

import (
	// "encoding/hex"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {

	_b, _ := GenerateHash("ministor")

	_is := CompareHash(_b, "ministor")

	if !_is {
		t.Errorf("CompareHash %b", _is)
	}

	// fmt.Println(hex.EncodeToString(_b))
	fmt.Println(string(_b))
}
