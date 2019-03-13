package strings

import (
	"strings"
)

// CompareVersion CompareVersion
// return big:1 smaill:2 equ:0 error:-1
func CompareVersion(ver1, ver2 string) (int, error) {

	_verStrArr1 := SpliteStrByNet(ver1)
	_verStrArr2 := SpliteStrByNet(ver2)

	_lenStr1 := len(_verStrArr1)
	_lenStr2 := len(_verStrArr2)

	_len := _lenStr1
	if _lenStr2 > _len {
		_len = _lenStr2
	}

	verArr1 := genStringArray(_verStrArr1, _len)
	verArr2 := genStringArray(_verStrArr2, _len)

	return CompareArrStrVers(verArr1, verArr2), nil
}

func genStringArray(strs []string, lens int) []string {
	_rets := make([]string, lens)
	_len := len(strs)
	for _index := range _rets {
		if _index >= _len {
			_rets[_index] = "0"
			continue
		}
		_rets[_index] = strs[_index]
	}
	return _rets
}

// CompareArrStrVers CompareArrStrVers
func CompareArrStrVers(ver1, ver2 []string) int {

	for index := range ver1 {

		littleResult := CompareLittleVer(ver1[index], ver2[index])

		if littleResult != 0 {
			return littleResult
		}
	}
	return 0
}

// CompareLittleVer CompareLittleVer
// 比较小版本号字符串
func CompareLittleVer(ver1, ver2 string) int {

	bytes1 := []byte(ver1)
	bytes2 := []byte(ver2)

	len1 := len(bytes1)
	len2 := len(bytes2)
	if len1 > len2 {
		return 1
	}

	if len1 < len2 {
		return 2
	}

	//如果长度相等则按byte位进行比较
	return CompareByBytes(bytes1, bytes2)
}

// CompareByBytes CompareByBytes 按byte位进行比较小版本号
func CompareByBytes(ver1, ver2 []byte) int {

	for index := range ver1 {
		if ver1[index] > ver2[index] {
			return 1
		}
		if ver1[index] < ver2[index] {
			return 2
		}
	}
	return 0
}

// SpliteStrByNet 按“.”分割版本号为小版本号的字符串数组
func SpliteStrByNet(version string) []string {
	return strings.Split(version, ".")
}
