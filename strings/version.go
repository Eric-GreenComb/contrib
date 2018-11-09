package strings

import (
	"errors"
	"strings"
)

// CompareVersion CompareVersion
// return big:1 smaill:2 equ:0 error:-1
func CompareVersion(ver1, ver2 string) (int, error) {

	verStrArr1 := SpliteStrByNet(ver1)
	verStrArr2 := SpliteStrByNet(ver2)

	lenStr1 := len(verStrArr1)
	lenStr2 := len(verStrArr2)

	if lenStr1 != lenStr2 {
		return -1, errors.New("version format error")
	}

	return CompareArrStrVers(verStrArr1, verStrArr2), nil
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
