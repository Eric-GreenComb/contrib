package inet

import (
	"errors"
	"fmt"
	"net"
)

// PrivateIP2Lower16 PrivateIP2Lower16
func PrivateIP2Lower16() int64 {
	_uint16, err := Lower16BitPrivateIP()
	if err != nil {
		fmt.Println(err.Error())
		return 1000
	}
	return int64(_uint16)
}

// Lower16BitPrivateIP Lower16BitPrivateIP
func Lower16BitPrivateIP() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}
	fmt.Println(ip)

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("no private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}
