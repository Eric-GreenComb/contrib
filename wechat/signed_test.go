package wechat

import (
	"strings"

	"testing"
)

func TestBindJSON(t *testing.T) {
	_reader := strings.NewReader(string("{\"id\":\"user01\",\"name\":\"eric\"}"))
	var props map[string]string
	BindJSON(_reader, &props)
	t.Logf("%v\n", props)

	if props["id"] != "user01" {
		t.Errorf("id error")
	}
}

func TestWxpayCalcSign(t *testing.T) {
	_reader := strings.NewReader(string("{\"id\":\"user01\",\"aname\":\"eric\"}"))
	var props map[string]string
	BindJSON(_reader, &props)
	t.Logf("%v\n", props)

	if props["id"] != "user01" {
		t.Errorf("id error")
	}

	// aname=eric&id=user01&key=greencomb
	_sign := WxpayCalcSign(props, "greencomb")

	if props["id"] != "user01" {
		t.Errorf("id error")
	}
	t.Logf("%s\n", _sign)
}
