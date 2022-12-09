package _92

//92. Reverse Linked List II
//Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.
//
//link:https://leetcode-cn.com/problems/reverse-linked-list-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy

	index := 0
	var prev *ListNode
	for index < left {
		prev = head
		head = head.Next
		index++
	}

	var next *ListNode
	mid := head
	for head != nil && index <= right {
		temp := head.Next
		head.Next = next
		next = head
		head = temp
		index++
	}

	prev.Next = next
	mid.Next = head
	return dummy.Next
}
