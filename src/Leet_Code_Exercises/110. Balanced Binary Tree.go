package main

import "math"

func main() {

}

func isBalanced(root *TreeNode) bool {
	// if helper function signals an inbalance the return value would be -1
	return helper(root) != -1
}

func helper(root *TreeNode) int {
	// base condition
	if root == nil {
		return 0
	}

	left := helper(root.Left)   // get height of Left tree at a given node
	right := helper(root.Right) // get height of Right tree at a given node

	// we will use -1 to signal inbalance instead of actual height of a node
	// report inbalance to parent node if difference between Left and Right is greater than 1, or
	// if Left or Right subtree returned with an inbalance signal instead of it's height
	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}

	// if tree is balanced till now return the height + 1 to parent
	return max(left, right) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func abs(i int) int {
	x := math.Abs(float64(i))
	return int(x)
}
