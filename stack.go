package main

import (
	"errors"
)

// Stack implements an integer based stack. It's currently implemented
// with a Go slice. However, a single linked list can work for this too
// as an underlying implementation.
type Stack struct {
	data         []int
	currentIndex int
}

// Push appends an element from the stack
//
// `O(1)` complexity
func (stack *Stack) Push(value int) error {
	if stack.currentIndex >= len(stack.data) {
		return errors.New("stack overflow")
	}

	stack.data[stack.currentIndex] = value
	stack.currentIndex++

	return nil
}

// Pop removes an element from the stack
//
// `O(1)` complexity
func (stack *Stack) Pop() (int, error) {
	if stack.currentIndex <= 0 {
		return 0, errors.New("no elements in stack")
	}

	value := stack.data[stack.currentIndex]

	// Just move the index. Leave the used position just right there.
	stack.currentIndex--

	return value, nil
}
