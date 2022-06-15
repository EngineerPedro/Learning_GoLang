package main

func main() {

}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// this is our height / maxDepth
	ans := 0
	// queue implementation
	var q []*TreeNode
	q = append(q, root, nil) // append root, and null to our queue

	for len(q) != 0 {
		// take the first element
		v := q[0]
		// pop the first element
		q = q[1:]

		// if the first element is nil
		if v == nil {
			ans += 1
			if len(q) == 0 {
				break
			}
			q = append(q, nil)
			continue
		}

		if v.Left != nil {
			q = append(q, v.Left)
		}
		if v.Right != nil {
			q = append(q, v.Right)
		}
	}
	return ans
}
