package rand

import (
	"fmt"
	"testing"
)

func TestRedPackage(t *testing.T) {
	fmt.Println(RedPackage(20, 500))
}

func TestXRandom(t *testing.T) {
	fmt.Println(XRandom(20, 500))
}
