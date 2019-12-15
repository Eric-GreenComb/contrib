package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JSONAuth JSONAuth
// FormAuth FormAuth
// salt:默认salt; appid:appid的key; sign: 签名的key; salts: salt组
//
// 如果请求参数中有appid的key，则匹配salts中配置的该app的salt，否则取默认的salt
func JSONAuth(salt, appid, sign string, salts map[string]interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 把request的内容读取出来
		var _bodyBytes []byte
		if c.Request.Body == nil {
			AbortWithError(c, http.StatusOK, GetLangContent("", "", "HTTP请求Body错误"))
			return
		}

		_bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		_reader := bytes.NewReader(_bodyBytes)
		var _props map[string]interface{}
		err := BindJSON(_reader, &_props)
		if err != nil {
			AbortWithError(c, http.StatusOK, GetLangContent("", "", "JSON数据格式错误"))
			return
		}

		// 处理各个appid对应的salt
		_salt := salt
		var _key string
		var _bapp bool
		// 如果没有输入sals列表，则采用默认salt
		if nil == salts {
			goto DefaultSalt
		}
		if len(appid) == 0 {
			goto DefaultSalt
		}
		// 如果没有函数接口appid参数的request请求参数，则采用默认salt
		_, _bapp = _props[appid]
		if !_bapp {
			goto DefaultSalt
		}
		// 如果接口salt列表没有匹配，则采用默认salt
		_key = _props[appid].(string)
		if _, _bsalt := salts[_key]; _bsalt {
			_salt = salts[_key].(string)
		}
	DefaultSalt:
		fmt.Println("salt:", _salt[:4])

		_sign := WechatSign(_props, _salt, sign)
		fmt.Println("====== api signed : ", _sign, _props[sign])
		if _props[sign] != _sign {
			AbortWithError(c, http.StatusOK, GetLangContent("", "", "签名错误"))
			return
		}
		// 把刚刚读出来的再写进去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(_bodyBytes))
		c.Next()
	}
}

// FormAuth FormAuth
// salt:默认salt; appid:appid的key; sign: 签名的key; salts: salt组
//
// 如果请求参数中有appid的key，则匹配salts中配置的该app的salt，否则取默认的salt
func FormAuth(salt, appid, sign string, salts map[string]interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 把request的内容读取出来
		var _bodyBytes []byte

		if c.Request.Body == nil {
			AbortWithError(c, http.StatusOK, GetLangContent("", "", "HTTP请求Body错误"))
			return
		}

		_bodyBytes, _ = ioutil.ReadAll(c.Request.Body)

		_formString := string(_bodyBytes)
		_maps, err := url.ParseQuery(_formString)
		if err != nil {
			AbortWithError(c, http.StatusOK, GetLangContent("", "", "解析HTTP参数错误"))
			return
		}

		_props := make(map[string]interface{})
		for _key, _value := range _maps {
			_props[_key] = _value[0]
		}

		// 处理各个appid对应的salt
		_salt := salt
		var _key string
		var _bapp bool
		// 如果没有输入sals列表，则采用默认salt
		if nil == salts {
			goto DefaultSalt
		}
		if len(appid) == 0 {
			goto DefaultSalt
		}
		// 如果没有函数接口appid参数的request请求参数，则采用默认salt
		_, _bapp = _props[appid]
		if !_bapp {
			goto DefaultSalt
		}
		// 如果接口salt列表没有匹配，则采用默认salt
		_key = _props[appid].(string)
		if _, _bsalt := salts[_key]; _bsalt {
			_salt = salts[_key].(string)
		}
	DefaultSalt:
		fmt.Println("salt:", _salt[:4])

		_sign := WechatSign(_props, salt, sign)
		fmt.Println("====== api signed : ", _sign, _props[sign])
		if _props[sign] != _sign {
			AbortWithError(c, http.StatusOK, GetLangContent("", "", "签名错误"))
			return
		}

		// 把刚刚读出来的再写进去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(_bodyBytes))
		c.Next()
	}
}

// GetLangContent GetLangContent
func GetLangContent(code, lang, _default string) string {
	if len(code) == 0 {
		return _default
	}

	if len(lang) == 0 {
		lang = "cn"
	}

	return _default
}

// AbortWithError AbortWithError
func AbortWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"flag": 2,
		"msg":  message,
		"data": "",
	})
	c.Abort()
}

// JWTAuth JWTAuth
func JWTAuth(issuer, key string) gin.HandlerFunc {
	return func(c *gin.Context) {

		_authorization := c.Request.Header.Get("Authorization")
		if len(_authorization) == 0 {
			JWTAbortWithError(c, http.StatusUnauthorized, GetLangContent("", "", "凭证无效"), "")
			return
		}

		_tokenString, err := StripBearerPrefixFromTokenString(_authorization)

		_j := &JWT{[]byte(key)}
		_claims, err := _j.Parse(_tokenString)
		if err != nil {
			fmt.Println("err:", err.Error(), _claims)
			JWTAbortWithError(c, http.StatusUnauthorized, GetLangContent("", "", "凭证无效"), "")
			return
		}

		if _claims.VerifyIssuer(issuer, false) == false {
			fmt.Println("凭证发放者不符")
			JWTAbortWithError(c, http.StatusUnauthorized, GetLangContent("", "", "凭证发放者不符"), "")
			return
		}

		c.Next()
	}
}

// CreateJWTString CreateJWTString
func CreateJWTString(id, sub, issuer, key string, expire int) string {

	_expire := time.Now().Add(time.Minute * time.Duration(expire)).Unix()
	// Claims schema of the data it will store
	claims := jwt.StandardClaims{
		Id:        id,
		Subject:   sub,
		NotBefore: int64(time.Now().Unix() - 600),
		ExpiresAt: _expire,
		Issuer:    issuer,
		IssuedAt:  time.Now().Unix(),
	}

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := _token.SignedString([]byte(key))
	return signedToken
}

// JWTAbortWithError JWTAbortWithError
func JWTAbortWithError(c *gin.Context, code int, message, realm string) {
	c.Header("WWW-Authenticate", "JWT realm="+realm)
	c.JSON(code, gin.H{
		"flag": code,
		"msg":  message,
	})
	c.Abort()
}

// WhitelistAuth WhitelistAuth
func WhitelistAuth(whitelist map[string]interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {

		_ip := c.ClientIP()
		if _, _inWhiteList := whitelist[_ip]; !_inWhiteList {
			AbortWithError(c, http.StatusOK, GetLangContent("", "", "非法请求"))
			return
		}

		c.Next()
	}
}
