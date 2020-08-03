package main

import (
	"fmt"
	"mypackages/graph"
)


func main() {
	edges := []graph.Edge{graph.Edge{K:"s", N:"a"}, graph.Edge{K:"a", N:"c"}, graph.Edge{K:"b", N:"d"}, 
	         			  graph.Edge{K:"s", N:"b"}, graph.Edge{K:"b", N:"c"}, graph.Edge{K:"c", N:"e"}, 
	         			  graph.Edge{K:"c", N:"d"}, graph.Edge{K:"d", N:"e"}}
	
// breadthFristSearch()
	g1 := graph.CreateGraph(edges...)
	g2 := g1.Copy()
	path := breadthFristSearch(g1, "s")
	fmt.Println("path =", path)

// augmentedBFS()
	steps := augmentedBFS(g2, "s")
	fmt.Println("No.of minimum steps you need to at least take to reach final node from starting node is", steps)
}

// Breadth First Search (From Page No. 39)
// Algorithm Complexity = [ 𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠 + 𝐸𝑑𝑔𝑒𝑠) ]
func breadthFristSearch(g graph.Graph, key string) []string {
	g.SetE(key, true)
	queue := []string{key}
	for i:=0; i<len(queue); i++ {
		for neighbour := range g.M[queue[i]].N{
			if !g.M[neighbour].E {
				g.SetE(neighbour, true)
				queue = append(queue, neighbour)
			}
		}
	}
	return queue
}

// Augmented Breadth First Search (From Page No. 44)
// Algorithm Complexity = [ 𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠 + 𝐸𝑑𝑔𝑒𝑠) ]
func augmentedBFS(g graph.Graph, key string) int {
	steps := map[string]int{}
	steps[key] = 0
	g.SetE(key, true)
	queue := []string{key}
	for i:=0; i<len(queue); i++ {
		for neighbour := range g.M[queue[i]].N{
			if !g.M[neighbour].E {
				g.SetE(neighbour, true)
				steps[neighbour] = steps[queue[i]] + 1
				queue = append(queue, neighbour)
			}
		}
	}
	return steps[queue[len(queue)-1]]
}