package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Slice() []int {
	ret := make([]int, 0)
	for head := n; head != nil; head = head.Next {
		ret = append(ret, head.Val)
	}
	return ret
}

func NewSingleFromArray(src []int) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for _, v := range src {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return dummy.Next
}

// SingleSetList 节点值不重复的单链表封装
type SingleSetList struct {
	len        int
	head, tail *ListNode
	nodes      map[int]*ListNode
}

func NewSingleSetList() *SingleSetList {
	l := &SingleSetList{
		len: 0,
		head: &ListNode{
			Next: nil,
		},
		tail:  nil,
		nodes: make(map[int]*ListNode),
	}
	l.tail = l.head
	return l
}

func (l *SingleSetList) Add(v int) {
	if n, ok := l.nodes[v]; ok {
		n.Val = v
	} else {
		l.tail.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		l.tail = l.tail.Next
		l.nodes[v] = l.tail
		l.len++
	}
}

func (l *SingleSetList) Head() *ListNode {
	return l.head
}
