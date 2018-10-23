package wechat

import (
	"strings"

	"testing"
)

func TestBindJSON(t *testing.T) {
	_reader := strings.NewReader(string("{\"id\":\"user01\",\"name\":\"eric\"}"))
	var props map[string]interface{}
	BindJSON(_reader, &props)
	t.Logf("%v\n", props)

	if props["id"] != "user01" {
		t.Errorf("id error")
	}
}

func TestWxpayCalcSign(t *testing.T) {
	_reader := strings.NewReader(string("{\"id\":123,\"aname\":\"eric\",\"name\":\"\"}"))
	var props map[string]interface{}
	BindJSON(_reader, &props)
	t.Logf("%v\n", props)

	// aname=eric&id=123&key=greencomb
	_sign := WxpayCalcSign(props, "greencomb")

	t.Logf("%s\n", _sign)
}
