package pbkdf2

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {

	_unique := Hash("13810167616", "f5d3c7fa", 16)
	fmt.Println(_unique)

	if _unique != "I7KFGKAXXVLQCPQ7TKKUSNEVDQ" {
		t.Errorf("Unique error")
	}

	_unique = Hash("13810167616", "f5d3c7fa", 12)
	fmt.Println(_unique)

	if _unique != "I7KFGKAXXVLQCPQ7TKKQ" {
		t.Errorf("Unique error")
	}

	_unique = Hash("13810167616", "f5d3c7fa", 8)
	fmt.Println(_unique)

	if _unique != "I7KFGKAXXVLQC" {
		t.Errorf("Unique error")
	}

}

func TestCheckHash(t *testing.T) {

	_unique := Hash("13810167616", "f5d3c7fa", 12)
	fmt.Println(_unique)

	_hash := "PBKI7KFGKAXXVLQCPQ7TKKQ"
	fmt.Println(_hash[:3], _hash[3:])
	if _hash[:3] == "PBK" {
		if _unique != _hash[3:] {
			t.Errorf("Unique error")
		}
	} else {
		t.Errorf("prefix error")
	}

	_unique = Hash("13810167616", "f5d3c7fa", 12)
	fmt.Println(_unique)

	if _unique != "I7KFGKAXXVLQCPQ7TKKQ" {
		t.Errorf("Unique error")
	}

	_unique = Hash("13810167616", "f5d3c7fa", 8)
	fmt.Println(_unique)

	if _unique != "I7KFGKAXXVLQC" {
		t.Errorf("Unique error")
	}

}
