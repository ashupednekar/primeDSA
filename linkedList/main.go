package main

import "fmt"


type ListNode struct{
  next *ListNode
  value int
}

type DoublyListNode struct{
  next *DoublyListNode
  prev *DoublyListNode
  value int
}


type LinkedList struct{
  head *ListNode
}

type DoublyLinkedList struct{
  head *DoublyListNode
  tail *DoublyListNode
}

func NewLinkedList() LinkedList{
  return LinkedList{nil}
}

func NewDoublyLinkedList() DoublyLinkedList{
  return DoublyLinkedList{nil, nil}
}


func (self *LinkedList) Push(value int){
  newItem := ListNode{nil, value}
  if self.head == nil{
    self.head = &newItem
  }else{
    curr := self.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = &newItem
  }
}

func (self *DoublyLinkedList) Push(value int){
  newItem := DoublyListNode{nil, nil, value}
  if self.head == nil{
    self.head = &newItem
  }else{
    curr := self.head
    for curr.next != nil{
      curr = curr.next
    }
    curr.next = &newItem
  }
  newItem.prev = self.head
}


func (self *LinkedList) Print() {
	if self.head == nil {
		fmt.Println("empty")
		return
	}
	curr := self.head
	for curr != nil {
		fmt.Printf("%d", curr.value)
		if curr.next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.next
	}
	fmt.Println()
}

func (self *DoublyLinkedList) Print() {
	if self.head == nil {
		fmt.Println("empty")
		return
	}
	curr := self.head
	for curr != nil {
		fmt.Printf("%d", curr.value)
		if curr.next != nil {
			fmt.Print(" <-> ")
		}
		curr = curr.next
	}
	fmt.Println()
}

func main(){
  l := NewLinkedList()
  l.Push(2)
  l.Push(3)
  l.Push(4)
  l.Print()

  d := NewDoublyLinkedList()
  d.Push(2)
  d.Push(3)
  d.Push(4)
  d.Print()
}
