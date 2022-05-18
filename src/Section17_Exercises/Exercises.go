package main

import "fmt"

func main() {
	//Exercise_01()
	Exercise_02()
}

func Exercise_01() {
	/*
		Hands-on exercise #1
		● Create a value and assign it to a variable.
		● Print the address of that value.
			code: https://play.golang.org/p/57kW8xd0qT
		video: 116
	*/
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
}

func Exercise_02() {
	/*Hands-on exercise #2
	● create a person struct
	● create a func called “changeMe” with a *person as a parameter
	○ change the value stored at the *person address
	■ important
	Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 57
	● to dereference a struct, use (*value).field
	○ p1.first
	○ (*p1).first
	● “As an exception, if the type of x is a named pointer type and (*x).f
	is a valid selector expression denoting a field (but not a method),
	x.f is shorthand for (*x).f.”
	○ https://golang.org/ref/spec#Selectors
	● create a value of type person
	○ print out the value
	● in func main
	○ call “changeMe”
	● in func main
	○ print out the value
	code: https://play.golang.org/p/AehM662HKS
	video: 117
	*/

	p1 := person{name: "Agent_Pedro"}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

//Work Area
type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Agent_Leslie"
}
