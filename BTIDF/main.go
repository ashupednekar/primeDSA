package main

import "fmt"

type Node struct {
	Value  int
	Height int
	Parent *Node
	Left   *Node
	Right  *Node
}


func printTree(node *Node, prefix string, isLeft bool) {
	if node == nil {
		return
	}
	printTree(node.Right, prefix+"│   ", false)
	if isLeft {
		fmt.Println(prefix + "└── " + fmt.Sprintf("%d(%d)", node.Value, node.Height))
	} else {
		fmt.Println(prefix + "┌── " + fmt.Sprintf("%d(%d)", node.Value, node.Height))
	}
	printTree(node.Left, prefix+"│   ", true)
}

func (node *Node) insert(num int) *Node {
	if node == nil {
		return &Node{num, 1, nil, nil, nil}
	}
	if num <= node.Value {
		node.Left = node.Left.insert(num)
		node.Left.Parent = node
	} else {
		node.Right = node.Right.insert(num)
		node.Right.Parent = node
	}
	node.Height = max(node.Left.getHeight(), node.Right.getHeight()) + 1
	return node
}

func (node *Node) getHeight() int {
	if node == nil{
		return 0
	}
	return node.Height
}

func (node *Node) find(num int) bool{
	if node == nil{
		return false
	}
	if num == node.Value{
		return true
	}
	if num < node.Value{
		return node.Left.find(num)
	}else{
		return node.Right.find(num)
	}
}

func (node *Node) findNode(num int) *Node{
	if node == nil{
		return nil 
	}
	if num == node.Value{
		return node 
	}
	if num < node.Value{
		return node.Left.findNode(num)
	}else{
		return node.Right.findNode(num)
	}
}

func (node *Node) delete(num int){
	node = node.findNode(num)
	if node.Left == nil && node.Right == nil{
		if node.Parent.Left == node{
			node.Parent.Left = nil
		}else{
			node.Parent.Right = nil
		}
	} else if node.Left == nil && node.Right != nil{
		node.Right.Parent = node.Parent
		node.Parent.Right = node.Right
	} else if node.Left != nil && node.Right == nil{
		node.Left.Parent = node.Parent
		node.Parent.Left = node.Left
	} else {
		// two children
		if node.Left.Height > node.Right.Height{
			largest := node.getLargest()
			largest.Parent.Right = nil
			node.Right.Parent = largest
			node.Left.Parent = largest
			largest.Left = node.Left
			largest.Right = node.Right
		}else{
			smallest := node.getSmallest()
			smallest.Parent.Left = nil
			node.Left.Parent = smallest
			node.Right.Parent = smallest
			smallest.Left = node.Left
			smallest.Right = node.Right
		}
	}

}

func (node *Node) getLargest() *Node {
	curr := node.Right
	for curr.Right != nil{
		curr = curr.Right
	}
	return curr
}

func (node *Node) getSmallest() *Node{
	curr := node.Left
	for curr.Left != nil{
		curr = curr.Left
	}
	return curr
}

func main() {
	tree := Node{17, 1, nil, nil, nil}
	tree.insert(18)
	tree.insert(6)
	tree.insert(3)
	tree.insert(9)
	tree.insert(50)
	tree.insert(37)
	printTree(&tree, "", false)
	fmt.Printf("50 exists: %v\n", tree.find(50))
	fmt.Printf("57 exists: %v\n", tree.find(57))
	fmt.Printf("deleting 3...\n")
	tree.delete(3)
	printTree(&tree, "", false)
	fmt.Printf("deleting 37...\n")
	tree.delete(37)
	printTree(&tree, "", false)
	fmt.Printf("deleting 6...\n")
	tree.delete(6)
	printTree(&tree, "", false)
}
//TODO: fix edge cases on delete, nil pointer dereference on root delete
















