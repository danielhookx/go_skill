package list

type LinkedList struct {
	len        int
	head, tail *Node
	nodes      map[interface{}]*Node
}

func NewLinkedList() *LinkedList {
	l := &LinkedList{
		len: 0,
		head: &Node{
			next: nil,
		},
		tail:  nil,
		nodes: make(map[interface{}]*Node),
	}
	l.tail = l.head
	return l
}

func (t *LinkedList) Add(v interface{}) {
	if n, ok := t.nodes[v]; ok {
		n.SetVal(v)
	} else {
		t.tail.Add(v)
		t.tail = t.tail.next
		t.nodes[v] = t.tail
		t.len++
	}
}

func (t *LinkedList) Size() int {
	return t.len
}

func (t *LinkedList) Head() *Node {
	return t.head.next
}

type Node struct {
	val  interface{}
	next *Node
}

func (t *Node) SetVal(v interface{}) {
	t.val = v
}

func (t *Node) Val() interface{} {
	return t.val
}

func (t *Node) Next() *Node {
	return t.next
}

func (t *Node) Add(v interface{}) {
	tmp := t.next
	t.next = &Node{
		val:  v,
		next: tmp,
	}
}
