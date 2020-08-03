/*
Usage:
	import this "graph" package in Go main file

Creating Graph:
Syntax: g := CreateGraph(e ...Edge) Graph {...}
		'e' is a variadic parameter of type 'Edge'. So it takes any number of aruguments (>= 0)
		Returns a 'Graph' type -> Graph{M:map[string]Node}; Node type = Node{N:map[string]int, E:bool} 
	Ex: g1 = CreateGraph() -> {map[]}
		or
		g2 = CreateGraph(Edge{"v1", "v2", 4}) -> {map[v1:{map[v2:4] false} v2:{map[] false}]}
		or 
		g3 = CreateGraph(Edge{"v1", "v2", 4}, Edge{"v3", "v4", 4}) -> {map[v1:{map[v2:4] false}, 
																		   v2:{map[] false}, 
																		   v3:{map[v4:4] false},
																		   v4:{map[] false}]}
		or 
		g4 := Graph{M:map[string]Node{}} -> {map[]} equivalent to CreateGraph()
Adding Vertex: We can first add vertices and then add edges or we can directly add Edges using
			   AddEdges() method, both work the same way 
Syntax:	g.AddVertices(vertices ...string) -> Adds Vertex 'v' to the graph
	Ex: g = CreateGraph()
		g.AddVertices() -> {map[]}
		g.AddVertices("v1") -> {map[v1:{map[] false}]}
		g.AddVertices("v2", "v3") -> {map[v1:{map[] false},
										  v2:{map[] false},
										  v3:{map[] false}]}
Adding Edges: We can directly add edges using AddEdges() method. This method adds both vertices and
			  edges directly to graph thus removing the need of adding vertices to graph initially
Syntax: g.AddEdges(edges ...Edge) {...}
		'edges' is a variadic parameter of type 'Edge'. So it takes any number of aruguments (>= 0)
	Ex: g = CreateGraph() -> {map[]}
		g.AddEdges() -> doesn't do anything -> {map[]}
		g.AddEdges(Edge{"v1", "v2", 4}) -> {map[v1:{map[v2:4] false} v2:{map[] false}]}
		g.AddEdges(Edge{"v2", "v3", 4}, Edge{"v3", "v4", 4}) -> {map[v1:{map[v2:4] false}, 
																     v2:{map[v3:4] false}, 
																     v3:{map[v4:4] false},
																     v4:{map[] false}]}
Deleting Edges: We can delete edges using DeleteEdges() method. This method deletes the given edges
Syntax: g.DeleteEdges(edges ...Edge) {...}
		'edges' is a variadic parameter of type 'Edge'. So it takes any number of aruguments (>= 0)
	Ex: g = CreateGraph() -> {map[]}
		g.AddEdges() -> doesn't do anything -> {map[]}
		g.AddEdges(Edge{"v1", "v2", 4}) -> {map[v1:{map[v2:4] false} v2:{map[] false}]}
		g.AddEdges(Edge{"v2", "v3", 4}, Edge{"v3", "v4", 4}) -> {map[v1:{map[v2:4] false}, 
																     v2:{map[v3:4] false}, 
																     v3:{map[v4:4] false},
																	 v4:{map[] false}]}
		g.DeleteEdges(Edge{"v3", "v4", 4}) -> 	{map[v1:{map[v2:4] false}, 
												     v2:{map[v3:4] false}, 
												     v3:{map[] false},
													 v4:{map[] false}]}
  Note: Here one important thing is to notice that after deleting the edge ("v3", "v4"), even though we
		don't have any incoming/outgoing reference/edge to vertex "v4" in graph (which means graph 
		no more requires "v4" to be in it really), we haven't still deleted this vertex "v4". So this
		is not so optimized implementation. One reason for this is to delete this vertex "v4" from graph
		permanently, we have to go through each key in graph and their respective neighbours to ensure
		that this vertex "v4" has any reference incoming edge to it. To verify this actually increases
		the complexity rather than just leaving it undeleted in graph. Because anyways this will not
		create any difference except that whenever we are iterating over keys, it will check this "v4"
		and skips the iteration as it has no neighbours in it. Leaving it undeleted also increases 
		memory storage as this never going to be used vertex still lies in the graph. So if you have large 
		no.of operations to be performed repeatedly on this graph you can consider deleting by verifying
		with all the keys and neighbours for any reference/incoming edge to this key before deleting it. 									 

Show Graph:
Syntax: g.Show() -> Prints graph g
	Ex: g = CreateGraph(Edge{"v1", "v2", 4})
		g.Show() -> {map[v1:{map[v2:4] false} v2:{map[] false}]}
Get Weight of an edgE:
Syntax: g.GetWeight(key, neighbour string) int {...} 
		returns weight of the edge from start node 'key' to end node 'neighbour'
	Ex: g = CreateGraph(Edge{"v1", "v2", 4})
		fmt.Println(graph.GetWeight("v1", "v2")) -> returns 4
Get Edges of a nodE:
Syntax: g.GetEdges(key string) []Edge {...}
		returns all the edges of the passed key node as a slice
	Ex: g = CreateGraph(Edge{"v1", "v2", 4})
	    fmt.Println(graph.GetEdges("v1")) -> [Edge{"v1", "v2", 4}]
Get all the connected neighbours of a nodE:
Syntax: g.GetNeighbours(key string) []string {...}
	Ex: g = CreateGraph(Edge{"v1", "v2", 4})
		fmt.Println(graph.GetNeighbours("v1")) -> ["v2"]
Get all edges in a graph:
Syntax: g.GetAllEdges() []Edge {...}
	Ex: g = CreateGraph(Edge{"v1", "v2", 4}, Edge{"v3", "v4", 4})
		fmt.Println(graph.GetAllEdges()) -> [Edge{"v1", "v2", 4}, Edge{"v3", "v4", 4}]
Get all vertices in a graph:
Syntax: g.GetAllVertices() []string {...}
	Ex: g = CreateGraph(Edge{"v1", "v2", 4}, Edge{"v3", "v4", 4})
		fmt.Println(graph.GetAllVertices()) -> ["v1", "v2", "v3", "v4"]
Deep Copy of a graph:
Syntax: g.Copy() Graph {...} -> returns a deep copy of the graph
	Ex: g = CreateGraph(Edge{"v1", "v2", 4})
		deepCopy = g.Copy() -> If you modify graph 'g' now, deepCopy will not be modified
		g.AddVertex("v3") -> Adds vertex "v3" to graph 'g', This doesn't modify deepCopy and
							 still deepCopy will not has vertex "v3"
Perform Normal Copy of a graph:
Syntax: g.GetGraph() Graph {...} -> returns a normal copy of the graph
	Ex: g = CreateGraph(Edge{"v1", "v2", 4})
		normalCopy = g.GetGraph() -> If you modify graph 'g' now, normalCopy will also be modified
		g.AddVertex("v3") -> Adds vertex "v3" to graph 'g', This also modify normalCopy and
							 now normalCopy will also have vertex "v3"
Set Explored E in Node type to true or false:
Syntax: g.SetE(key string, explored bool) {...}
	Ex: g = CreateGraph(Edge{"v1", "v2", 4}) -> {map[v1:{map[v2:4] false} v2:{map[] false}]}
		g.SetE("v1", true) -> {map[v1:{map[v2:4] true} v2:{map[] false}]}
    If we directly try to assign a struct field in a map, it will lead to an error;
    For example, g.M["v1"].E = true -> Error: cannot assign to struct field g.M["v1"].E in map
    So, We can only change explored E in Node type to true or false by using SetE() method 

Make all the Vertices the Graph Unvisited:
Syntax: g.MakeUnvisited() {...}
	Ex: If graph g = {map[v1:{map[v2:4] true}, 
					      v2:{map[v3:4] true}, 
					      v3:{map[v4:4] true},
						  v4:{map[] true}]}
	  	g.MakeUnivisited() -> g = {map[v1:{map[v2:4] false}, 
								       v2:{map[v3:4] false}, 
								       v3:{map[v4:4] false},
									   v4:{map[] false}]}
*/




