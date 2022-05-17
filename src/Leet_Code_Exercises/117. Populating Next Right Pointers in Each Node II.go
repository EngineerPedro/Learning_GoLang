package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {

	//BreathFirstSearch(test)

}

func BreathFirstSearch(root *Node) *Node {
	if root == nil {
		return root
	}
	que := []*Node{root}

	for len(que) != 0 {
		// Reset next node at the start of each level
		var next *Node

		for _ = range que {
			current := que[0]
			que = que[1:]

			// The previous node in the reverse level order traversal is `current.next`
			current.Next = next

			// Add right first and left second for reverse level order
			if current.Right != nil {
				que = append(que, current.Right)
			}
			if current.Left != nil {
				que = append(que, current.Left)
			}
			next = current
		}
	}

	return root
}
