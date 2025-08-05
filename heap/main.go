package main

import "fmt"

type MinHeap struct{
	data [50]int
	length int
}

func (heap *MinHeap) parent(idx int) int{
	return (idx - 1) / 2
}

func (heap *MinHeap) left(idx int) int{
	return 2*idx + 1
}

func (heap *MinHeap) right(idx int) int {
	return heap.left(idx) + 1
}

func (heap *MinHeap) heapifyUp(idx int) {
	if idx == 0{
		return
	}
	parentIdx := heap.parent(idx)
	if heap.data[parentIdx] > heap.data[idx]{
		heap.data[parentIdx], heap.data[idx] = heap.data[idx], heap.data[parentIdx]
		heap.heapifyUp(parentIdx)
	}
}

func (heap *MinHeap) heapifyDown(idx int){
	smallest := idx
	leftIdx := heap.left(idx) 
	rightIdx := heap.right(idx)
	if idx >= heap.length{
		return
	}
	if leftIdx < heap.length && heap.data[leftIdx] < heap.data[smallest]{
		smallest = leftIdx
	}
	if rightIdx < heap.length && heap.data[rightIdx] < heap.data[smallest]{
		smallest = rightIdx
	}
	if idx != smallest{
		heap.data[idx], heap.data[smallest] = heap.data[smallest], heap.data[idx]
	  heap.heapifyDown(smallest)
	}
}

func (heap *MinHeap) Push(num int){
	heap.data[heap.length] = num
	heap.heapifyUp(heap.length)
	heap.length++
}

func (heap *MinHeap) Pop() int {
	heap.length--
	if heap.length == 0{
		return -1
	}
	popped := heap.data[0]
	if heap.length == 1{
		heap.data = [50]int{}
		return popped
	}
	heap.data[0] = heap.data[heap.length]
	heap.data[heap.length] = 0
	heap.heapifyDown(0)
	return popped
}

func main(){
	heap := MinHeap{}
	heap.Push(2)
	heap.Push(5)
	heap.Push(9)
	heap.Push(7)
	heap.Push(3)
	heap.Push(3)
	heap.Push(5)
	heap.Push(7)
	heap.Push(6)
	heap.Push(4)
	fmt.Printf("heap: %v\n", heap)
	fmt.Printf("popped: %d, %v\n", heap.Pop(), heap)
	fmt.Printf("popped: %d, %v\n", heap.Pop(), heap)
	fmt.Printf("popped: %d, %v\n", heap.Pop(), heap)
	fmt.Printf("popped: %d, %v\n", heap.Pop(), heap)
	fmt.Printf("popped: %d, %v\n", heap.Pop(), heap)
	fmt.Printf("popped: %d, %v\n", heap.Pop(), heap)
	fmt.Printf("popped: %d, %v\n", heap.Pop(), heap)
	fmt.Printf("popped: %d, %v\n", heap.Pop(), heap)
	fmt.Printf("popped: %d, %v\n", heap.Pop(), heap)
	fmt.Printf("popped: %d, %v\n", heap.Pop(), heap)
	fmt.Printf("length: %d\n", heap.length)
}
