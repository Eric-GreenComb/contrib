package inet

import (
	"fmt"
	"testing"
)

func TestNtoa(t *testing.T) {

}

func TestAton(t *testing.T) {
	fmt.Println(Aton("192.168.1.1"))

	fmt.Println(Aton("127.0.0.1"))
}

func TestIsBelong(t *testing.T) {

	fmt.Println(IsBelong(`10.187.102.200`, `10.187.102.0/24`))

	fmt.Println(IsBelong(`10.187.101.8`, `10.187.102.0/24`))

	fmt.Println(IsBelong(`192.168.3.1`, `192.168.3.0/24`))
}
