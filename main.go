package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const SIZE int = 10

	// Init a new rand object
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Declare a new array
	var arr [SIZE]int

	// Assign the array with random values
	for i := 0; i < SIZE; i++ {
		arr[i] = r1.Intn(10) + 1
	}

	fmt.Println("Array before sort:", arr)

	// Insertion sort runs in-place
	for i := 1; i < SIZE; i++ {
		n := arr[i]
		j := i
		for j > 0 && arr[j-1] > n {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = n
	}

	fmt.Println("Array after sort:", arr)
}
