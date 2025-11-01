package main

import (
	"fmt"
	"time"
)

func RemoveItem(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func SearchSmallestIndex(array []int) int {
	smallestNum := array[0]
	smallestIndex := 0

	for i := 0; i < len(array); i++ {
		if array[i] < smallestNum {
			smallestNum = array[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}

func SelectSort(array []int) []int {
	start := time.Now()
	defer func() {
		fmt.Printf("The time spent in the function was %v\n", time.Since(start))
	}()

	newSlice := []int{}
	for i := 0; i < len(array); i++ {
		newSlice = append(newSlice, array[SearchSmallestIndex(array)])

		array = RemoveItem(array, SearchSmallestIndex(array))
	}
	return newSlice
}

func main() {
	newArray := []int{}

	for i := 1_000; i > -1; i-- {
		newArray = append(newArray, i)
	}
	fmt.Println("oiee")

	SelectSort(newArray)
}
