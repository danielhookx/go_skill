package stack

type Stack struct {
	stack []interface{}
}

func NewStack() *Stack {
	return &Stack{
		stack: make([]interface{}, 0),
	}
}

func (q *Stack) Push(item interface{}) {
	q.stack = append(q.stack, item)
}

func (q *Stack) Pop() interface{} {
	v := q.stack[len(q.stack)-1]
	q.stack = q.stack[:len(q.stack)-1]
	return v
}

func (q *Stack) Len() int {
	return len(q.stack)
}
