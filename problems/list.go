package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList reverses a singly linked list.
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
