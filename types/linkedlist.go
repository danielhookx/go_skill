package types

type SingleListNode[T any] struct {
	Val  T
	Next *SingleListNode[T]
}

func NewSingleListFromArray[T any](src []T) *SingleListNode[T] {
	dummy := &SingleListNode[T]{}
	p := dummy
	for _, v := range src {
		p.Next = &SingleListNode[T]{Val: v}
		p = p.Next
	}
	return dummy.Next
}

// SingleListSet 节点值不重复的单链表封装
type SingleListSet[T comparable] struct {
	len        int
	head, tail *SingleListNode[T]
	nodes      map[T]*SingleListNode[T]
}

func NewSingleListSet[T comparable]() *SingleListSet[T] {
	l := &SingleListSet[T]{
		len: 0,
		head: &SingleListNode[T]{
			Val:  *new(T),
			Next: nil,
		},
		tail:  nil,
		nodes: make(map[T]*SingleListNode[T]),
	}
	l.tail = l.head
	return l
}

func (l *SingleListSet[T]) Add(val T) {
	if n, ok := l.nodes[val]; ok {
		n.Val = val
		return
	}
	l.tail.Next = &SingleListNode[T]{
		Val:  val,
		Next: nil,
	}
	l.tail = l.tail.Next
	l.nodes[val] = l.tail
	l.len++
}

func (l *SingleListSet[T]) Head() *SingleListNode[T] {
	return l.head.Next
}

func (l *SingleListSet[T]) Tail() *SingleListNode[T] {
	return l.tail
}
