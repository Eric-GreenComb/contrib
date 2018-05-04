package crypto

import (
	"golang.org/x/crypto/bcrypt"
	mathRand "math/rand"
	"time"
)

// GenerateHash generates bcrypt hash from plaintext password
func GenerateHash(password string) ([]byte, error) {
	hex := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(hex, 10)
	if err != nil {
		return hashedPassword, err
	}
	return hashedPassword, nil
}

// CompareHash compares bcrypt password with a plaintext one. Returns true if passwords match
// and false if they do not.
func CompareHash(digest []byte, password string) bool {
	hex := []byte(password)
	if err := bcrypt.CompareHashAndPassword(digest, hex); err == nil {
		return true
	}
	return false
}

// RandomPwd 随机密码
// num 几位
func RandomPwd(num int) string {
	chars := make([]byte, 62)
	j := 0
	for i := 48; i <= 57; i++ {
		chars[j] = byte(i)
		j++
	}
	for i := 65; i <= 90; i++ {
		chars[j] = byte(i)
		j++
	}
	for i := 97; i <= 122; i++ {
		chars[j] = byte(i)
		j++
	}
	j--

	str := ""
	mathRand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		x := mathRand.Intn(j)
		str += string(chars[x])
	}

	return str
}

// GetRandomString 生成随机字符串
func GetRandomString(num int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

const (
	// KcRandKindNum 纯数字
	KcRandKindNum = 0
	// KcRandKindLower 小写字母
	KcRandKindLower = 1
	// KcRandKindUpper 大写字母
	KcRandKindUpper = 2
	// KcRandKindAll 数字、大小写字母
	KcRandKindAll = 3
)

// Krand 随机字符串
func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	mathRand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = mathRand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + mathRand.Intn(scope))
	}
	return string(result)
}

// GetRandomItNum 返回输入参数以下的随机数字，如输入参数为5，随机返回0,1,2,3,4
func GetRandomItNum(num int) int {
	if num == 0 {
		return 0
	}
	r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	return r.Intn(num)
}
