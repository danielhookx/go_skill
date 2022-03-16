package heap

type Interface interface {
	Less(p Interface) bool
}

type Heap struct {
	nodes    []Interface
	last     int
	maxCount int
}

func BuildHeap(src []Interface) *Heap {
	h := Heap{
		nodes:    src,
		last:     len(src) - 1,
		maxCount: len(src),
	}
	for i := len(src) / 2; i >= 0; i-- {
		h.down(i)
	}
	return &h
}

func BuildHeap2(src []Interface) *Heap {
	h := Heap{
		nodes:    make([]Interface, 0),
		last:     -1,
		maxCount: -1,
	}
	for _, i := range src {
		h.Add(i)
	}
	return &h
}

func (h *Heap) Add(item Interface) {
	if cap(h.nodes) <= h.last+1 {
		//grow
		h.grow(cap(h.nodes))
	}
	h.last++
	h.nodes[h.last] = item
	h.up(h.last)
}

func (h *Heap) Pop() Interface {
	if h.last <= 0 {
		return nil
	}
	h.swap(0, h.last)
	h.last--
	h.down(0)
	return h.nodes[h.last+1]
}

func (h *Heap) up(j int) {
	for {
		i := (j - 1) / 2 //parent
		if i >= j || h.nodes[j].Less(h.nodes[i]) {
			break
		}
		h.swap(i, j)
		j = i
	}
}

func (h *Heap) down(i int) {
	for {
		//left
		j1 := 2*i + 1
		if j1 > h.last || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 <= h.last && !h.nodes[j2].Less(h.nodes[j1]) {
			j = j2
		}
		if !h.nodes[i].Less(h.nodes[j]) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

func (h *Heap) swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

func (h *Heap) grow(size int) {
	if size == 0 {
		size = 100
	}
	new := make([]Interface, cap(h.nodes)+size)
	copy(new, h.nodes)
	h.nodes = new
}
