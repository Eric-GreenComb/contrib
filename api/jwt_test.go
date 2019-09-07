package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestCreateToken(t *testing.T) {

	_j := &JWT{[]byte("keyjwt")}

	_expire := time.Now().Add(time.Minute * time.Duration(10)).Unix()

	_claims := jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix() - 600),
		ExpiresAt: _expire,
		Issuer:    "unionpro",
		IssuedAt:  time.Now().Unix(),
	}

	_token, err := _j.Create(_claims)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(_token)
	}
}

func TestParseToken(t *testing.T) {

	_j := &JWT{[]byte("keyjwt")}

	_claims, err := _j.Parse("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Njc3OTAzMjUsImlhdCI6MTU2Nzc4OTcyNSwiaXNzIjoidW5pb25wcm8iLCJuYmYiOjE1Njc3ODkxMjV9.SahYxxfHezPcNhWl9eJ3-9izQ5Xlzh_PV37vlxJerKA")
	if err != nil {
		fmt.Println("err:", err.Error(), _claims)
		return
	}

	if _claims.VerifyIssuer("unionpro", false) == false {
		fmt.Println("凭证发放者不符")
		return
	}
}

func TestRefresh(t *testing.T) {

	_j := &JWT{[]byte("keyjwt")}

	_token, err := _j.Refresh("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Njc3OTAzOTcsImlhdCI6MTU2Nzc4OTcyNSwiaXNzIjoidW5pb25wcm8iLCJuYmYiOjE1Njc3ODkxMjV9.pTUnTuxvFCd01hTK1YJWfkdk56D0iv0OAGknroNBVoc", "unionpro", 10)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	fmt.Println(_token)
}
