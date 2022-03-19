package linked_list

import (
	"fmt"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	h := removeNthFromEnd(head, 2)
	for h != nil {
		fmt.Printf("%d,", h.Val)
		h = h.Next
	}
}

func Test_reverseList1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	h := reverseList1(head)
	for h != nil {
		fmt.Printf("%d,", h.Val)
		h = h.Next
	}
}

func Test_reverseList2(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	h := reverseList2(head)
	for h != nil {
		fmt.Printf("%d,", h.Val)
		h = h.Next
	}
}

func Test_isPalindrome(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}
	t.Log(isPalindrome(head))
}

func Test_hasCycle(t *testing.T) {
	n2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	n2.Next = n1
	t.Log(hasCycle(n1))
}
