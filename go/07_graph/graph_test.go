package graph

import "testing"

func TestGraph(t *testing.T) {
	g := NewGraph(10)
	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(2, 3)
	g.addEdge(5, 3)
	g.Print()
	// g.BFS(1, 5)
	g.DFS(1, 5)
}
