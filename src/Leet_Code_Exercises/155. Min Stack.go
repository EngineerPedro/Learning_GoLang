package main

import "fmt"

func main() {

	obj := Constructor_01()
	obj.Push(4)
	obj.Push(40)
	obj.Push(87)
	obj.Push(8)
	obj.Push(2)
	fmt.Println(obj.minStack)
	fmt.Println(obj.stack)
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
	obj.Pop()
	obj.Pop()
	fmt.Println(obj.minStack)
	fmt.Println(obj.stack)
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
	obj.Pop()
	obj.Pop()
	fmt.Println(obj.minStack)
	fmt.Println(obj.stack)
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())

}

type MinStack struct {
	minStack []int // This stores the minumum elements
	stack    []int // This stores all the elements
}

// Initialize your data structure here
func Constructor_01() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	//insert new value into stack
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 { // check if the stack is empty
		this.minStack = append(this.minStack, x) // if is is also append the new value to the minStack
		return
	}
	// IF NOT compare the values
	// Currrent element less than previously known minimum, push it in the min stack.
	if x <= this.minStack[len(this.minStack)-1] { // comparing elements
		this.minStack = append(this.minStack, x) // if true add new minimum
	}
}

func (this *MinStack) Pop() {
	x := this.stack[len(this.stack)-1]          // Store desired element wanting to be removed in x
	this.stack = this.stack[:len(this.stack)-1] // Gets reid of that element

	if x == this.minStack[len(this.minStack)-1] { // Check if element is on top of min stack
		this.minStack = this.minStack[:len(this.minStack)-1] // If so get rid of it
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 { // Check if the stack is empty
		return 0 // If so return 0
	}
	return this.stack[len(this.stack)-1] // return whats on top of the stack
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 { // Check if the minStack is empty
		return 0 //Return Zero if True
	}
	return this.minStack[len(this.minStack)-1] // return whats on the top of the min stack
}
