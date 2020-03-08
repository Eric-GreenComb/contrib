package perm

// Str2Bin 二进制字符串转十进制数
func Str2Bin(s string) (num int64) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int64(s[l-i-1]) & 0xf) << uint8(i)
	}
	return
}
