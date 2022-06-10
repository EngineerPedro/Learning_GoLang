package main

import (
	"container/list"
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	st := list.New()
	closeB := map[rune]rune{'}': '{', ']': '[', ')': '('}
	for _, b := range s {
		// if a closing bracket, check top of stack has corresponding opening bracket
		if openB, ok := closeB[b]; ok {
			// stack is empty or its not the opening bracket needed
			if st.Len() <= 0 || st.Back().Value.(rune) != openB {
				return false
			}
			st.Remove(st.Back())
		} else {
			st.PushBack(b)
		}
	}
	if st.Len() > 0 {
		return false
	}
	return true
}
