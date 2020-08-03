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
	// Directed Acyclic Graph with 4 strongly connected components which are themselves directed cyclic graphs
// (Example from Page No. 67)
	edges := []graph.Edge{graph.Edge{K:"v1", N:"v3"}, graph.Edge{K:"v3", N:"v5"}, 
						  graph.Edge{K:"v5", N:"v1"}, graph.Edge{K:"v3", N:"v11"}, 
						  graph.Edge{K:"v11", N:"v8"}, graph.Edge{K:"v11", N:"v6"}, 
						  graph.Edge{K:"v6", N:"v10"}, graph.Edge{K:"v10", N:"v8"}, 
						  graph.Edge{K:"v8", N:"v6"}, graph.Edge{K:"v5", N:"v9"}, 
						  graph.Edge{K:"v5", N:"v7"}, graph.Edge{K:"v9", N:"v2"}, 
						  graph.Edge{K:"v2", N:"v4"}, graph.Edge{K:"v4", N:"v7"}, 
						  graph.Edge{K:"v7", N:"v9"}, graph.Edge{K:"v9", N:"v4"}, 
						  graph.Edge{K:"v9", N:"v8"}, graph.Edge{K:"v2", N:"v10"}}
 	kg1 := graph.CreateGraph(edges...)
 	scc := kosaraju(kg1)
 	fmt.Println("Strongly Connected Components =", scc)
 // Directed Acyclic Graph (converted from undirected graph in Example from Page No. 52)
 	edges = []graph.Edge{graph.Edge{K:"s", N:"a"}, graph.Edge{K:"s", N:"b"}, graph.Edge{K:"a", N:"c"}, 
 						 graph.Edge{K:"b", N:"c"}, graph.Edge{K:"c", N:"d"}, graph.Edge{K:"c", N:"e"}, 
 						 graph.Edge{K:"d", N:"e"}, graph.Edge{K:"b", N:"d"}}
 	kg2 := graph.CreateGraph(edges...)
 	scc = kosaraju(kg2)
 	fmt.Println("Strongly Connected Components =", scc)
 // Directed Cyclic Graph (with s -> a -> c -> b -> s)
 	edges = []graph.Edge{graph.Edge{K:"s", N:"a"}, graph.Edge{K:"a", N:"c"}, graph.Edge{K:"c", N:"b"}, 
 						 graph.Edge{K:"b", N:"s"}}
 	kg3 := graph.CreateGraph(edges...)
 	scc = kosaraju(kg3)
 	fmt.Println("Strongly Connected Components =", scc)
}



// Strongly Connected Components for Directed Graph
// (Kosaraju Algorithm from Page No. 74) (Depends on topologicalSort() function from above steps)
// Algorithm Complexity = [ 𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠 + 𝐸𝑑𝑔𝑒𝑠) ]
func kosaraju(g graph.Graph) [][]string {
	revG := graph.CreateGraph()
	reverseGraph := func() {
		for vertex := range g.M {
			for neighbour := range g.M[vertex].N {
				if _, ok := revG.M[neighbour]; !ok {
					revG.M[neighbour] = graph.Node{N:map[string]int{vertex:g.M[vertex].N[neighbour]},
												   E:g.M[neighbour].E}
					continue
				}
				revG.M[neighbour].N[vertex] = g.M[vertex].N[neighbour]
			}
		}
	}
	var scc [][]string
	reverseGraph()
	topoOrder := topologicalSort(revG)
	var numSCC int
	var dfsSCC func(ele string)
	dfsSCC = func(ele string) {
		if g.M[ele].E {
			return
		}
		g.SetE(ele, true)
		if len(scc)-1 < numSCC-1{
			scc = append(scc, []string{ele})
		} else {
			scc[numSCC-1] = append(scc[numSCC-1], ele)
		}
		for item := range g.M[ele].N {
			if !g.M[item].E {
				dfsSCC(item)
			}
		}
	}
	for i:=len(topoOrder)-1; i>=0; i-- {
		key := topoOrder[i].s
		if !g.M[key].E {
			numSCC++
			dfsSCC(key)
		}
	}
	return scc
}


// Dependency: Topological Ordering (Topological Sort) using Depth-First Search [From Page No. 62]
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
