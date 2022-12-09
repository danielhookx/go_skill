package _86

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	headDummy := &ListNode{}
	tailDummy := &ListNode{}

	headDummy.Next = head

	tail := tailDummy
	head = headDummy

	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			tmp := head.Next
			head.Next = head.Next.Next
			tmp.Next = nil

			tail.Next = tmp
			tail = tail.Next
		}
	}
	tail.Next = nil
	head.Next = tailDummy.Next
	return headDummy.Next
}
