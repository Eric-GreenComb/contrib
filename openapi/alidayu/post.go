package alidayu

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func init() {
	var (
		appKey    = flag.String("alidayu.appkey", "changeit", "alidayu appkey")
		appSecret = flag.String("alidayu.appsecret", "changeit", "alidayu appsecret")
	)
	AppKey = *appKey
	AppSecret = *appSecret
}

// DoPost post data 2 alidayu
func DoPost(m map[string]string) (success bool, response string) {
	if AppKey == "" || AppSecret == "" {
		return false, "AppKey or AppSecret is requierd!"
	}

	body, size := getRequestBody(m)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", URL, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.ContentLength = size

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		response = err.Error()
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	response = string(data)
	if strings.Contains(response, "success") {
		return true, response
	}
	return false, response
}

func getRequestBody(m map[string]string) (reader io.Reader, size int64) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	v := url.Values{}

	signString := AppSecret
	for _, k := range keys {
		v.Set(k, m[k])
		signString += k + m[k]
	}
	signString += AppSecret

	signByte := md5.Sum([]byte(signString))
	sign := strings.ToUpper(fmt.Sprintf("%x", signByte))
	v.Set("sign", sign)

	return ioutil.NopCloser(strings.NewReader(v.Encode())), int64(len(v.Encode()))
}
