package main

import "fmt"

type Node struct{
  value int
  next *Node
}

type Stack struct{
  head *Node
  length int
}

func NewStack() Stack{
  return Stack{nil, 0}
}

func (self *Stack) Print(){
  if self.head == nil{
    fmt.Println("empty")
    return
  }
  curr := self.head
  for curr != nil{
    fmt.Printf("%d -> ", curr.value)
    curr = curr.next
  }
  fmt.Printf("[of length %d]\n", self.length)
}

func (self *Stack) Push(value int){
  newItem := Node{value, nil}
  if self.head == nil{
    self.head = &newItem
  }else{
    newItem.next = self.head
    self.head = &newItem
  }
  self.length += 1
}

func (self *Stack) Pop() (int, error){
  if self.head == nil{
    return -1, fmt.Errorf("empty")
  }else{
    del := self.head
    self.head = del.next
    return del.value, nil
  }
}

func callPop(s Stack) Stack {
  v, err := s.Pop()
  if err != nil{
    fmt.Println(err)
  }
  fmt.Printf("popped: %d\n", v)
  return s
}

func main(){
  s := NewStack()
  s.Print()
  s.Push(1)
  s.Print()
  s.Push(2)
  s.Print()
  s.Push(3)
  s.Print()
  s = callPop(s) 
  s.Print()
  s = callPop(s) 
  s.Print()
  s = callPop(s)
  s.Print()
}
