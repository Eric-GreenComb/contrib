package api

import (
	"html/template"
)

// FilterXSS FilterXSS
func FilterXSS(s string) string {
	return template.HTMLEscapeString(s)
}
