package rand

import (
	"fmt"
	"math/rand"
	"time"
)

// Gen6Number gen 6 rand number
func Gen6Number() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

// GetRandomString 生成随机字符串
func GetRandomString(num int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandomItNum 返回输入参数以下的随机数字，如输入参数为5，随机返回0,1,2,3,4
func GetRandomItNum(num int) int {
	if num == 0 {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(num)
}
