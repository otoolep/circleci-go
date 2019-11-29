package mod

import "testing"

func Test_Add3(t *testing.T) {
	if Add(1,2,3) != 6 {
		t.Fatal("add (3) doesn't work!")
	}
}
