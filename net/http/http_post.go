package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	httpbase "net/http"
	"net/url"
	"strings"
	"time"
)

// PostRawJSON post json []byte
func PostRawJSON(finalURL string, req []byte, response interface{}) (err error) {
	httpResp, err := httpbase.DefaultClient.Post(finalURL, "application/json;charset=utf-8", bytes.NewReader(req))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != httpbase.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	if err = json.NewDecoder(httpResp.Body).Decode(response); err != nil {
		return
	}
	return
}

// PostJSONObj post json object
func PostJSONObj(finalURL, _json string, response interface{}) (err error) {

	httpResp, err := httpbase.DefaultClient.Post(finalURL, "application/json;charset=utf-8", bytes.NewReader([]byte(_json)))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != httpbase.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	if err = json.NewDecoder(httpResp.Body).Decode(response); err != nil {
		return
	}
	return
}

// PostJSONString post json string and return response string
func PostJSONString(finalURL, _json string) (response string, err error) {

	httpResp, err := httpbase.DefaultClient.Post(finalURL, "application/json;charset=utf-8", bytes.NewReader([]byte(_json)))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != httpbase.StatusOK {
		return "", fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	_responseBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return "", err
	}

	return string(_responseBody), nil
}

// PostHTTPSJSONString post json string and return response string
func PostHTTPSJSONString(finalURL, _json string) (response string, err error) {

	tr := &httpbase.Transport{ //解决x509: certificate signed by unknown authority
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	_client := &httpbase.Client{
		Timeout:   10 * time.Second,
		Transport: tr, //解决x509: certificate signed by unknown authority
	}

	httpResp, err := _client.Post(finalURL, "application/json;charset=utf-8", bytes.NewReader([]byte(_json)))
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != httpbase.StatusOK {
		fmt.Println("POST Error : http.Status", httpResp.Status)
		return "", fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	_responseBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return "", err
	}

	return string(_responseBody), nil
}

// PostForm post form data
func PostForm(url string, data url.Values) (response string, err error) {
	resp, err := httpbase.PostForm(url, data)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// PostFormTimeout 通过client的Do方法执行
func PostFormTimeout(url string, data url.Values) (string, error) {
	/*     生成client,参数默认;
	*    这个结构体有四个属性
	*    Transport(RoundTrepper);指定执行的独立、单次http请求的机制
	*    CheckRedirect(func(req *Request, via []*Request)err):指定处理重定向的策略,如果不为nil,客户端在执行重定向之前调用本函数字段.如果CheckRedirect 返回一个错误，本类型中的get方法不会发送请求,如果CheckRedirect为nil,就会采用默认策略:连续请求10次后停止；
	＊    jar(CookieJar):jar指定cookie管理器,若果为nil请求中不会发送cookie,回复中的cookie会被忽略
	＊    TimeOut(time.Duration):指定本类型请求的时间限制，为0表示不设置超时
	*/
	//client := &http.Client{}    这里初始化了一个默认的client，并没有配置一些请求的设置

	//可以通过client中transport的Dail函数,在自定义Dail函数里面设置建立连接超时时长和发送接受数据超时
	_client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*3) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 2)) //设置发送接受数据超时
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 2,
		},
	}
	_reqest, err := http.NewRequest("POST", url, strings.NewReader(data.Encode())) //提交请求;用指定的方法，网址，可选的主体放回一个新的*Request
	_reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	_response, err := _client.Do(_reqest) //前面预处理一些参数，状态，Do执行发送；处理返回结果;Do:发送请求,
	if err != nil {
		return "", err
	}

	defer _response.Body.Close()

	body, err := ioutil.ReadAll(_response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Post post body
func Post(url string, bodyType string, body io.Reader) (response string, err error) {
	resp, err := httpbase.Post(url, bodyType, body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(_body), nil
}
