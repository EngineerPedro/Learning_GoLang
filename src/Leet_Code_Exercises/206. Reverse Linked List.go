package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Println(reverseList(head))
}

func reverseList(head *ListNode) *ListNode {
	//non or 1 node
	if head == nil || head.Next == nil {
		return head
	}

	//over 2 nodes
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
