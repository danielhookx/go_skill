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
