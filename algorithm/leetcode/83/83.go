package _83

//83. Remove Duplicates from Sorted List
//Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.
//
//link: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil {
		for node.Next != nil && node.Val == node.Next.Val {
			//delete
			node.Next = node.Next.Next
		}
		node = node.Next
	}
	return head
}
