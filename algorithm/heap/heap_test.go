package heap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Integer struct {
	value int
}

func (t *Integer) Less(p Interface) bool {
	return t.value < p.(*Integer).value
}

func TestHeapSort(t *testing.T) {
	const count = 20
	var src = make([]Interface, count)
	rand.Seed(time.Now().Unix())
	for i := 0; i < count; i++ {
		src[i] = &Integer{
			value: rand.Intn(count),
		}
	}
	for _, node := range src {
		fmt.Printf("%v,", node.(*Integer).value)
	}
	fmt.Println()
	h := BuildHeap(src)
	for _, node := range h.nodes {
		fmt.Printf("%v,", node.(*Integer).value)
	}
	fmt.Println()
	for v := h.Pop(); v != nil; v = h.Pop() {
	}
	for _, node := range h.nodes {
		fmt.Printf("%v,", node.(*Integer).value)
	}
}

func TestHeap_up(t *testing.T) {
	h := &Heap{
		nodes: []Interface{
			&Integer{
				value: 1,
			}, &Integer{
				value: 2,
			}, &Integer{
				value: 3,
			}, &Integer{
				value: 4,
			}, &Integer{
				value: 5,
			}, &Integer{
				value: 6,
			}, &Integer{
				value: 7,
			},
		},
	}
	h.up(len(h.nodes) - 1)
	for _, node := range h.nodes {
		fmt.Printf("%v,", node.(*Integer).value)
	}
}

func TestHeap_down(t *testing.T) {
	h := &Heap{
		nodes: []Interface{
			&Integer{
				value: 7,
			}, &Integer{
				value: 2,
			}, &Integer{
				value: 1,
			}, &Integer{
				value: 4,
			}, &Integer{
				value: 5,
			}, &Integer{
				value: 6,
			}, &Integer{
				value: 3,
			},
		},
		maxCount: 0,
	}
	h.last = len(h.nodes) - 1
	h.down(2)
	for _, node := range h.nodes {
		fmt.Printf("%v,", node.(*Integer).value)
	}
}

func TestHeap_Add(t *testing.T) {
	h := &Heap{
		nodes: []Interface{
			&Integer{
				value: 1,
			}, &Integer{
				value: 2,
			}, &Integer{
				value: 3,
			},
		},
		maxCount: 0,
	}
	h.last = len(h.nodes) - 1
	h.Add(&Integer{
		value: 4,
	})
	for _, node := range h.nodes {
		if node == nil {
			continue
		}
		fmt.Printf("%v,", node.(*Integer).value)
	}
}

func TestHeap_Pop(t *testing.T) {
	h := &Heap{
		nodes: []Interface{
			&Integer{
				value: 4,
			}, &Integer{
				value: 2,
			}, &Integer{
				value: 3,
			},
		},
		maxCount: 0,
	}
	h.last = len(h.nodes) - 1
	h.Pop()
	for _, node := range h.nodes {
		if node == nil {
			continue
		}
		fmt.Printf("%v,", node.(*Integer).value)
	}
}

func TestBuildHeap(t *testing.T) {
	h := BuildHeap([]Interface{
		&Integer{
			value: 7,
		}, &Integer{
			value: 2,
		}, &Integer{
			value: 1,
		}, &Integer{
			value: 4,
		}, &Integer{
			value: 5,
		}, &Integer{
			value: 6,
		}, &Integer{
			value: 3,
		},
	})
	for _, node := range h.nodes {
		fmt.Printf("%v,", node.(*Integer).value)
	}
}

func TestHeapSort1(t *testing.T) {
	h := BuildHeap([]Interface{
		&Integer{
			value: 7,
		}, &Integer{
			value: 2,
		}, &Integer{
			value: 1,
		}, &Integer{
			value: 4,
		}, &Integer{
			value: 5,
		}, &Integer{
			value: 6,
		}, &Integer{
			value: 3,
		},
	})
	for v := h.Pop(); v != nil; v = h.Pop() {
	}
	for _, node := range h.nodes {
		fmt.Printf("%v,", node.(*Integer).value)
	}
}
