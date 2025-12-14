package problems

import "testing"

func TestHello(t *testing.T) {
  if 2+2 != 4 {
    t.Fatal("math is broken")
  }
}
