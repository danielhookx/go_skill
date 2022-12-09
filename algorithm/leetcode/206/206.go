package _206

//206. Reverse Linked List
//Given the head of a singly linked list, reverse the list, and return the reversed list.
//
//link:https://leetcode-cn.com/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = newHead
		//move
		newHead = head
		head = tmp
	}
	return newHead
}
