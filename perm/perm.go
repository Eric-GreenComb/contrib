// 使用二进制位来控制权限
// Copyright 2020 The Gopper.in. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package perm

import (
	"fmt"
)

// Add 权限的增加，参数为字符串
// 规则如下
// 原二进制位权限 | 要添加的权限二进制位
// 首先得到权限的perm: 用到位运算的左移运算
// 1 << bit_pos
// 1 << 5 = 00100000
// 再和00001111 进行或运算
// 	00100000 | 00001111 = 00101111
func Add(perm, bitpos string, n int) (retn int64, rets string) {
	_perm := Str2Bin(perm)
	_bitpos := Str2Bin(bitpos)
	retn = _perm | _bitpos
	rets = fmt.Sprintf("%0*b", n, retn)
	return
}

// AddN 权限的增加，参数为int64
// 规则如下
// 原二进制位权限 | 要添加的权限二进制位
// 首先得到权限的perm: 用到位运算的左移运算
// 1 << bit_pos
// 1 << 5 = 00100000
// 再和00001111 进行或运算
// 	00100000 | 00001111 = 00101111
func AddN(perm, bitpos int64, n int) (retn int64, rets string) {
	retn = perm | bitpos
	rets = fmt.Sprintf("%0*b", n, retn)
	return
}

// Del 删除权限，参数为字符串
// 原二进制位权限 ^ 也就是 00101111 ^ 00100000 也就是 00001111
func Del(perm, bitpos string, n int) (retn int64, rets string) {
	_perm := Str2Bin(perm)
	_bitpos := Str2Bin(bitpos)
	retn = _perm ^ _bitpos
	rets = fmt.Sprintf("%0*b", n, retn)
	return
}

// DelN 删除权限，参数为int64
// 原二进制位权限 ^ 也就是 00101111 ^ 00100000 也就是 00001111
func DelN(perm, bitpos int64, n int) (retn int64, rets string) {
	retn = perm ^ bitpos
	rets = fmt.Sprintf("%0*b", n, retn)
	return
}

// Check 校验权限，参数为字符串
// 原二进制位权限 & 也就是 00001111 & 00000100 也就是 00000100
func Check(perm, bitpos string) bool {
	_perm := Str2Bin(perm)
	_bitpos := Str2Bin(bitpos)
	_check := Str2Bin(bitpos)
	_retn := _perm & _bitpos
	return _retn == _check
}

// CheckN 校验权限，参数为int64
// 原二进制位权限 & 也就是 00001111 & 00000100 也就是 00000100
func CheckN(perm, bitpos int64) bool {
	_retn := perm & bitpos
	return _retn == bitpos
}

// Merge 合并权限，参数为字符串
// 原二进制位权限 | 也就是 00001001 | 00000100 也就是 00001101
func Merge(perm, bitpos string, n int) (retn int64, rets string) {
	_perm := Str2Bin(perm)
	_bitpos := Str2Bin(bitpos)
	_retn := _perm | _bitpos
	return _retn, fmt.Sprintf("%0*b", n, _retn)
}

// MergeN 合并权限，参数为int64
// 原二进制位权限 | 也就是 00001001 | 00000100 也就是 00001101
func MergeN(perm, bitpos int64, n int) (retn int64, rets string) {
	_retn := perm | bitpos
	return _retn, fmt.Sprintf("%0*b", n, _retn)
}
