package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func compareBT(treeA *Node, treeB *Node) bool {
	if treeA == nil && treeB == nil{
		return true
	}
	if treeA == nil || treeB == nil{
		return false
	}
	if treeA.Value != treeB.Value{
		return false
	}
	return compareBT(treeA.Left, treeB.Left) && compareBT(treeA.Right, treeB.Right)
}

func main() {
	treeA := Node{
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
	treeB := Node{
		Value: 34,
		Left: &Node{
			Value: 56,
			Left: &Node{
				Value: 23,
				Left: &Node{
					Value: 36,
					Right: nil,
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
	fmt.Println("treeA == treeA: %v", compareBT(&treeA, &treeA))
	fmt.Println("treeA == treeB: %v", compareBT(&treeA, &treeB))
}
