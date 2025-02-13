package main

import "fmt"

func search(l []int, v int) (int, error) {
  for i:=0; i<len(l); i++{
    if l[i] == v{
      return i, nil
    }
  } 
  return 0, fmt.Errorf("not found")
}


func main(){

  l := []int{1,2,3,4,5,6,7,8}
  res, err := search(l, 5)
  if err != nil{
    fmt.Errorf("error: %v", err)
  }
  fmt.Println("result: %v", res)
  

}
