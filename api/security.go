package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// JSONAuth JSONAuth
func JSONAuth(salt, sign string, rule map[string]interface{}) gin.HandlerFunc {
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

		if _, ok := rule["mobi"]; ok {
			_param := rule["mobi"].(string)
			if _, ok := _props[_param]; ok {
				fmt.Println("检验", _props[_param])
			}
		}

		if _, ok := rule["email"]; ok {
			_param := rule["email"].(string)
			if _, ok := _props[_param]; ok {
				fmt.Println("检验", _props[_param])
			}
		}

		_sign := CalcSign(_props, salt, sign)
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
func FormAuth(salt, sign string, rule map[string]interface{}) gin.HandlerFunc {
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

		if _, ok := rule["mobi"]; ok {
			_param := rule["mobi"].(string)
			if _, ok := _props[_param]; ok {
				fmt.Println("检验", _props[_param])
			}
		}

		if _, ok := rule["email"]; ok {
			_param := rule["email"].(string)
			if _, ok := _props[_param]; ok {
				fmt.Println("检验", _props[_param])
			}
		}

		_sign := CalcSign(_props, salt, sign)
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
		_token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(key))
			return b, nil
		})
		if err != nil {
			fmt.Println(err.Error())
			JWTAbortWithError(c, http.StatusUnauthorized, GetLangContent("", "", "校验凭证错误"), "")
			return
		}

		claims := _token.Claims.(jwt.MapClaims)

		fmt.Println(claims)
		_now := time.Now().Unix()
		if claims.VerifyExpiresAt(_now, false) == false {
			JWTAbortWithError(c, http.StatusUnauthorized, GetLangContent("", "", "凭证过期"), "")
			return
		}

		if claims.VerifyIssuer(issuer, false) == false {
			JWTAbortWithError(c, http.StatusUnauthorized, GetLangContent("", "", "凭证发放者不符"), "")
			return
		}

		c.Next()
	}
}

// CreateJWTString CreateJWTString
func CreateJWTString(issuer, key string, expire int) string {

	_expire := time.Now().Add(time.Minute * time.Duration(expire)).Unix()
	// Claims schema of the data it will store
	claims := jwt.StandardClaims{
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
