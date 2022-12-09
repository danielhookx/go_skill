package algorithm

type Queue struct {
	queue []int
}

func NewQueue() *Queue {
	return &Queue{
		queue: make([]int, 0),
	}
}

func (q *Queue) EnQueue(v int) {
	q.queue = append(q.queue, v)
}

func (q *Queue) DeQueue() int {
	if len(q.queue) == 0 {
		return 0
	}
	v := q.queue[0]
	q.queue = q.queue[1:]
	return v
}

func (q *Queue) Len() int {
	return len(q.queue)
}
