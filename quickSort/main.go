package main

import "fmt"

func qs(arr []int, low int, high int){
	if low < high{
		pivotIdx := partition(arr, low, high)
		qs(arr, low, pivotIdx - 1)
		qs(arr, pivotIdx + 1, high)
	}
}

func partition(arr []int, low int, high int) int {
	idx := low - 1
	for i:=low; i<high; i++{
		if arr[i] <= arr[high]{
			idx++
			arr[idx], arr[i] = arr[i], arr[idx]
		}
	}
	idx++
	arr[idx], arr[high] = arr[high], arr[idx]
	return idx
}

func main(){
	arr := []int{3,5,5,2,4,6,8,3,6,4,6,7,4,67,6,43,6,8,2,3,4,5}
	fmt.Printf("given arr: %v\n", arr)
	qs(arr, 0, len(arr) - 1)
	fmt.Printf("sorted arr: %v\n", arr)
}
