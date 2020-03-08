package perm

import (
	"fmt"
	"testing"
)

func TestPermAdd(t *testing.T) {

	fmt.Println(Add("1111", "100000"))

	fmt.Println(AddN(Str2Bin("1111"), 1<<5))
}

func TestPermDel(t *testing.T) {

	fmt.Println(Del("101111", "100000"))

	fmt.Println(DelN(Str2Bin("101111"), Str2Bin("100000")))
}

func TestPermCheck(t *testing.T) {

	fmt.Println(Check("00001111", "00000100", "00000100"))

	fmt.Println(CheckN(Str2Bin("00001111"), Str2Bin("00000100"), Str2Bin("00000100")))
}

func TestPermMerge(t *testing.T) {

	fmt.Println(Merge("00001001", "00000100"))

	fmt.Println(MergeN(Str2Bin("00001001"), Str2Bin("00000100")))
}
