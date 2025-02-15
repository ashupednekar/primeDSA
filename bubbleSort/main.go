package main

import "fmt"


func bubbleSort(l []int) ([]int, error){
  for i:=0; i<len(l); i++{
    for j:=0; j<len(l)-1-i; j++{
      if l[j] > l[j+1]{
        tmp := l[j]
        l[j] = l[j+1]
        l[j+1] = tmp
      }
    }
  }
  return l, nil
}


func main(){
  l := []int{1,3,7,4,2}
  res, err := bubbleSort(l)
  if err != nil{
    fmt.Errorf("%s", err)
  }
  fmt.Printf("res: %v\n", res)
}
