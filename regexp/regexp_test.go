package regexp

import (
	"fmt"
	"testing"
)

func TestIsEmail(t *testing.T) {

	_is := IsEmail("ministor@126.com")

	if !_is {
		t.Errorf("Isn't Mail")
	}

	_is = IsEmail("ministor11@126.cn.com")

	if !_is {
		t.Errorf("Isn't Mail")
	}

	_is = IsEmail("minis@tor11@126.cn.com")

	if _is {
		t.Errorf("Isn't Mail")
	}
}

func TestIsMobile(t *testing.T) {

	_is := IsMobile("13810167616")

	if !_is {
		t.Errorf("Isn't Mobile")
	}

}

func TestMobileReplaceRepl(t *testing.T) {

	_str := MobileReplaceRepl("13810167616")

	fmt.Println(_str)
	if _str != "138****7616" {
		t.Errorf("MobileReplaceRepl")
	}
}
