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

func TestFormCalcSignAppend(t *testing.T) {
	query := "id=123&name=eric&api_token=41d4bf7a6e3aac84669493fd2a8df473&a="

	_maps, _ := url.ParseQuery(query)
	_props := make(map[string]interface{})
	for _key, _value := range _maps {
		_props[_key] = _value[0]
	}

	t.Logf("%v\n", _props)

	// VeZ16GuXyKaYhKbA?a=&id=123&name=eric
	_sign := CalcSignAppend(_props, "VeZ16GuXyKaYhKbA?", "api_token")

	t.Logf("%s\n", _sign)
	fmt.Println(_sign)

	t.Logf("%s\n", _props["api_token"])

	if _props["api_token"] != _sign {
		t.Errorf("signed error")
	}
}

func TestFormAPICalcSignNil(t *testing.T) {
	query := "api_token=53f29a0e5a243dd78639d2dc7120a8af&a="

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
