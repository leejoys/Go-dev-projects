package main

// // ErrVertExist - vertex already exist
// var ErrVertExist = fmt.Errorf("vertex already exist")

// // ErrNotAllVertExist - not all vertices exist
// var ErrNotAllVertExist = fmt.Errorf("not all vertices exist")

// //Vertex of graph
// type Vertex struct {
// 	Key      int
// 	Vertices map[int]*Vertex
// }

// // NewVertex is a constructor function for the Vertex
// func NewVertex(key int) *Vertex {
// 	return &Vertex{
// 		Key:      key,
// 		Vertices: map[int]*Vertex{},
// 	}
// }

// // Graph is graph struct
// type Graph struct {
// 	Vertices map[int]*Vertex
// 	directed bool
// }

// // NewDirectedGraph create directed graph
// func NewDirectedGraph() *Graph {
// 	return &Graph{
// 		Vertices: map[int]*Vertex{},
// 		directed: true,
// 	}
// }

// // NewUndirectedGraph create undirected graph
// func NewUndirectedGraph() *Graph {
// 	return &Graph{
// 		Vertices: map[int]*Vertex{},
// 	}
// }

// // AddVertex creates a new vertex with the given
// // key and adds it to the graph
// func (g *Graph) AddVertex(key int) error {

// 	if g.Vertices[key] != nil {
// 		return ErrVertExist
// 	}

// 	v := NewVertex(key)
// 	g.Vertices[key] = v
// 	return nil
// }

// // AddEdge adds an edge between two vertices in the graph
// func (g *Graph) AddEdge(k1, k2 int) error {
// 	v1 := g.Vertices[k1]
// 	v2 := g.Vertices[k2]

// 	if v1 == nil || v2 == nil {
// 		return ErrNotAllVertExist
// 	}

// 	if _, ok := v1.Vertices[v2.Key]; ok {
// 		return nil
// 	}

// 	v1.Vertices[v2.Key] = v2
// 	if !g.directed && v1.Key != v2.Key {
// 		v2.Vertices[v1.Key] = v1
// 	}

// 	g.Vertices[v1.Key] = v1
// 	g.Vertices[v2.Key] = v2
// 	return nil
// }
