package main

import "fmt"

func main() {
	a := []int{-44, 4, 86, 9, 32, 6}
	Bubblesort(a)
	fmt.Println(a)
}

func Bubblesort(a []int) {
	noswaps := true
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				noswaps = false
			}
		}
		if noswaps {
			break
		}
	}
}
