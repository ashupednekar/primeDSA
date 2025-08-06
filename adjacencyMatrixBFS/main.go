package main

import "fmt"

type Q struct{
	arr []int
}

func NewQ() Q {
	return Q{arr: []int{}}
}

func (q *Q) enqueue(num int){
	q.arr = append(q.arr, num)
}

func (q *Q) dequeue() int {
	if len(q.arr) == 0{
		return -1
	}
	out := q.arr[0]
	q.arr = q.arr[1:]
	return out
}

func searchBFS(graph [][]int, source int, needle int) []int {
	seen := make([]bool, len(graph))
	prevs := []int{}
	for range(len(graph)){ 
		prevs = append(prevs, -1)
	}

	seen[source] = true
	q := NewQ()
	q.enqueue(source)

	curr := q.dequeue()
	for curr != -1{
		tos := graph[curr]
		for to, edge := range tos{
			if edge == 0{
				continue
			}
			if seen[to] {
				continue
			}
			seen[to] = true
			prevs[to] = curr
			q.enqueue(to)
		}
		curr = q.dequeue()
	}
	out := []int{}
	curr = needle
	for curr != -1{
		out = append([]int{curr}, out...)
		curr = prevs[curr]
	}
	return out
}

func main() {
	matrix := [][]int{
		{0, 1, 1, 1, 0}, // from node 0 to 1, 2, 3
		{1, 0, 0, 0, 0}, // from node 1 to 0,2
		{0, 0, 0, 1, 0}, // from node 2 to 3
		{0, 0, 0, 0, 1}, // from node 3 to 4
		{0, 0, 0, 0, 0}, // node 4 has no outgoing edges
	}
	fmt.Printf("path: %v", searchBFS(matrix, 0, 4))
}
