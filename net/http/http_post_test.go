package http

import (
	"net/url"
	"strings"
	"testing"
)

func TestPostForm(t *testing.T) {

	urlStr := "http://localhost:3000/post"
	form := make(url.Values)
	form.Set("email", "ministor@126.com")
	form.Set("pwd", "a11111")
	_ok, _ := PostForm(urlStr, form)

	if _ok != "OK" {
		t.Errorf("PostForm error")
	}
}

func TestPost(t *testing.T) {

	url := "http://localhost:3000/post"
	bodyType := "application/x-www-form-urlencoded"
	body := strings.NewReader("email=ministor@126.com&pwd=b11111")

	_ok, _ := Post(url, bodyType, body)

	if _ok != "OK" {
		t.Errorf("Post error")
	}
}
