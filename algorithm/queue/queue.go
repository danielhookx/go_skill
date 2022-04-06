package queue

type Queue struct {
	queue []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		queue: make([]interface{}, 0),
	}
}

func (q *Queue) Enqueue(item interface{}) {
	q.queue = append(q.queue, item)
}

func (q *Queue) Dequeue() interface{} {
	v := q.queue[0]
	q.queue = q.queue[1:]
	return v
}

func (q *Queue) Len() int {
	return len(q.queue)
}
