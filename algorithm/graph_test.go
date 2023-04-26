package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndirectedGraphBFS(t *testing.T) {
	A := int('A' - 'A')
	B := int('B' - 'A')
	C := int('C' - 'A')
	D := int('D' - 'A')
	E := int('E' - 'A')
	F := int('F' - 'A')

	g := NewUndirectedGraph(6)
	g.AddEdge(A, B)
	g.AddEdge(A, C)
	g.AddEdge(B, A)
	g.AddEdge(B, C)
	g.AddEdge(B, D)
	g.AddEdge(C, A)
	g.AddEdge(C, B)
	g.AddEdge(C, D)
	g.AddEdge(C, E)
	g.AddEdge(D, C)
	g.AddEdge(D, E)
	g.AddEdge(D, F)
	g.AddEdge(E, C)
	g.AddEdge(E, D)
	g.AddEdge(F, D)

	ret := UndirectedGraphBFS(g, A)
	assert.EqualValues(t, []int{A, B, C, D, E, F}, ret)
}

func TestUndirectedGraphDFS(t *testing.T) {
	A := int('A' - 'A')
	B := int('B' - 'A')
	C := int('C' - 'A')
	D := int('D' - 'A')
	E := int('E' - 'A')
	F := int('F' - 'A')

	g := NewUndirectedGraph(6)
	g.AddEdge(A, B)
	g.AddEdge(A, C)
	g.AddEdge(B, A)
	g.AddEdge(B, C)
	g.AddEdge(B, D)
	g.AddEdge(C, A)
	g.AddEdge(C, B)
	g.AddEdge(C, D)
	g.AddEdge(C, E)
	g.AddEdge(D, C)
	g.AddEdge(D, E)
	g.AddEdge(D, F)
	g.AddEdge(E, C)
	g.AddEdge(E, D)
	g.AddEdge(F, D)

	ret := UndirectedGraphDFS(g, A)
	assert.EqualValues(t, []int{A, B, C, D, E, F}, ret)
}
