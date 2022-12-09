package algorithm

type Stack struct {
	stack []int
}

func NewStack() *Stack {
	return &Stack{
		stack: make([]int, 0),
	}
}

func (s *Stack) Push(item int) {
	s.stack = append(s.stack, item)
}

func (s *Stack) Pop() int {
	if len(s.stack) == 0 {
		return 0
	}
	v := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return v
}

func (s *Stack) Len() int {
	return len(s.stack)
}
