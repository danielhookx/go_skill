package list

import "testing"

func TestLinkedList_Add(t1 *testing.T) {
	l := NewLinkedList()

	l.Add(2)
	l.Add(2)
	l.Add(3)
	l.Add(4)

	for node := l.Head(); node != nil; node = node.next {
		t1.Log(node.val)
	}
}
