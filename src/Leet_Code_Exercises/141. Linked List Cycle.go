package main

func main() {

}

type ListNode_03 struct {
	Val  int
	Next *ListNode_03
}

func hasCycle(head *ListNode_03) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
