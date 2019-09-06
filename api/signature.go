package api

import (
	"encoding/json"
	request "github.com/banerwai/global/bean"
	"github.com/banerwai/global/constant"
	"github.com/banerwai/gommon/crypto"
	"sort"
	"strconv"
)

// CheckAPIJson check api signature
func CheckAPIJson(jsonRequest string) bool {
	var request request.APIRequest
	json.Unmarshal([]byte(jsonRequest), &request)
	return CheckAPISignature(request)
}

// CheckAPISignature check api signature
func CheckAPISignature(request request.APIRequest) bool {
	var _sortKeys []string
	for k := range request.Payload.Params {
		_sortKeys = append(_sortKeys, k)
	}
	sort.Strings(_sortKeys)
	param := strconv.FormatInt(request.Payload.Iat, 10) + "."
	for _, k := range _sortKeys {
		param += request.Payload.Params[k]
	}

	bRet := false
	switch request.Signature.Alg {
	case "md5":
		param += constant.BanerwaiAPISignKey
		bRet = crypto.CompareMd5(param, request.Signature.Sign)
	case "hs256":
		bRet = crypto.CompareHS256Hex(param, constant.BanerwaiAPISignKey, request.Signature.Sign)
	case "hs512":
		bRet = crypto.CompareHS512Hex(param, constant.BanerwaiAPISignKey, request.Signature.Sign)
	case "sha256":
		param += constant.BanerwaiAPISignKey
		bRet = crypto.CompareSHA256Hex(param, request.Signature.Sign)
	case "sha512":
		param += constant.BanerwaiAPISignKey
		bRet = crypto.CompareSHA512Hex(param, request.Signature.Sign)
	}

	return bRet
}

// TransferMap2ArrayMap transfer map[string]string 2 map[string][]string( like url.Values)
func TransferMap2ArrayMap(mapIn map[string]string) map[string][]string {
	_aMap := make(map[string][]string)
	for k, v := range mapIn {
		_aMap[k] = []string{v}
	}
	return _aMap
}
