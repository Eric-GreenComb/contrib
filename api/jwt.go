package api

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Err
var (
	ErrTokenExpired     error = errors.New("凭证过期")
	ErrTokenNotValidYet error = errors.New("凭证还未校验")
	ErrTokenMalformed   error = errors.New("凭证格式错误")
	ErrTokenInvalid     error = errors.New("凭证无效")
	ErrTokenIssuer      error = errors.New("凭证发放者不符")
)

// JWT JWT
type JWT struct {
	SigningKey []byte
}

// Create 创建token
func (j *JWT) Create(claims jwt.StandardClaims) (string, error) {
	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return _token.SignedString(j.SigningKey)
}

// Parse 解析token
func (j *JWT) Parse(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.SigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// Refresh 更新token
func (j *JWT) Refresh(tokenString, issuer string, expire int) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.SigningKey, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return "", ErrTokenInvalid
	}

	_claims, _ok := token.Claims.(*jwt.StandardClaims)
	if _ok && token.Valid {

		if _claims.VerifyIssuer(issuer, false) == false {
			fmt.Println("凭证发放者不符")
			return "", ErrTokenIssuer
		}

		_expire := time.Now().Add(time.Minute * time.Duration(expire)).Unix()
		jwt.TimeFunc = time.Now
		_claims.ExpiresAt = _expire
		return j.Create(*_claims)
	}

	return "", ErrTokenInvalid
}

// StripBearerPrefixFromTokenString Strips 'Bearer ' prefix from bearer token string
func StripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, ErrTokenInvalid
}
