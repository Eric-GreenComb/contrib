package wechat

import (
	"strings"

	"testing"
)

func TestBindJSON(t *testing.T) {
	_reader := strings.NewReader(string("{\"id\":\"user01\",\"name\":\"eric\",\"key\":\"092982528307ED4B4C7654DF0790E4B2\"}"))
	var props map[string]interface{}
	BindJSON(_reader, &props)
	t.Logf("%v\n", props)

	if props["id"] != "user01" {
		t.Errorf("id error")
	}
}

func TestParseQuery(t *testing.T) {
	query := "id=123&name=eric&signed=3c41117cb2dfdf70000840abd381646a&a="
	// var _props map[string]interface{}
	_props := make(map[string]interface{})
	ParseQuery(query, _props)
	t.Logf("%v\n", _props)
	t.Logf("====%s====\n", _props["a"])
}

func TestWxpayCalcSign(t *testing.T) {
	_reader := strings.NewReader(string("{\"id\":123,\"name\":\"eric\",\"signed\":\"3C41117CB2DFDF70000840ABD381646A\"}"))
	var props map[string]interface{}
	BindJSON(_reader, &props)
	t.Logf("%v\n", props)

	// id=123&name=eric&key=VeZ16GuXyKaYhKbA
	_sign := WxpayCalcSign(props, "VeZ16GuXyKaYhKbA")

	t.Logf("%s\n", _sign)

	t.Logf("%s\n", props["signed"])

	if props["signed"] != _sign {
		t.Errorf("signed error")
	}
}

func TestFormWxpayCalcSign(t *testing.T) {
	query := "id=123&name=eric&signed=3C41117CB2DFDF70000840ABD381646A&a="
	// var _props map[string]interface{}
	_props := make(map[string]interface{})
	ParseQuery(query, _props)
	t.Logf("%v\n", _props)

	// id=123&name=eric&key=VeZ16GuXyKaYhKbA
	_sign := WxpayCalcSign(_props, "VeZ16GuXyKaYhKbA")

	t.Logf("%s\n", _sign)

	t.Logf("%s\n", _props["signed"])

	if _props["signed"] != _sign {
		t.Errorf("signed error")
	}
}
