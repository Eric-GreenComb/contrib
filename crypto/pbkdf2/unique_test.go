package pbkdf2

import (
	"fmt"
	"testing"
)

func TestUnique(t *testing.T) {

	_unique := Unique("123456", "f5d3c7fa")
	fmt.Println(_unique)

	if _unique != "NNMDSO6U5BUSIXMBINZHCKU4HY" {
		t.Errorf("Unique error")
	}
}
