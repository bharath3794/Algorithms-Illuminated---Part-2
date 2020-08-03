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
	// topologicalSort()
// Topological Sort only works on DAG (Directed Acyclic Graphs) and it 
// doesn't work on Directed Cyclic Graphs i.e. directed graph with cycles
	fmt.Println("-----Topological Sort-----")
	// Directed Acyclic Graph (Example from Quiz 8.3, Page No. 57, 58)
	edges := []graph.Edge{graph.Edge{K:"s", N:"v"}, graph.Edge{K:"s", N:"w"}, graph.Edge{K:"v", N:"t"}, 
	 					 graph.Edge{K:"w", N:"t"}}
	tg1 := graph.CreateGraph(edges...)
	fValues := topologicalSort(tg1)
	fmt.Println("fValues =", fValues)
	// Directed Cyclic Graph (Converted from above Example from Quiz 8.3, Page No. 57, 58)
	edges = []graph.Edge{graph.Edge{K:"s", N:"v"}, graph.Edge{K:"s", N:"w"}, graph.Edge{K:"v", N:"t"}, 
						 graph.Edge{K:"w", N:"t"}, graph.Edge{K:"t", N:"s"}}
	tg2 := graph.CreateGraph(edges...)
	fValues = topologicalSort(tg2)
	fmt.Println("fValues =", fValues)
	// Directed Acyclic Graph (converted from undirected graph in Example from Page No. 52)
	edges = []graph.Edge{graph.Edge{K:"s", N:"a"}, graph.Edge{K:"s", N:"b"}, graph.Edge{K:"a", N:"c"}, 
						 graph.Edge{K:"b", N:"c"}, graph.Edge{K:"c", N:"d"}, graph.Edge{K:"c", N:"e"}, 
						 graph.Edge{K:"d", N:"e"}, graph.Edge{K:"b", N:"d"}}
	tg3 := graph.CreateGraph(edges...)
	fValues = topologicalSort(tg3)
	fmt.Println("fValues =", fValues)
	// (Directed Cyclic Graph) (converted from above Example from Page No. 52)
	// Topological Sort doesn't work on Directed Cyclic Graphs i.e. directed graph with cycles
	edges = []graph.Edge{graph.Edge{K:"s", N:"a"}, graph.Edge{K:"s", N:"b"}, graph.Edge{K:"a", N:"c"}, 
						 graph.Edge{K:"b", N:"c"}, graph.Edge{K:"c", N:"d"}, graph.Edge{K:"c", N:"e"}, 
						 graph.Edge{K:"d", N:"e"}, graph.Edge{K:"b", N:"d"}, graph.Edge{K:"e", N:"s"}, 
						 graph.Edge{K:"d", N:"s"}}
	tg4 := graph.CreateGraph(edges...)
	fValues = topologicalSort(tg4)
	fmt.Println("fValues =", fValues)
	// Directed Cyclic Graph (Example from Fig. 8.11, Page No. 59)
	// Topological Sort doesn't work on Directed Cyclic Graphs i.e. directed graph with cycles
	edges = []graph.Edge{graph.Edge{K:"u", N:"v"}, graph.Edge{K:"v", N:"w"}, graph.Edge{K:"w", N:"x"}, 
						 graph.Edge{K:"x", N:"y"}, graph.Edge{K:"y", N:"z"}, graph.Edge{K:"z", N:"u"}}
	tg5 := graph.CreateGraph(edges...)
	fValues = topologicalSort(tg5)
	fmt.Println("fValues =", fValues)
}



// Topological Ordering (Topological Sort) using Depth-First Search [From Page No. 62]
// Algorithm Complexity = [ 𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠 + 𝐸𝑑𝑔𝑒𝑠) ]
func topologicalSort(g graph.Graph) []fvalue {
	var fValues []fvalue
	curLabel := len(g.M)
	var dfsTopo func(ele string)
	// Closure
	dfsTopo = func(ele string) {
		if g.M[ele].E {
			return
		}
		g.SetE(ele, true)
		for item := range g.M[ele].N {
			if !g.M[item].E {
				dfsTopo(item)
			}
		}
		fValues = append(fValues, fvalue{ele, curLabel})
		curLabel--
	}
	for key := range g.M {
		if !g.M[key].E {
			dfsTopo(key)
		}
	}
	return fValues
}
