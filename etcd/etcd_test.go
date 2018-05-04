package etcd

import (
	"testing"
)

func TestGetValue(t *testing.T) {
	Set("/banerwai/test/value", "value")
	_value, _ := GetValue("/banerwai/test/value")

	if _value != "value" {
		t.Errorf("etcd GetValue error")
	}
}
