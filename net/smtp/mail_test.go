package smtp

import (
	"testing"
)

func TestEmailServiceSend(t *testing.T) {
	var _email Email
	_email.Server("smtp.126.com:25", "xxx@126.com", "xxx")
	err := _email.Send("xxx@126.com", "Test", "Test", "html")

	if err != nil {
		t.Fatalf("send err : %v", err)
	}
}
