package graph

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/queue"
)

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
