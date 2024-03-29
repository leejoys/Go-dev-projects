package main

// Вариант программы для поиска кратчайшего пути в неориентированном графе

import (
	"fmt"
)

//Graph strucrture ==================================================================

// ErrVertExist - vertex already exist
var ErrVertExist = fmt.Errorf("vertex already exist")

// ErrNotAllVertExist - not all vertices exist
var ErrNotAllVertExist = fmt.Errorf("not all vertices exist")

//Vertex of graph
type Vertex struct {
	Key      int
	Vertices map[int]*Vertex
	Parent   *Vertex
	Distance int
	InQueue  bool
}

// NewVertex is a constructor function for the Vertex
func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[int]*Vertex{},
	}
}

// Graph is graph struct
type Graph struct {
	Vertices map[int]*Vertex
	directed bool
}

// NewDirectedGraph create directed graph
func NewDirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
		directed: true,
	}
}

// NewUndirectedGraph create undirected graph
func NewUndirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
	}
}

// AddVertex creates a new vertex with the given
// key and adds it to the graph
func (g *Graph) AddVertex(key int) error {

	if g.Vertices[key] != nil {
		return ErrVertExist
	}

	v := NewVertex(key)
	g.Vertices[key] = v
	return nil
}

// AddEdge adds an edge between two vertices in the graph
func (g *Graph) AddEdge(k1, k2 int) error {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	if v1 == nil || v2 == nil {
		return ErrNotAllVertExist
	}

	if _, ok := v1.Vertices[v2.Key]; ok {
		return nil
	}

	v1.Vertices[v2.Key] = v2
	if !g.directed && v1.Key != v2.Key {
		v2.Vertices[v1.Key] = v1
	}

	return nil
}

// Queue structure ===================================================================

// node that holds the graphs vertex as data
type node struct {
	v    *Vertex
	next *node
}

// queue data structure
type queue struct {
	head *node
	tail *node
}

// enqueue adds a new node to the tail of the queue
func (q *queue) enqueue(v *Vertex) {
	n := &node{v: v}

	if q.tail == nil {
		q.head = n
		q.tail = n
		return
	}

	q.tail.next = n
	q.tail = n
}

//  enqueueHead adds a new node to the head of the queue
func (q *queue) enqueueHead(v *Vertex) {
	n := &node{v: v}

	if q.tail == nil {
		q.head = n
		q.tail = n
		return
	}
	n.next = q.head
	q.head = n
}

// dequeue removes the head from the queue and returns it
func (q *queue) dequeue() *Vertex {
	n := q.head

	if n == nil {
		return nil
	}

	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return n.v
}

// BFS function ==========================================================================

//BFS - breadth-first search
func BFS(g *Graph, start *Vertex) map[int]*Vertex { //

	for _, v := range g.Vertices {
		v.Parent, v.Distance, v.InQueue = nil, 0, false
	}

	q := &queue{}
	visited := map[int]*Vertex{}
	current := start

	for {

		visited[current.Key] = current

		for _, v := range current.Vertices {
			if _, ok := visited[v.Key]; ok || v.InQueue {
				continue
			}
			q.enqueue(v)
			v.InQueue = true
			v.Parent = current
			v.Distance = current.Distance + 1
		}

		current = q.dequeue()
		if current == nil {
			break
		}
	}
	return visited
}

// ShortestBFS function ===================================================================

// ErrNoWay - there is no way between vertices
var ErrNoWay = fmt.Errorf("there is no way between vertices")

// // ShortestBFS search with ReQueue
// func ShortestBFS(k1, k2 int, g *Graph) *queue {
// 	searchTree := BFS(g, g.Vertices[k1])
// if _, ok := searchTree[k2]; !ok {
// 	return nil, ErrNoWay
// }
// 	q := &queue{}
// 	callBack := func(v *Vertex) {
// 		q.enqueueHead(v)
// 	}
// 	ReQueue(k2, searchTree, callBack)
// 	return q
// }

// // ReQueue recursively fill the queue
// func ReQueue(k int, searchTree map[int]*Vertex, callBack func(v *Vertex)) {
// 	v := searchTree[k]
// 	callBack(v)
// 	if v.Parent != nil {
// 		ReQueue(v.Parent.Key, searchTree, callBack)
// 	}
// }

// ShortestBFS search with anonimus recursion
func ShortestBFS(k1, k2 int, g *Graph) (*queue, error) {
	searchTree := BFS(g, g.Vertices[k1])
	if _, ok := searchTree[k2]; !ok {
		return nil, ErrNoWay
	}
	q := &queue{}
	var f func(int)

	f = func(k int) {
		v := searchTree[k]
		q.enqueueHead(v)
		if v.Parent != nil {
			f(v.Parent.Key)
		}
	}

	f(k2)
	return q, nil
}

func main() {

	// graph
	g := NewUndirectedGraph()

	// ErrNotAllVertExist test
	err := g.AddEdge(1, 2)
	if err != nil {
		fmt.Println(err)
	}

	// create vertices
	for i := 0; i < 8; i++ {
		err := g.AddVertex(i + 1)
		if err != nil {
			fmt.Println(err)
		}
	}

	// ErrVertExist test
	err = g.AddVertex(5)
	if err != nil {
		fmt.Println(err)
	}

	// begin create edges
	err = g.AddEdge(1, 2)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(2, 3)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(2, 4)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(4, 1)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(4, 5)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(5, 6)
	if err != nil {
		fmt.Println(err)
	}
	err = g.AddEdge(6, 7)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(5, 7)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(4, 7)
	if err != nil {
		fmt.Println(err)
	}
	// end of edges

	path, err := ShortestBFS(1, 6, g)
	if err != nil {
		fmt.Println(err)
	} else {
		for path.tail != nil {
			fmt.Println(path.dequeue().Key)
		}
	}

	fmt.Println("")

	path, err = ShortestBFS(6, 3, g)
	if err != nil {
		fmt.Println(err)
	} else {
		for path.tail != nil {
			fmt.Println(path.dequeue().Key)
		}
	}

	// fmt.Println(g.Vertices[1])
	// fmt.Println(g.Vertices[2])
	// fmt.Println(g.Vertices[3])
	// fmt.Println(g.Vertices[4])
	// fmt.Println(g.Vertices[5])
	// fmt.Println(g.Vertices[6])
	// fmt.Println(g.Vertices[7])
	// fmt.Println(g.Vertices[8])
}
