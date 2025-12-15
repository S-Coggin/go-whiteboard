package problems

import "testing"

func TestTwoSum(t *testing.T) {
	i, j := TwoSum([]int{2, 7, 11, 15}, 9)
	if i != 0 || j != 1 {
		t.Fatalf("expected (0,1), got (%d,%d)", i, j)
	}

	i, j = TwoSum([]int{3, 2, 4}, 6)
	if i != 1 || j != 2 {
		t.Fatalf("expected (1,2), got (%d,%d)", i, j)
	}

	i, j = TwoSum([]int{1, 2, 3}, 7)
	if i != -1 || j != -1 {
		t.Fatalf("expected (-1,-1), got (%d,%d)", i, j)
	}
}
