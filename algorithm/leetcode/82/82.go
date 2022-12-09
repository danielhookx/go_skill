package _82

//82. Remove Duplicates from Sorted List II
//Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.
//
//link:https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	node := dummy
	delVar := 0
	for node.Next != nil && node.Next.Next != nil {
		if node.Next.Val == node.Next.Next.Val {
			delVar = node.Next.Val
			for node.Next != nil && delVar == node.Next.Val {
				//delete
				node.Next = node.Next.Next
			}
		} else {
			node = node.Next
		}
	}
	return dummy.Next
}
