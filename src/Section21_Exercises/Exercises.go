package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	Exercise_01()
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
