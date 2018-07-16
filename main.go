package main

import (
	"fmt"
)

func main() {
	a := [9]int{5, 3, 17, 10, 84, 19, 6, 22, 9}

	// Copy the data to a new array to work with
	tmp := make([]int, len(a))
	copy(tmp, a[:])

	// Try the sorting algorithm
	QuickSort(tmp, 0, len(tmp)-1)

	fmt.Println(a, tmp)
}
