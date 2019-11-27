package crypto

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {

	_encodeMd5 := Md5("go语言中，判断两个字符串是否相等")

	t.Logf("%s\n", _encodeMd5)

	_bTrue := CompareMd5("go语言中，判断两个字符串是否相等", _encodeMd5)

	if !_bTrue {
		t.Errorf("CompareMd5 error")
	}

	_encodeDoubleMd5 := Md5(_encodeMd5)

	_bTrue = CompareDoubleMd5("go语言中，判断两个字符串是否相等", _encodeDoubleMd5)
	if !_bTrue {
		t.Errorf("CompareDoubleMd5 error")
	}

	fmt.Println("MD5: " + _encodeMd5)
}

func BenchmarkMd5(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DoubleMd5("13810167616")
	}

}
