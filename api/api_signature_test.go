package api

import (
	"testing"

	"encoding/json"
	"fmt"
	request "github.com/banerwai/global/bean"
	"github.com/banerwai/global/constant"
	"github.com/banerwai/gommon/crypto"
	"sort"
	"strconv"
	"time"
)

func TestCheckAPIJson(t *testing.T) {

	var APIRequest request.APIRequest

	var APIRequestHeader request.APIRequestHeader
	APIRequestHeader.Method = "invoke"
	APIRequestHeader.Ver = "1.0.0"
	APIRequest.Header = APIRequestHeader

	var APIRequestPayload request.APIRequestPayload
	APIRequestPayload.Iss = "Eric"
	APIRequestPayload.Iat = time.Now().Unix()
	APIRequestPayload.Aud = "www.banerwai.com"

	_mapParams := make(map[string]string)

	_mapParams["name"] = "eric"
	_mapParams["tel"] = "13811111111"
	APIRequestPayload.Params = _mapParams

	APIRequest.Payload = APIRequestPayload

	var APIRequestSignature request.APIRequestSignature

	var _sortKeys []string
	for k := range _mapParams {
		_sortKeys = append(_sortKeys, k)
	}
	sort.Strings(_sortKeys)
	param := strconv.FormatInt(APIRequestPayload.Iat, 10) + "."
	for _, k := range _sortKeys {
		param += _mapParams[k]
	}

	// APIRequestSignature.Alg = "md5"
	// param += constant.BanerwaiAPISignKey
	// APIRequestSignature.Sign = crypto.Md5(param)

	APIRequestSignature.Alg = "hs256"
	APIRequestSignature.Sign = crypto.HS256Hex(param, constant.BanerwaiAPISignKey)

	// APIRequestSignature.Alg = "hs512"
	// APIRequestSignature.Sign = crypto.HS512Hex(param, constant.BanerwaiAPISignKey)

	// APIRequestSignature.Alg = "sha256"
	// param += constant.BanerwaiAPISignKey
	// APIRequestSignature.Sign = crypto.SHA256Hex(param)

	// APIRequestSignature.Alg = "sha512"
	// param += constant.BanerwaiAPISignKey
	// APIRequestSignature.Sign = crypto.SHA512Hex(param)

	APIRequest.Signature = APIRequestSignature

	b, err := json.Marshal(APIRequest)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

	if !CheckAPIJson(string(b)) {
		t.Errorf("CheckAPIJson error")
	}
}
