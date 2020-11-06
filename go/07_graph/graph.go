package graph

import (
	"container/list"
	"fmt"
)

// Graph 图结构体
type Graph struct {
	v   int // 顶点个数
	adj []*list.List
}

// NewGraph 新建图
func NewGraph(v int) *Graph {
	newGraph := &Graph{v, make([]*list.List, v)}
	for i := range newGraph.adj {
		newGraph.adj[i] = list.New()
	}
	return newGraph
}

// 添加边
func (g *Graph) addEdge(s, t int) {
	g.adj[s].PushBack(t)
	g.adj[t].PushBack(s)
}

// Print 打印
func (g *Graph) Print() {
	for i, l := range g.adj {
		fmt.Printf("顶点%2d: ", i)
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("%+v->", e.Value)
		}
		fmt.Printf("\n")
	}
}

// BFS 广度优先
func (g *Graph) BFS(s, t int) {
	// 构建和初始化Prev，用来记录路径
	prev := make([]int, g.v)
	for i := range prev {
		prev[i] = -1
	}
	var queue []int              // 遍历节点的队列
	visited := make([]bool, g.v) // 标记是否已经访问过
	// 将s放入队列中
	queue = append(queue, s)
	visited[s] = true
	// 从队列中取值处理
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		linkedList := g.adj[cur]
		// 遍历链表
		for e := linkedList.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] { //节点从来没有被访问
				prev[k] = cur
				if k == t {
					break
				}
				queue = append(queue, k)
				visited[k] = true
			}
		}
	}
	if len(prev) > 0 {
		printPrev(prev, s, t)
		fmt.Printf("\n")
	} else {
		fmt.Printf("No path %d -> %d", s, t)
	}
}

//print path recursively
func printPrev(prev []int, s, t int) {
	if s == t || prev[t] == -1 {
		fmt.Printf(" %d", t)
	} else {
		printPrev(prev, s, prev[t])
		fmt.Printf(" %d", t)
	}
}

// DFS 深度遍历优先
func (g *Graph) DFS(s, t int) {
	prev := make([]int, g.v)
	for i := range prev {
		prev[i] = -1
	}
	visited := make([]bool, g.v)

	isFound := false
	g.recurse(s, t, prev, visited, isFound)
	printPrev(prev, s, t)
	fmt.Printf("\n")
}

// 递归查询
func (g *Graph) recurse(s, t int, prev []int, visited []bool, isFound bool) {
	if isFound {
		return
	}
	visited[s] = true
	if s == t {
		isFound = true
		return
	}
	linkedList := g.adj[s]
	for e := linkedList.Front(); e != nil; e = e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			prev[k] = s
			g.recurse(k, t, prev, visited, false)
		}
	}
}
