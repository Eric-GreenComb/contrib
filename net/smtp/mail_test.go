package smtp

import (
	"testing"
)

func TestEmailServiceSend(t *testing.T) {
	var _email Email
	_email.Server("smtp.126.com:25", "user@u-sharing.com", "QWEasd@01")
	err := _email.Send("ministor@126.com", "Test", "Test", "html")

	if err != nil {
		t.Fatalf("send err : %v", err)
	}
}
