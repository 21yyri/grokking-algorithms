package main

import (
	"fmt"
	"time"
)

func Quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	lesser := []int{}
	greater := []int{}

	pivot := arr[0]

	for _, num := range arr[1:] {
		if pivot >= num {
			lesser = append(lesser, num)
		} else if pivot < num {
			greater = append(greater, num)
		}
	}

	return append(append(Quicksort(lesser), pivot), Quicksort(greater)...)
}

func main() {
	array := []int{}
	for x := 1000; x > 0; x-- {
		array = append(array, x)
	}

	startTime := time.Now()
	sortedArray := Quicksort(array)
	fmt.Printf("The sorted array is: %v\nIt took the program %v to sort it.\n", sortedArray, time.Since(startTime))
}
