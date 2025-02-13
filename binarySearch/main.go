package main

import "fmt"


func search(l []int, v int, offset int) (int, error) {
	if len(l) == 0 {
		return -1, fmt.Errorf("value not found")
	}

	midIdx := len(l) / 2
	midVal := l[midIdx]

	if midVal > v {
		return search(l[:midIdx], v, offset) 
	} else if midVal < v {
		return search(l[midIdx+1:], v, offset+midIdx+1) 
	} else {
		return offset + midIdx, nil 
	}
}

func searchIterative(l []int, v int) (int, error){
  low := 0
  high := len(l)
  for low < high{
    mid := (low + high)/2
    if l[mid] == v{
      return mid, nil
    }else if l[mid] < v{
      low = mid + 1
    }else{
      high = mid
    }
  }
  return -1, fmt.Errorf("not found")
}

func main(){

  l := []int{1,2,3,4,5,6,7,8}
  res, err := search(l, 2, 0)
  if err != nil{
    fmt.Errorf("error: %v", err)
  }
  fmt.Printf("result: %v\n", res)

  res, err = searchIterative(l, 2)
  if err != nil{
    fmt.Errorf("error: %v", err)
  }
  fmt.Printf("result: %v\n", res)

}
