package mod

import "testing"

func Test_Add(t *testing.T) {
	if Add(1,2) != 3 {
		t.Fatal("add doesn't work!")
	}
}
