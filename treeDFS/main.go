package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func PreOrder(node *Node, path []int) []int {
	if node == nil {
		return path
	}
	path = append(path, node.Value)
	path = PreOrder(node.Left, path)
	path = PreOrder(node.Right, path)
	return path
}

func InOrder(node *Node, path []int) []int {
	if node == nil {
		return path
	}
	path = InOrder(node.Left, path)
	path = append(path, node.Value)
	path = InOrder(node.Right, path)
	return path
}

func PostOrder(node *Node, path []int) []int {
	if node == nil {
		return path
	}
	path = PostOrder(node.Left, path)
	path = PostOrder(node.Right, path)
	path = append(path, node.Value)
	return path
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
	fmt.Printf("pre order: %v", PreOrder(&root, []int{}))
	fmt.Printf("in order: %v", InOrder(&root, []int{}))
	fmt.Printf("post order: %v", PostOrder(&root, []int{}))
}
