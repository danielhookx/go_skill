package algorithm

// UndirectedGraph 无向图邻接表形式
type UndirectedGraph struct {
	v   int
	adj map[int]*SingleSetList
}

func NewUndirectedGraph(v int) *UndirectedGraph {
	g := &UndirectedGraph{
		v:   v,
		adj: make(map[int]*SingleSetList),
	}
	return g
}

func (g *UndirectedGraph) AddEdge(s, t int) {
	if _, ok := g.adj[s]; !ok {
		g.adj[s] = NewSingleSetList()
	}
	if _, ok := g.adj[t]; !ok {
		g.adj[t] = NewSingleSetList()
	}
	g.adj[s].Add(t)
	g.adj[t].Add(s)
}

func UndirectedGraphBFS(g *UndirectedGraph, start int) []int {
	ret := make([]int, 0)
	q := NewQueue[int]()
	visited := make(map[int]bool)
	q.EnQueue(start)
	visited[start] = true

	for q.Len() > 0 {
		vertex := q.DeQueue()
		l := g.adj[vertex]
		for node := l.Head(); node != nil; node = node.Next {
			if v, ok := visited[node.Val]; v && ok {
				continue
			}
			q.EnQueue(node.Val)
			visited[node.Val] = true
		}
		ret = append(ret, vertex)
	}
	return ret
}

func UndirectedGraphDFS(g *UndirectedGraph, start int) []int {
	ret := make([]int, 0)
	q := NewStack()
	visited := make(map[int]bool)
	q.Push(start)
	visited[start] = true

	for q.Len() > 0 {
		vertex := q.Pop()
		l := g.adj[vertex]
		for node := l.Head(); node != nil; node = node.Next {
			if v, ok := visited[node.Val]; v && ok {
				continue
			}
			q.Push(node.Val)
			visited[node.Val] = true
		}
		ret = append(ret, vertex)
	}
	return ret
}
