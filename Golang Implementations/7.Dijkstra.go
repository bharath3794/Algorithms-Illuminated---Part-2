package main

import (
	"fmt"
	"mypackages/graph"
	"math"
	"container/heap"
)


func main() {
	// Directed Acyclic Graph (Example from Pg. No. 94)
	edges := []graph.Edge{graph.Edge{K:"s", N:"v", W:1}, graph.Edge{K:"s", N:"w", W:4}, 
						  graph.Edge{K:"v", N:"w", W:2}, graph.Edge{K:"v", N:"t", W:6}, 
						  graph.Edge{K:"w", N:"t", W:3}}
	
	g1 := graph.CreateGraph(edges...)
	// Normal Dijkstra Algorithm (From Page No. 92)
	fmt.Println("-----Dijkstra Algorithm for Shortest Path-----")
	shortestDistances := dijkstra(g1, "s")
	fmt.Println("shortestDistances =", shortestDistances)
	// Dijkstra Algorithm using Heaps (Explanation From Pg. No. 121-123)
	fmt.Println("-----Dijkstra Algorithm Using Heaps-----")
	shortestDistances = dijkstraHeapBased(g1, "s")
	fmt.Println("shortestDistances =", shortestDistances)

	//One other example
	edges = []graph.Edge{graph.Edge{K:"a", N:"b", W:7}, graph.Edge{K:"a", N:"e", W:14}, 
						 graph.Edge{K:"a", N:"f", W:9}, graph.Edge{K:"b", N:"f", W:8}, 
						 graph.Edge{K:"b", N:"c", W:15}, graph.Edge{K:"f", N:"e", W:2}, 
						 graph.Edge{K:"f", N:"c", W:6}, graph.Edge{K:"c", N:"d", W:4}, 
						 graph.Edge{K:"e", N:"d", W:9}}
	g2 := graph.CreateGraph(edges...)
	// Normal Dijkstra Algorithm (From Page No. 92)
	fmt.Println("-----Dijkstra Algorithm for Shortest Path-----")
	shortestDistances = dijkstra(g2, "a")
	fmt.Println("shortestDistances =", shortestDistances)
	// Dijkstra Algorithm using Heaps (Explanation From Pg. No. 121-123)
	fmt.Println("-----Dijkstra Algorithm Using Heaps-----")
	shortestDistances = dijkstraHeapBased(g2, "a")
	fmt.Println("shortestDistances =", shortestDistances)
}

/* 
* Dijkstra Algorithm for computing shortest path (From Page No. 92)
* This algorithm works on any directed graph with a source vertex; (source vertex is the vertex
  with no incoming edges);
* The one other requirement for this algorithm to be working is that the length of the edges shouldn't be
  negative

* Algorithm Complexity = [ 𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠 ∗ 𝐸𝑑𝑔𝑒𝑠) ]
*/
func dijkstra(g graph.Graph, source string) map[string]int {
	xMap := map[string]int{source:0}
	for len(xMap) < len(g.M) {
		var ele string
		minDist := math.MaxInt64
		for key := range xMap {
			for neighbour := range g.M[key].N {
				if _, ok := xMap[neighbour]; ok{
					continue
				}
				curDist := xMap[key] + g.M[key].N[neighbour]
				if curDist < minDist {
					minDist = curDist
					ele = neighbour
				}
			}
		}
		xMap[ele] = minDist
	}
	return xMap
}

/*
* Dijkstra Algorithm using Heaps (Explanation From Pg. No. 121-123)
* Based on Book: [ 𝑂((𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠+𝐸𝑑𝑔𝑒𝑠)∗𝑙𝑜𝑔(𝑛)) ];
* Below Implementation: [ 𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠∗𝐸𝑑𝑔𝑒𝑠∗𝑙𝑜𝑔(𝑛)) ];
* While we delete the element in Heap (see Pg. No. 123, Dijkstra (Heap-Based, Part 2)), 
  in the below code we always need to search for the element first in the heap to know its index and 
  then we would delete that element based on the index. If we don't search for key each time deleting the 
  key, the complexity would normally be the mentioned complexity of [𝑂((𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠+𝐸𝑑𝑔𝑒𝑠)∗𝑙𝑜𝑔(𝑛))]. 
  As we are searching each time when we want to delete the key, it adds up to the complexity and the 
  complexity would be [𝑂(𝑉𝑒𝑟𝑡𝑖𝑐𝑒𝑠*𝐸𝑑𝑔𝑒𝑠∗𝑙𝑜𝑔(𝑛))]
*/
func dijkstraHeapBased(g graph.Graph, source string) map[string]int {
	xMap := map[string]int{}
	heapSlice := &heapVertices{}
	for vert := range g.M {
		var val int
		if vert == source {
			val = 0
		} else {
			val = math.MaxInt64
		}
		heap.Push(heapSlice, &vertex{vert, val})
	}
	for len(*heapSlice) > 0 {
		popped := heap.Pop(heapSlice)
		w := popped.(*vertex)
		xMap[w.s] = w.i
		if len(*heapSlice) > 0 {
			for neighbour := range g.M[w.s].N {
				dist := xMap[w.s] + g.M[w.s].N[neighbour]
				// This update() method will perform operations from line 13-15 in Dijkstra (Heap-Based, Part 2)
				// (see Pg. No. 123)
				heapSlice.update(neighbour, dist)
			}
		}
	}
	return xMap
}


/* Dependency for dijkstraHeapBased() to delete or to update a term in heap */
type vertex struct {
	s string
	i int
}

type heapVertices []*vertex

func (h heapVertices) Len() int { return len(h) }
func (h heapVertices) Less(i, j int) bool { return h[i].i < h[j].i }
func (h heapVertices) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *heapVertices) Push(x interface{}) {
	*h = append(*h, x.(*vertex))
}

func (h *heapVertices) Pop() interface{} {
	popped := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return popped
}

// This update() method will perform operations from line 13-15 in Dijkstra (Heap-Based, Part 2)
// (see Pg. No. 123)
func (h *heapVertices) update(key string, length int) {
	var idx int
	for j, item := range *h {
		if (*item).s == key {
			length = int(math.Min(float64((*item).i), float64(length)))
			(*item).i = length
			idx = j
			break
		}
 	}

 	heap.Fix(h, idx)
}