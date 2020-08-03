package main

import (
	"fmt"
	"mypackages/graph"
)


func main() {
	//connectedComponents()
	// Undirected graph (Example from Page No. 31)
    edges := []graph.Edge{graph.Edge{K:"s", N:"u"}, graph.Edge{K:"s", N:"v"}, graph.Edge{K:"u", N:"v"}, 
    					  graph.Edge{K:"u", N:"w"}, graph.Edge{K:"v", N:"w"}, graph.Edge{K:"u", N:"s"}, 
    					  graph.Edge{K:"v", N:"s"}, graph.Edge{K:"v", N:"u"}, graph.Edge{K:"w", N:"u"}, 
    					  graph.Edge{K:"w", N:"v"}}
    // Undirected graph
    edges = []graph.Edge{graph.Edge{K:"s", N:"u"}, graph.Edge{K:"v", N:"w"}, graph.Edge{K:"u", N:"s"}, 
        				 graph.Edge{K:"w", N:"v"}}
    g3 := graph.CreateGraph(edges...)
	cc := connectedComponents(g3)
	fmt.Println("Connected Components =", cc)
}


// Connected Components for Undirected Graph (UCC Algorithm from Page No. 49)
// Algorithm Complexity = [ 𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠 + 𝐸𝑑𝑔𝑒𝑠) ]
func connectedComponents(g graph.Graph) [][]string {
	var cc [][]string
	numCC := 0
	for key := range g.M {
		if !g.M[key].E {
			numCC += 1
			g.SetE(key, true)
			queue := []string{key}
			for i:=0; i<len(queue); i++ {
				if len(cc)-1 < numCC-1{
					cc = append(cc, []string{queue[i]})
				} else {
					cc[numCC-1] = append(cc[numCC-1], queue[i])
				}
				for neighbour := range g.M[queue[i]].N{
					if !g.M[neighbour].E {
						g.SetE(neighbour, true)
						queue = append(queue, neighbour)
					}
				}
			}
		}
	}
	return cc
}