package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	left  *node
	right *node
}

type bst struct {
	root *node //points to first element in array
	len  int   //keeps track of the size of the tree
}

func main() {
	n := &node{1, nil, nil}
	n.left = &node{0, nil, nil}
	n.right = &node{2, nil, nil}
	b := bst{
		root: n,
		len:  3,
	}
	fmt.Println(b)

	b.add(5)
	b.add(4)
	b.add(6)
}

//two string method for the node object
func (n node) String() string {
	return strconv.Itoa(n.value)
}

func (b bst) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()
}

func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root)
}

func (b bst) inOrderTraversalByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s", root))
	b.inOrderTraversalByNode(sb, root.right)
}

func (b *bst) add(value int) {
	b.root = b.addByNode(b.root, value)
	b.len++
}

func (b *bst) addByNode(root *node, value int) *node {
	if root == nil {
		return &node{value: value}
	}
	if value < root.value {
		root.left = b.addByNode(root.left, value)
	} else {
		root.right = b.addByNode(root.right, value)
	}
	return root
}
