package main

import "fmt"

type Node struct{
  value int
  next *Node
}

type Queue struct{
  head *Node
  tail *Node
}

func NewQueue() Queue{
  return Queue{nil, nil}
}

func (self *Queue) Push(value int){
  newItem := Node{value, nil}
  if self.tail == nil{
    self.head = &newItem
    self.tail = &newItem
  } else{
    self.tail.next = &newItem
    self.tail = &newItem
  }
}

func (self *Queue) Pop() (int, error) {
  if self.head == nil{
    return -1, fmt.Errorf("empty")
  }
  del := self.head
  self.head = del.next
  return del.value, nil
}

func (self *Queue) Print(){
  if self.head == nil{
    fmt.Println("empty")
    return
  }
  curr := self.head
  for curr != nil{
    fmt.Printf("%d ->", curr.value)
    curr = curr.next
  }
  fmt.Println()
}

func callPop(q Queue) Queue {
  v, err := q.Pop()
  if err != nil{
    fmt.Println(err)
  }
  fmt.Printf("popped: %d\n", v)
  return q
}

func main(){
  q := NewQueue()
  q.Print()
  q.Push(1)
  q.Print()
  q.Push(2)
  q.Print()
  q.Push(3)
  q.Print()
  q = callPop(q) 
  q.Print()
  q = callPop(q) 
  q.Print()
  q = callPop(q)
  q.Print()
}
