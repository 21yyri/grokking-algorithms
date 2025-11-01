package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func BinarySearch(slice []int, target int) (int, error) {
	start := time.Now()
	defer func() {
		fmt.Printf("Duration of binary search: %v\n", time.Since(start))
	}()

	low, high := 0, len(slice)-1

	for {
		if low > high {
			break
		}

		half := (low + high) / 2

		if slice[half] == target {
			return half, nil
		}

		if slice[half] > target {
			high = half - 1
		} else {
			low = half + 1
		}
	}

	return -1, errors.New("index not found")
}

func LinearSearch(array []int, target int) (int, error) {
	start := time.Now()
	defer func() {
		fmt.Printf("Duration of linear search: %v\n", time.Since(start))
	}()

	for id, item := range array {
		if item == target {
			return id, nil
		}
	}

	return -1, errors.New("index not found")
}

// MAIN
func main() {
	array := []int{}

	for i := 0; i <= 1e8; i++ {
		array = append(array, i)
	}

	var target int
	fmt.Scan(&target)

	_, err := LinearSearch(array, target)
	if err != nil {
		log.Fatal(err)
	}

	_, errorVar := BinarySearch(array, target)
	if errorVar != nil {
		log.Fatal(err)
	}

}
