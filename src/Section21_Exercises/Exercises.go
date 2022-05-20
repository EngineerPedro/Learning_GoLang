package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//Exercise_01()
	//Exercise_02()
	Exercise_03()
}

func Exercise_01() {
	/*
		Hands-on exercise #1
		● in addition to the main goroutine, launch two additional goroutines
		○ each additional goroutine should print something out
		● use waitgroups to make sure each goroutine finishes before your program exists
		code: https://github.com/GoesToEleven/go-programming
		video: 148

	*/
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("hello from gold ")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from silver ")
		wg.Done()
	}()

	fmt.Println("mid begin cpu", runtime.NumCPU())
	fmt.Println("mid begin gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
}

func Exercise_02() {

	/*
		This exercise will reinforce our understanding of method sets:
		● create a type person struct
		● attach a method speak to type person using a pointer receiver
		○ *person
		● create a type human interface
		○ to implicitly implement the interface, a human must have the speak method
		● create func “saySomething”
		○ have it take in a human as a parameter
		○ have it call the speak method
		● show the following in your code
		○ you CAN pass a value of type *person into saySomething
		○ you CANNOT pass a value of type person into saySomething
		● here is a hint if you need some help
		○ https://play.golang.org/p/FAwcQbNtMG
		Receivers Values
		-----------------------------------------------
		(t T) T and *T
		(t *T) *T
		code: https://github.com/GoesToEleven/go-programming
		video: 149

	*/
	p1 := person{
		First: "Leslie",
	}

	saysomething(&p1)
}

func Exercise_03() {
	/*
		● Using goroutines, create an incrementer program
		○ have a variable to hold the incrementer value
		○ launch a bunch of goroutines
		■ each goroutine should
		Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 65
		● read the incrementer value
		○ store it in a new variable
		● yield the processor with runtime.Gosched()
		● increment the new variable
		● write the value in the new variable back to the incrementer
		variable
		● use waitgroups to wait for all of your goroutines to finish
		● the above will create a race condition.
		● Prove that it is a race condition by using the -race flag
		● if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
		code: https://github.com/GoesToEleven/go-programming
		video: 150

	*/
	var wg sync.WaitGroup
	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			fmt.Println(incrementer)
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(incrementer)
}

//Work Area

type person struct {
	First string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hi there")
}

func saysomething(h human) {
	h.speak()
}
