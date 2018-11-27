package strings

import (
	"fmt"
	"testing"
)

func TestVersion(t *testing.T) {
	version1 := "1.0.1"
	version2 := "1.0.1.1"
	fmt.Println(CompareVersion(version1, version2))

	version1 = "1.0.13"
	version2 = "1.0.1a"
	fmt.Println(CompareVersion(version1, version2))

	version1 = "1.0.131"
	version2 = "1.0.1a"
	fmt.Println(CompareVersion(version1, version2))

	version1 = "1.1.131"
	version2 = "1.10.1a"
	fmt.Println(CompareVersion(version1, version2))
}