package _21

//21. Merge Two Sorted Lists
//You are given the heads of two sorted linked lists list1 and list2.
//
//Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
//Return the head of the merged linked list.
//
//link:https://leetcode-cn.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p := &ListNode{}
	head := p

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			p.Next = list2
			list2 = list2.Next
			p = p.Next
		} else {
			p.Next = list1
			list1 = list1.Next
			p = p.Next
		}
	}
	if list1 == nil {
		p.Next = list2
	}
	if list2 == nil {
		p.Next = list1
	}
	return head.Next
}