package graph

import (
	"fmt"
)

type Edge struct {
	K string // K for key
	N string // N for neighbour
	W int // W for weight
}

type Node struct {
	N map[string]int // N for neighbours
	E bool // E for explored
}


type Graph struct {
	M map[string]Node // M for map
}


func CreateGraph(e ...Edge) Graph {
	g := Graph{M:map[string]Node{}}
	g.AddEdges(e...)
	return g
}

func (g *Graph) AddVertices(vertices ...string) {
	for _, v := range vertices {
		if _, ok := g.M[v]; !ok{
			g.M[v] = Node{N:map[string]int{}}
		}
	}
}


func (g *Graph) AddEdges(edges ...Edge) {
	for _, edge := range edges {
		if _, ok := g.M[edge.K]; !ok {
			g.M[edge.K] = Node{N:map[string]int{}}
		}
		if _, ok := g.M[edge.N]; !ok {
			g.M[edge.N] = Node{N:map[string]int{}}
		}
		g.M[edge.K].N[edge.N] = edge.W
	}
}
 
func(g *Graph) DeleteEdges(edges ...Edge) {
	for _, edge :=  range edges {
		delete(g.M[edge.K].N, edge.N)
	}
}

func (g *Graph) Show() {
	fmt.Println(*g)
}

func (g *Graph) GetWeight(key, neighbour string) int {
	if _, ok := g.M[key]; ok {
		if _, ok := g.M[key].N[neighbour]; ok {
			return g.M[key].N[neighbour]
		}
	}
	return 0
}

func (g *Graph) GetEdges(key string) []Edge {
	var edges []Edge
	if _, ok := g.M[key]; ok {
		for neighbour, weight := range g.M[key].N{
			edges = append(edges, Edge{key, neighbour, weight})
		}
	}
	return edges
}


func (g *Graph) GetNeighbours(key string) []string {
	var neighbours []string
	for item := range g.M[key].N {
		neighbours = append(neighbours, item)
	}
	return neighbours
}

func (g *Graph) GetAllEdges() []Edge {
	var edges []Edge
	for key := range g.M {
		for neighbour, weight := range g.M[key].N{
			edges = append(edges, Edge{key, neighbour, weight})
		}
	}
	return edges
}

func (g *Graph) GetAllVertices() []string {
	var vertices []string
	for key := range g.M {
		vertices = append(vertices, key)
	}
	return vertices
}

func (g *Graph) Copy() Graph {
	deepCopy := Graph{map[string]Node{}}
	for key := range g.M {
		deepCopy.M[key] = Node{N:map[string]int{}, E:g.M[key].E}
		for neighbour, weight := range g.M[key].N {
			deepCopy.M[key].N[neighbour] = weight
		}
	}
	return deepCopy
}

func (g *Graph) GetGraph() Graph {
	return *g
}

func (g *Graph) SetE(key string, explored bool) {
	if temp, ok := g.M[key]; ok {
		temp.E = explored
		g.M[key] = temp
	}
}

func (g *Graph) MakeUnvisited() {
	for key := range g.M {
		g.SetE(key, false)
	}
}