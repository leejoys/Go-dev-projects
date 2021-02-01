package main

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

	// fmt.Println(g.Vertices[v1.Key])
	// fmt.Println(v1)
	// fmt.Println(g.Vertices[v2.Key])
	// fmt.Println(v2)

	// g.Vertices[v1.Key] = v1
	// g.Vertices[v2.Key] = v2
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
func BFS(g *Graph, startVertex *Vertex) map[int]*Vertex { //

	q := &queue{}
	visited := map[int]*Vertex{}
	current := startVertex

	for {

		visited[current.Key] = current

		for _, v := range current.Vertices {
			if _, ok := visited[v.Key]; !ok {
				q.enqueue(v)
			}
		}

		current = q.dequeue()
		if current == nil {
			break
		}
	}
	return visited
}

func main() {

	g := NewUndirectedGraph()

	err := g.AddEdge(1, 2)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 7; i++ {
		err := g.AddVertex(i + 1)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = g.AddVertex(5)
	if err != nil {
		fmt.Println(err)
	}

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

	searchTree := BFS(g, g.Vertices[1])

	fmt.Println(searchTree)
	fmt.Println(g)
	fmt.Println(g.Vertices[1])
	fmt.Println(g.Vertices[2])
	fmt.Println(g.Vertices[3])
	fmt.Println(g.Vertices[4])
	fmt.Println(g.Vertices[5])
	fmt.Println(g.Vertices[6])
	fmt.Println(g.Vertices[7])

	q := &queue{}
	for _, v := range g.Vertices {
		fmt.Println(v)
		q.enqueue(v)
	}

	for range g.Vertices {
		fmt.Println(q.dequeue())
	}

}
