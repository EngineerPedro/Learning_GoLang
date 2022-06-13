package main

import "fmt"

type ListNode_01 struct {
	Val  int
	Next *ListNode_01
}

func main() {
	List_1 := &ListNode_01{1, &ListNode_01{3, &ListNode_01{5, nil}}}
	List_2 := &ListNode_01{1, &ListNode_01{2, &ListNode_01{4, nil}}}
	fmt.Println(mergeTwoLists(List_1, List_2))
}

func mergeTwoLists(list1 *ListNode_01, list2 *ListNode_01) *ListNode_01 {
	//Both lists already sorted
	Final_List := &ListNode_01{}
	k := Final_List
	i, j := list1, list2
	//iterate through both lists
	for i != nil && j != nil {
		if i.Val < j.Val { //check if I is greater then
			k.Next = i
			i = i.Next
		} else { // if not
			k.Next = j
			j = j.Next
		}
		//k moves forward but final will still be at one
		k = k.Next
	}
	if i != nil {
		k.Next = i
	} else if j != nil {
		k.Next = j
	}
	return Final_List.Next
}
