package rand

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestRandomStringLength(t *testing.T) {
	ls := []int{1, 5, 10, 50, 100}
	dict := CharsASCII

	for _, l := range ls {
		s, err := Generate(l, dict)
		if err != nil {
			t.Error(err)
		}

		if len(s) != l {
			t.Errorf("unexpected length of generated string: want %d, got %d", l, len(s))
		}
	}
}

func TestRandomStringInvalidLength(t *testing.T) {
	ls := []int{-1, 0}

	for _, l := range ls {
		if _, err := Generate(l, CharsAlpha); err != ErrInvalidLengthSpecified {
			t.Errorf("unexpected error for length %d: %v", l, err)
		}
	}
}

func TestRandomStringInvalidChars(t *testing.T) {
	if _, err := Generate(1, ""); err != ErrInvalidDictSpecified {
		t.Errorf("unexpected error using empty dictionary: %v", err)
	}
}

func TestRandomStringWithNonASCII(t *testing.T) {
	const (
		dict = "世界"
		l    = 5
	)

	s, err := Generate(l, dict)
	if err != nil {
		t.Fatal(err)
	}
	if got := utf8.RuneCountInString(s); got != l {
		t.Fatalf("invalid length of string, got %d, want %d", got, l)
	}
}

func TestPwdGenerate(t *testing.T) {
	dict := CharsASCII
	fmt.Println(Generate(20, dict))
}
