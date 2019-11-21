package pbkdf2

import (
	"fmt"
	"testing"
)

func BenchmarkHashCustom(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HashCustom("13810167616", "f5d3c7fa", 1, 16)
	}

}

func TestCheckHashCustom(t *testing.T) {

	_unique := HashCustom("13810167616", "f5d3c7fa", 1000, 12)
	fmt.Println(_unique)

	_hash := "PBKDYMJJHZXJPQHGCODH3PQ"
	fmt.Println(_hash[:3], _hash[3:])
	if _hash[:3] == "PBK" {
		if _unique != _hash[3:] {
			t.Errorf("Unique error")
		}
	} else {
		t.Errorf("prefix error")
	}

	_unique = HashCustom("13810167616", "f5d3c7fa", 1000, 12)
	fmt.Println(_unique)

	if _unique != "DYMJJHZXJPQHGCODH3PQ" {
		t.Errorf("Unique error")
	}

	_unique = HashCustom("13810167616", "f5d3c7fa", 1000, 8)
	fmt.Println(_unique)

	if _unique != "DYMJJHZXJPQHG" {
		t.Errorf("Unique error")
	}

}

func BenchmarkHash(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hash("13810167616", 1, 12)
	}

}

func TestCheckHash(t *testing.T) {

	_unique := Hash("13810167616", 1, 12)
	fmt.Println(_unique)

	_unique = Hash("ministor@126.com", 1, 12)
	fmt.Println(_unique)

	_pwd := Hash("123456", 1, 8)
	fmt.Println(_pwd)
}
