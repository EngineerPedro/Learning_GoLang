package main

import "fmt"

//Node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert will add a node to the tree
//the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

//Search the tree for a specific value
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k) //recursion
	} else if n.Key > k {
		return n.Left.Search(k) //recursion
	}

	return true
}

func invertTree(root *Node) *Node {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(500)
	tree.Insert(50)
	fmt.Printf("%+v\n", tree)
	invertTree(tree)
}
