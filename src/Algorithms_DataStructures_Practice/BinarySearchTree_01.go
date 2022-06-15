package main

import "fmt"

//Represents the start of the tree
var root *Node = nil

//Node represents the components of a binary search tree
type Node struct {
	data  int
	Left  *Node
	Right *Node
}

//Make a new node for what is to be inserted into the tree
func createNewNode(newData int) *Node {
	var newNode *Node = new(Node)
	newNode.data = newData
	newNode.Left = nil
	newNode.Right = nil
	return newNode
}

//In-Order traversal of the tree
func inOrder(root *Node) {
	if root == nil {
		return
	}
	//Use recursion to go through the
	inOrder(root.Left) //Traversing the left subtree
	fmt.Printf("%d", root.data)
	inOrder(root.Right) //Traversing the right subtree
}

//Insert will add a node to the tree
//the key to add should not be already in the tree
func insert(node *Node, newData int) {
	//If no tree exists we are at the
	if root == nil {
		root = &Node{data: newData, Left: nil, Right: nil}
	}

	var compareValue = newData - node.data

	if compareValue < 0 {
		if node.Left == nil {
			node.Left = createNewNode(newData)
		} else {
			insert(node.Left, newData)
		}
	} else if compareValue > 0 {
		if node.Right == nil {
			node.Right = createNewNode(newData)
		} else {
			insert(node.Right, newData)
		}
	}
}

//Search the tree for a specific value
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.data < k {
		return n.Right.Search(k) //recursion
	} else if n.data > k {
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
	//Constructing a binary search tree
	insert(root, 42)
	insert(root, 66)
	insert(root, 67)
	insert(root, 68)
	insert(root, 25)
	insert(root, 32)
	insert(root, 12)
	insert(root, 40)
	fmt.Printf("In-order traversal binary search tree: \n")
	inOrder(root)
	invertTree(root)
	inOrder(root)

}
