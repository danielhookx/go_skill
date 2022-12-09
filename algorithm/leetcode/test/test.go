package test

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	i := 0
	slip := &ListNode{}
	current := &ListNode{Next: head}
	head = current
	for ; i < left-1; i++ {
		current = current.Next
	}
	startTemp := current
	endTemp := current.Next

	var tmp *ListNode
	for ; i < right; i++ {
		tmp = current.Next
		current.Next = slip.Next
		slip.Next = current
		current = tmp
	}
	startTemp.Next = current
	endTemp.Next = tmp
	return head.Next
}
