package main

import (
	"fmt"
	"math/rand"
	"time"
	"mypackages/graph"
)


func main() {
// Undirected graph (Example from Page No. 31)
	edges := []graph.Edge{graph.Edge{K:"s", N:"u"}, graph.Edge{K:"s", N:"v"}, graph.Edge{K:"u", N:"v"}, 
						  graph.Edge{K:"u", N:"w"}, graph.Edge{K:"v", N:"w"}, graph.Edge{K:"u", N:"s"}, 
						  graph.Edge{K:"v", N:"s"}, graph.Edge{K:"v", N:"u"}, graph.Edge{K:"w", N:"u"}, 
						  graph.Edge{K:"w", N:"v"}}
// directed graph (Example from Page No. 31)
	edges  = []graph.Edge{graph.Edge{K:"s", N:"u"}, graph.Edge{K:"s", N:"v"}, graph.Edge{K:"u", N:"v"}, 
	                     graph.Edge{K:"w", N:"u"}, graph.Edge{K:"w", N:"v"}}
	
	g := graph.CreateGraph(edges...)
	rand.Seed(time.Now().UnixNano())
	genericSearch(g, "s")

}


// Generic Search Algorithm (From Page No. 32)
// (Based on explanation with example in Page No. 33)
func genericSearch(g graph.Graph, key string) {
	path := []string{}
	choose := []string{}
	g.SetE(key, true)
	for ele := range g.M {
		if g.M[ele].E == true {
			for item := range g.M[ele].N {
				if g.M[item].E == false {
					choose = append(choose, item)
				}
			}
		}
		if len(choose) > 0 {
			idx := rand.Intn(len(choose))
			if g.M[choose[idx]].E == false {
				g.SetE(choose[idx], true)
				path = append(path, choose[idx])
			}
			choose = choose[:len(choose)-1]
		}
	}
	for len(choose) > 0 {
		idx := rand.Intn(len(choose))
		if g.M[choose[idx]].E == false {
			g.SetE(choose[idx], true)
			path = append(path, choose[idx])
		}
		choose = choose[:len(choose)-1]
	}
	for _, item := range path {
		fmt.Printf("We can reach %v from %v\n", item, key)
	}
}
