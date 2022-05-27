package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generateSlice_02(20)
	a := []int{-35, -60, 486, 55}
	fmt.Println("\n--- Unsorted --- \n\n", slice, "\n")
	fmt.Println("\n--- Unsorted --- \n\n", a)
	insertionsort(a)
	insertionsort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
	fmt.Println("\n--- Unsorted --- \n\n", a)
}

// Generates a slice of size, size filled with random numbers
func generateSlice_02(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func insertionsort(items []int) {
	for currentvalue := 1; currentvalue < len(items); currentvalue++ {
		for comparriing := currentvalue; comparriing > 0; comparriing-- {
			if items[comparriing-1] > items[comparriing] {
				items[comparriing-1], items[comparriing] = items[comparriing], items[comparriing-1]
			}
		}
	}
}
