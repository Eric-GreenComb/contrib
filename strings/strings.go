package strings

import (
	"bytes"
)

// SubstringByte 以byte来截取
func SubstringByte(str string, start int) string {
	return substr(str, start, len(str)-start, false)
}

// Substring sub string
func Substring(str string, start int) string {
	return substr(str, start, len(str)-start, true)
}

// Substr sub string
func Substr(str string, start, length int) string {
	return substr(str, start, length, true)
}

func substr(str string, start, length int, isRune bool) string {
	rs := []rune(str)
	rs2 := []byte(str)

	rl := len(rs)
	if !isRune {
		rl = len(rs2)
	}
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	if isRune {
		return string(rs[start:end])
	}
	return string(rs2[start:end])
}

// ConstructString construct string
func ConstructString(strings ...string) string {
	var buf bytes.Buffer
	for _, str := range strings {
		buf.WriteString(str)
	}
	return buf.String()
}
