package problems

import "testing"

func buildList(vals []int) *ListNode {
	var head, tail *ListNode
	for _, v := range vals {
		n := &ListNode{Val: v}
		if head == nil {
			head, tail = n, n
		} else {
			tail.Next = n
			tail = n
		}
	}
	return head
}

func listToSlice(h *ListNode) []int {
	out := []int{}
	for h != nil {
		out = append(out, h.Val)
		h = h.Next
	}
	return out
}

func TestReverseList(t *testing.T) {
	h := buildList([]int{1, 2, 3, 4})
	r := ReverseList(h)

	got := listToSlice(r)
	want := []int{4, 3, 2, 1}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("expected %v, got %v", want, got)
		}
	}
}
