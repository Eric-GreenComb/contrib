package inet

import (
	"fmt"
	"net"
	"testing"
)

func TestNtoa(t *testing.T) {

}

func TestAton(t *testing.T) {
	_ipnr := net.ParseIP("192.168.1.1")
	fmt.Println(Aton(_ipnr))

	_ipnr = net.ParseIP("192.168.1.2")
	fmt.Println(Aton(_ipnr))
}

func TestIsBelong(t *testing.T) {

	fmt.Println(IsBelong(`10.187.102.200`, `10.187.102.0/24`))

	fmt.Println(IsBelong(`10.187.101.8`, `10.187.102.0/24`))

	fmt.Println(IsBelong(`192.168.3.1`, `192.168.3.0/24`))
}
