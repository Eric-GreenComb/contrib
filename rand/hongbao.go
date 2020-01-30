package rand

import (
	"math"
	"math/rand"
	baserand "math/rand"
	"time"
)

// 随机红包
// remainCount: 剩余红包数
// remainMoney: 剩余红包金额（单位：分)
func randomMoney(remainCount int, remainMoney int64) int64 {
	if remainCount == 1 {
		return remainMoney
	}

	// baserand.Seed(time.Now().UnixNano())

	var _min int64
	_min = 1
	_max := remainMoney / int64(remainCount) * 2
	// return baserand.Int63n(_max) + int64(_min)
	return xRandom(_min, _max)
}

// xRandom xRandom
// 生产min和max之间的随机数，但是概率不是平均的，从min到max方向概率逐渐加大。
// 先平方，然后产生一个平方值范围内的随机数，再开方，这样就产生了一种“膨胀”再“收缩”的效果。
func xRandom(min, max int64) int64 {
	_sqr := sqr(max - min)
	baserand.Seed(time.Now().UnixNano())
	_next := baserand.Int63n(_sqr)
	_float64 := float64(_next)
	_sqrt := math.Sqrt(_float64)
	return int64(_sqrt)
}

func sqr(x int64) int64 {
	return x * x
}

// RedPackage 发红包
// count: 红包数量
// money: 红包金额（单位：分)
func RedPackage(count int, money int64) []int64 {
	var _ret []int64
	for i := 0; i < count; i++ {
		_m := randomMoney(count-i, money)
		_ret = append(_ret, _m)
		money -= _m
	}
	return _ret
}

// Shuffle 洗牌算法
func Shuffle(arr []int64) {
	rand.Seed(time.Now().UnixNano())
	var i, j int
	var temp int64
	for i = len(arr) - 1; i > 0; i-- {
		j = rand.Intn(i + 1)
		temp = arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
}
