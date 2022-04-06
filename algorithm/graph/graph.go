package graph

import (
	"fmt"
	"github.com/oofpgDLD/go_skill/algorithm/list"
	"github.com/oofpgDLD/go_skill/algorithm/queue"
	"github.com/oofpgDLD/go_skill/algorithm/stack"
)

type UndirectedGraph struct {
	v   int //
	adj map[interface{}]*list.LinkedList
}

func NewUndirectedGraph(v int) *UndirectedGraph {
	g := &UndirectedGraph{
		v:   v,
		adj: make(map[interface{}]*list.LinkedList),
	}
	return g
}

func (g *UndirectedGraph) AddEdge(s, t interface{}) {
	if _, ok := g.adj[s]; !ok {
		g.adj[s] = list.NewLinkedList()
	}
	if _, ok := g.adj[t]; !ok {
		g.adj[t] = list.NewLinkedList()
	}
	g.adj[s].Add(t)
	g.adj[t].Add(s)
}

//
func BFS(g *UndirectedGraph, s interface{}) {
	//
	q := queue.NewQueue()
	visited := make(map[interface{}]bool)
	q.Enqueue(s)
	visited[s] = true

	for q.Len() > 0 {
		vertex := q.Dequeue()
		l := g.adj[vertex]
		for node := l.Head(); node != nil; node = node.Next() {
			if v, ok := visited[node.Val()]; v && ok {
				continue
			}
			q.Enqueue(node.Val())
			visited[node.Val()] = true
		}
		fmt.Printf("%c\n", vertex)
	}
}

//
func DFS(g *UndirectedGraph, s interface{}) {
	//
	q := stack.NewStack()
	visited := make(map[interface{}]bool)
	q.Push(s)
	visited[s] = true

	for q.Len() > 0 {
		vertex := q.Pop()
		l := g.adj[vertex]
		for node := l.Head(); node != nil; node = node.Next() {
			if v, ok := visited[node.Val()]; v && ok {
				continue
			}
			q.Push(node.Val())
			visited[node.Val()] = true
		}
		fmt.Printf("%c\n", vertex)
	}
}
