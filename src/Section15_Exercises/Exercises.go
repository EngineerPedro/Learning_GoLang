package main

func main() {
	Exercise_01()
	Exercise_02()
	Exercise_03()
	Exercise_04()
	Exercise_05()
	Exercise_06()
	Exercise_07()
	Exercise_08()
	Exercise_09()
	Exercise_10()
}
func Exercise_01() {
	/*	Hands-on exercise #1 ● Review
		○ functions
		○ purpose of functions
		■ abstract code
		■ code reusability
		● DRY - Don’t Repeat Yourself
		○ func, receiver, identifier, params, returns
		○ parameters vs arguments
		○ variadic funcs
		■ multiple “variadic” params
		■ multiple “variadic” args
		○ returns
		■ multiple returns
		■ named returns - yuck!
		○ func expressions
		■ assigning a func to a variable ○ callbacks
		■ passing a func into another func as an argument
		○ closure
		■ one scope enclosing another
		■ variables declared in the outer scope are accessible in inner scopes
		■ closure helps us limit the scope of variables
		○ recursion
		■ a func that calls itself
		■ factorial ● life philosophy
		○ focus on what’s important; not upon what’s urgent ● Hands on exercise
		○ create a func with the identifier foo that returns an int
		○ create a func with the identifier bar that returns an int and a string
		○ call both funcs
		○ print out their results
	code: https://play.golang.org/p/owEJNF5WAD*/

	println(foo())
	println(bar())

}

func Exercise_02() {
	/*	Hands-on exercise #2
		● create a func with the identifier foo that
		○ takes in a variadic parameter of type int
		○ pass in a value of type []int into your func (unfurl the []int)
		○ returns the sum of all values of type int passed in
		● create a func with the identifier bar that
		○ takes in a parameter of type []int
		○ returns the sum of all values of type int passed in
	code: https://play.golang.org/p/B0yRxtBQPD video: 103*/
}
func Exercise_03() {
	/*Hands-on exercise #3
	Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
		code: https://play.golang.org/p/XluEuUD0Nw video: 104*/
}
func Exercise_04() {
	/*Hands-on exercise #4
	● Create a user defined struct with
		○ the identifier “person”
		○ the fields:
		■ first ■ last ■ age
		● attach a method to type person with
		○ the identifier “speak”
		○ the method should have the person say their name and age
		● create a value of type person
		● call the method from the value of type person
		code: https://play.golang.org/p/NnXyWdqbbw video: 105*/

}
func Exercise_05() {
	/*	Hands-on exercise #5
		● create a type SQUARE
		● create a type CIRCLE
		● attach a method to each that calculates AREA and returns it
		○ circle area= π r 2
		○ squarearea=L*W
		● create a type SHAPE that defines an interface as anything that has the AREA method
		● create a func INFO which takes type shape and then prints the area
		Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 54
		● create a value of type square
		● create a value of type circle
		● use func info to print the area of square
		● use func info to print the area of circle
		code: https://play.golang.org/p/NGGikHNvHv video: 106*/

}
func Exercise_06() {
	/*Hands-on exercise #6
		● Build and use an anonymous func
	code: https://play.golang.org/p/DQX3xEIcRe video: 107*/
}
func Exercise_07() {
	/*Hands-on exercise #7
		● Assign a func to a variable, then call that func
	code: https://play.golang.org/p/_Qu7ZAyFDH video: 108*/
}
func Exercise_08() {
	/*	Hands-on exercise #8
		● Create a func which returns a func
		● assign the returned func to a variable
		● call the returned func
	code: https://play.golang.org/p/c2HwqVE1Rd video: 109*/
}
func Exercise_09() {
	/*	Hands-on exercise #9
		A “callback” is when we pass a func into a func as an argument. For this exercise,
		● pass a func into a func as an argument
	code: https://play.golang.org/p/0yGYPKh1y7 video: 110*/
}
func Exercise_10() {
	/*Hands-on exercise #10
		Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable:
	code: https://play.golang.org/p/a56uWtoFSL
	video: 111*/
}

//Work Area

func foo() int {
	x := 15
	return x
}

func bar() string {
	x := "Pedro"
	return x
}
