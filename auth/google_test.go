package auth

import (
	"encoding/base32"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestOneTimePassword(t *testing.T) {

	// DPI45HKISEXU6HG7
	secretKey := "FIQNXC5VT7QBGDFRRX7AXBTC2HJ6RNIO" // 16*N 个字符
	secretKeyUpper := strings.ToUpper(secretKey)
	fmt.Println(secretKeyUpper)
	key, err := base32.StdEncoding.DecodeString(secretKeyUpper)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// generate a one-time password using the time at 30-second intervals
	epochSeconds := time.Now().Unix()
	pwd := OneTimePassword(key, toBytes(epochSeconds/30))

	secondsRemaining := 30 - (epochSeconds % 30)
	fmt.Printf("%06d (%d second(s) remaining)\n", pwd, secondsRemaining)

}
