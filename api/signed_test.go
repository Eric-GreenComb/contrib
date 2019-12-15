package api

import (
	"fmt"
	"net/url"
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

func TestCalcSign(t *testing.T) {
	_reader := strings.NewReader(string("{\"id\":123,\"name\":\"\",\"api_token\":\"499456AFE7368947C23C5AA8EF64F939\"}"))
	var props map[string]interface{}
	BindJSON(_reader, &props)
	t.Logf("%v\n", props)

	// VeZ16GuXyKaYhKbA?id=123&name=
	_sign := CalcSign(props, "VeZ16GuXyKaYhKbA?", "api_token")

	t.Logf("%s\n", _sign)

	t.Logf("%s\n", props["api_token"])

	if props["api_token"] != _sign {
		t.Errorf("signed error")
	}
}

func TestFormCalcSign(t *testing.T) {
	query := "id=123&name=eric&api_token=02c32e5f1eb9bb0707f3b185719c42e5&a="

	_maps, _ := url.ParseQuery(query)
	_props := make(map[string]interface{})
	for _key, _value := range _maps {
		_props[_key] = _value[0]
	}

	t.Logf("%v\n", _props)

	// VeZ16GuXyKaYhKbA?a=&id=123&name=eric
	_sign := CalcSign(_props, "VeZ16GuXyKaYhKbA?", "api_token")

	t.Logf("%s\n", _sign)

	t.Logf("%s\n", _props["api_token"])

	if _props["api_token"] != _sign {
		t.Errorf("signed error")
	}
}

func TestFormWechatSign(t *testing.T) {
	query := "id=123&name=eric&api_token=487fb41ebe2db640d8842a01bfb4bb69&a="

	_maps, _ := url.ParseQuery(query)
	_props := make(map[string]interface{})
	for _key, _value := range _maps {
		_props[_key] = _value[0]
	}

	t.Logf("%v\n", _props)

	// id=123&name=eric&VeZ16GuXyKaYhKbA?
	_sign := WechatSign(_props, "VeZ16GuXyKaYhKbA?", "api_token")

	t.Logf("%s\n", _sign)
	fmt.Println(_sign)

	t.Logf("%s\n", _props["api_token"])

	if _props["api_token"] != _sign {
		t.Errorf("signed error")
	}
}

func TestFormWechatSignNil(t *testing.T) {
	query := "api_token=ab1d2b98f8f1c53a930074166502b570&a="

	_maps, _ := url.ParseQuery(query)
	_props := make(map[string]interface{})
	for _key, _value := range _maps {
		_props[_key] = _value[0]
	}
	t.Logf("%v\n", _props)

	// VeZ16GuXyKaYhKbA?
	_sign := WechatSign(_props, "VeZ16GuXyKaYhKbA?", "api_token")

	t.Logf("%s\n", _sign)

	t.Logf("%s\n", _props["api_token"])

	if _props["api_token"] != _sign {
		t.Errorf("signed error")
	}
}

func TestJSONWechatSign(t *testing.T) {
	_reader := strings.NewReader(string("{\"id\":123,\"name\":\"\",\"api_token\":\"44940fa863af338f69b70af2d7ea4acd\"}"))
	var props map[string]interface{}
	BindJSON(_reader, &props)
	fmt.Println("props:", props)

	// id=123&VeZ16GuXyKaYhKbA?
	_sign := WechatSign(props, "VeZ16GuXyKaYhKbA?", "api_token")

	t.Logf("%s\n", _sign)

	t.Logf("%s\n", props["api_token"])

	if props["api_token"] != _sign {
		t.Errorf("signed error")
	}
}

func TestJSONWechatSignNil(t *testing.T) {
	_reader := strings.NewReader(string("{\"name\":\"\",\"api_token\":\"ab1d2b98f8f1c53a930074166502b570\"}"))
	var props map[string]interface{}
	BindJSON(_reader, &props)
	fmt.Println("props:", props)

	// VeZ16GuXyKaYhKbA?
	_sign := WechatSign(props, "VeZ16GuXyKaYhKbA?", "api_token")

	t.Logf("%s\n", _sign)

	t.Logf("%s\n", props["api_token"])

	if props["api_token"] != _sign {
		fmt.Println(props["api_token"], _sign)
		t.Errorf("signed error")
	}
}
