package main

import "fmt"

type Node struct{
	Value int
	Left *Node
	Right *Node
}

func BFSExistsCheckChans(node *Node, needle int) bool{
	q := make(chan *Node, 10)
	q <- node
	curr := <-q
	if curr.Value == needle{
		return true
	}
	if curr.Left != nil{
		q <- curr.Left
	}
	if curr.Right != nil{
		q <- curr.Right
	}
	return false
}

func BFSExistsCheck(node *Node, needle int) bool{
	q := []*Node{node}
	curr := q[0]
	q = q[1:]
	if curr.Value == needle{
		return false
	}
	if curr.Left != nil{
		q = append(q, curr.Left)
	}
	if curr.Right != nil{
		q = append(q, curr.Right)
	}
	return false
}

func main() {
	root := Node{
		Value: 34,
		Left: &Node{
			Value: 56,
			Left: &Node{
				Value: 23,
				Left: &Node{
					Value: 36,
					Right: &Node{
						Value: 8,
					},
				},
			},
		},
		Right: &Node{
			Value: 83,
			Left: &Node{
				Value: 4,
			},
		},
	}
	fmt.Printf("given tree: %v\n", root)
	fmt.Printf("num: %d exists in tree: %v", 23, BFSExistsCheckChans(&root, 23))
	fmt.Printf("num: %d exists in tree: %v\n\n", 87, BFSExistsCheckChans(&root, 87))
  fmt.Printf("num: %d exists in tree: %v", 23, BFSExistsCheck(&root, 23))
	fmt.Printf("num: %d exists in tree: %v", 87, BFSExistsCheck(&root, 87))

}

