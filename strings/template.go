package strings

import (
	"bytes"
	"html/template"
)

// ParseTpl parse tpl
func ParseTpl(tplname, tpl string, mmap map[string]string) string {
	_tpl, _err := template.New(tplname).Parse(tpl)
	if _err != nil {
		return ""
	}

	b := bytes.NewBuffer(make([]byte, 0))

	_err = _tpl.Execute(b, mmap)
	if _err != nil {
		return ""
	}
	return b.String()
}
