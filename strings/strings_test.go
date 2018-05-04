package strings

import (
	"testing"
)

func TestSubstring(t *testing.T) {

	_is := Substring("ministor@126.com", 5)

	if _is != "tor@126.com" {
		t.Errorf("Substring error")
	}

}

func TestMarkdown(t *testing.T) {
	b := MarkdownBasic2Html("# Test")
	if b != "<h1>Test</h1>\n" {
		t.Errorf("MarkdownBasic2Html error")
	}
}

func TestParseIDCard(t *testing.T) {
	p, y := ParseIDCard("231011198111266811")
	if p != "黑龙江省" || y != "1981" {
		t.Errorf("ParseIDCard error")
	}
}

func TestConstructString(t *testing.T) {
	s := ConstructString("1", "2")
	if s != "12" {
		t.Errorf("ConstructString error")
	}

	q := ConstructString("http://www.xxx.com/s?", "q=2", "&pt=0")
	if q != "http://www.xxx.com/s?q=2&pt=0" {
		t.Errorf("ConstructString error")
	}
}
