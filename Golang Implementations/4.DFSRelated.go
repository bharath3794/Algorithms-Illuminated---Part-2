package main

import (
	"fmt"
	"mypackages/graph"
)


type fvalue struct {
	s string
	i int
}

func main() {
	// Undirected graph (Example from Page No. 52)
	edges := []graph.Edge{graph.Edge{K:"s", N:"a"}, graph.Edge{K:"s", N:"b"}, graph.Edge{K:"a", N:"c"}, 
						  graph.Edge{K:"b", N:"c"}, graph.Edge{K:"c", N:"d"}, graph.Edge{K:"c", N:"e"}, 
						  graph.Edge{K:"b", N:"s"}, graph.Edge{K:"a", N:"s"}, graph.Edge{K:"d", N:"e"}, 
						  graph.Edge{K:"b", N:"d"}, graph.Edge{K:"c", N:"a"}, graph.Edge{K:"c", N:"b"}, 
						  graph.Edge{K:"d", N:"c"}, graph.Edge{K:"e", N:"c"}, graph.Edge{K:"e", N:"d"}, 
						  graph.Edge{K:"d", N:"b"}}
	// Above example converted to directed graph (Above Undirected graph Example from Page No. 52)
	edges = []graph.Edge{graph.Edge{K:"s", N:"a"}, graph.Edge{K:"s", N:"b"}, graph.Edge{K:"a", N:"c"}, 
						 graph.Edge{K:"b", N:"c"}, graph.Edge{K:"c", N:"d"}, graph.Edge{K:"c", N:"e"}, 
						 graph.Edge{K:"d", N:"e"}, graph.Edge{K:"b", N:"d"}}
	
// depthFristSearchIterative() - Iterative DFS
	g1 := graph.CreateGraph(edges...)
	g2 := g1.Copy()
	fmt.Println("-----Iterative DFS-----")
	path := depthFristSearchIterative(g1, "s")
	fmt.Println("path =", path)
// depthFristSearchRecursive() - Recursive DFS
	fmt.Println("-----Recursive DFS-----")
	path = depthFristSearchRecursive(g2, "s")
	fmt.Println("path =", path)
}


// Depth First Search (Iterative Version) (Page No. 54)
// Algorithm Complexity = [ 𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠 + 𝐸𝑑𝑔𝑒𝑠) ]
func depthFristSearchIterative(g graph.Graph, key string) []string {
	var path []string
	stack := []string{key}
	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !g.M[v].E {
			path = append(path, v)
			g.SetE(v, true)
			for item := range g.M[v].N {
				if !g.M[item].E {
					stack = append(stack, item)
				}
			}
		}
	}
	return path
}

// Depth First Search (Recursive Version) (Page No. 55)
// Algorithm Complexity = [ 𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠 + 𝐸𝑑𝑔𝑒𝑠) ]
func depthFristSearchRecursive(g graph.Graph, key string) []string {
	if g.M[key].E {
		return []string{key}
	}
	var path []string
	if !g.M[key].E {
		path = append(path, key)
		g.SetE(key, true)
		for item := range g.M[key].N {
			if !g.M[item].E {
				lst := depthFristSearchRecursive(g, item)
				path = append(path, lst...)
			}
		}
	}
	return path
}