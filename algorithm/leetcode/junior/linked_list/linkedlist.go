package linked_list

//Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	del := node.Next
	node.Val = del.Val
	node.Next = del.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	interval := n
	step := 0
	last := head
	target := head
	for last.Next != nil {
		last = last.Next
		if step >= interval {
			target = target.Next
		} else {
			step++
		}
	}

	if target.Next == nil {
		head = nil
	} else {
		if step < interval {
			head = target.Next
		}
		del := target.Next
		target.Next = del.Next
	}
	return head
}

//O(n)
func reverseList1(head *ListNode) *ListNode {
	p := head
	font := head
	if head == nil {
		return head
	}
	//swap
	for p.Next != font {
		for p.Next != nil && p.Next != font {
			p.Val, p.Next.Val = p.Next.Val, p.Val
			p = p.Next
		}
		font = p
		p = head
	}
	return head
}

//O(log(n))
func reverseList2(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}
	return newHead
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var p = &ListNode{}
	head := p
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
			p = p.Next
		} else {
			p.Next = list2
			list2 = list2.Next
			p = p.Next
		}
	}
	if list1 != nil {
		p.Next = list1
	} else {
		p.Next = list2
	}
	return head.Next
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	slow = reverseList2(slow)
	fast = head
	for slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		if fast.Next == slow || fast.Next.Next == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
