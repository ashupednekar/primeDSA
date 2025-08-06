package main

import (
	"fmt"
)

type Adj struct {
	to     int
	weight int
}


func walk(graph [][]Adj, curr int, needle int, seen []bool) []int{
	if curr == needle{
		return []int{curr}
	}
	if seen[curr]{
		return nil
	}
	seen[curr] = true
	for _, edge := range(graph[curr]){
		if path := walk(graph, edge.to, needle, seen); path != nil{
			return append([]int{curr}, path...)
		}
	}
	return nil
} 


func searchDFS(graph [][]Adj, source int, needle int) []int{
	seen := make([]bool, len(graph))
	return walk(graph, source, needle, seen)
}

func main() {
	graph := [][]Adj{
		{{to: 2, weight: 1}, {to: 3, weight: 1}},
		{},
		{{to: 3, weight: 1}},
		{{to: 4, weight: 1}, {to: 1, weight: 1}},
		{{to: 2, weight: 1}},
	}
	fmt.Printf("path: %v\n", searchDFS(graph, 0, 4))
}
