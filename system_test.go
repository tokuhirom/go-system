package system

import "testing"

func TestSimple(t *testing.T) {
	if System("ls system_test.go") != 0 {
		t.Fail()
	}
}
