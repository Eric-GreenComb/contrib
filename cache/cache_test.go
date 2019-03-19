package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestCache(t *testing.T) {

	type Token struct {
		Expire int64
		UserID string
		Name   string
	}
	var _new Token
	_new.Expire = time.Now().Add(time.Hour * 3).Unix()
	_new.UserID = "138"
	_new.Name = "eric"

	GCache.set(_new.UserID, _new)

	_token := GCache.get(_new.UserID)
	_obj := _token.(Token)
	fmt.Println(_obj.Name)

	type User struct {
		UserID string
		Name   string
		Phone  string
	}

	var _user User
	_user.UserID = "138"
	_user.Name = "eric"
	_user.Phone = "110"

	GCache.set(_new.UserID, _user)

	_user.Phone = "120"
	GCache.set(_new.UserID, _user)

	_user01 := GCache.get(_new.UserID)
	if _user01 == nil {
		fmt.Println("NIL")
		return
	}
	_obj1 := _user01.(User)
	fmt.Println(_obj1.Phone)

	_user01 = GCache.get(_new.UserID + "1")
	if _user01 == nil {
		fmt.Println("NIL")
		return
	}

	return

}
