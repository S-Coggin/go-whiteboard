package problems

import "testing"

// This is a just a smoke test to confirm that the testing framework is working.
func TestHello(t *testing.T) {
	if 2+2 != 4 {
		t.Fatal("math is broken")
	}
}
