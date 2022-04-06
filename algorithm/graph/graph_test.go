package graph

import "testing"

func TestBFS(t *testing.T) {
	g := NewUndirectedGraph(6)
	g.AddEdge('A', 'B')
	g.AddEdge('A', 'C')
	g.AddEdge('B', 'A')
	g.AddEdge('B', 'C')
	g.AddEdge('B', 'D')
	g.AddEdge('C', 'A')
	g.AddEdge('C', 'B')
	g.AddEdge('C', 'D')
	g.AddEdge('C', 'E')
	g.AddEdge('D', 'C')
	g.AddEdge('D', 'E')
	g.AddEdge('D', 'F')
	g.AddEdge('E', 'C')
	g.AddEdge('E', 'D')
	g.AddEdge('F', 'D')

	BFS(g, 'A')
	BFS(g, 'E')
}

func TestDFS(t *testing.T) {
	g := NewUndirectedGraph(6)
	g.AddEdge('A', 'B')
	g.AddEdge('A', 'C')
	g.AddEdge('B', 'A')
	g.AddEdge('B', 'C')
	g.AddEdge('B', 'D')
	g.AddEdge('C', 'A')
	g.AddEdge('C', 'B')
	g.AddEdge('C', 'D')
	g.AddEdge('C', 'E')
	g.AddEdge('D', 'C')
	g.AddEdge('D', 'E')
	g.AddEdge('D', 'F')
	g.AddEdge('E', 'C')
	g.AddEdge('E', 'D')
	g.AddEdge('F', 'D')

	DFS(g, 'A')
	DFS(g, 'E')
}
