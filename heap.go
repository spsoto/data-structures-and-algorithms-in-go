package main

import (
	"errors"
	"fmt"
)

// Heap represents a binary heap of integers.
type Heap struct {
	data       []int
	heapLength int
}

// At returns the value of the heap at a particular position.
//
// `O(1)` complexity.
func (h *Heap) At(i int) (int, error) {
	if i >= h.heapLength {
		return 0, errors.New("out of bound")
	}

	return h.data[i], nil
}

// Set sets the value at position i with value key
func (h *Heap) Set(i int, key int) error {
	if i >= h.heapLength {
		return errors.New("out of bound")
	}

	h.data[i] = key

	return nil
}

// Swap exchanges two positions withing the heap. This method
// **does not** ensures that the max / min heap is preserved.
//
// `O(1)` complexity.
func (h *Heap) Swap(i int, j int) error {
	if i >= h.heapLength || j >= h.heapLength {
		return errors.New("out of bound")
	}

	elemCopy, _ := h.At(i)
	second, _ := h.At(j)
	h.Set(i, second)
	h.Set(j, elemCopy)

	return nil
}

// Parent returns the parent element of a particular key of the heap.
//
// `O(1)` complexity.
func (h *Heap) Parent(i int) int {
	return (i+1)/2 - 1
}

// Left returns the value of index of the left
// child of a node in the binary heap.
//
// **Does not** ensure that the positions is already
// in use.
//
// `O(1)` complexity.
func (h *Heap) Left(i int) int {
	return 2*(i+1) - 1
}

// Right returns the value of the index of the right
// child of a node in the binary heap.
//
// **Does not** ensure that the positions is already
// in use.
//
// `O(1)` complexity.
func (h *Heap) Right(i int) int {
	return 2*(i+1) + 1 - 1
}

// MaxHeapify takes a position of a heap at position i and moves it
// down the heap to ensure the max heap property of the heap.
//
// `O(log n)` complexity.
func (h *Heap) MaxHeapify(i int) {
	// Get the current element. If it doesn't exist, there's actually
	// nothing to do here.
	currentElem, err := h.At(i)

	if err != nil {
		return
	}

	largest, leftIndex, rightIndex := 0, h.Left(i), h.Right(i)

	// Check if the left element is larger
	leftElem, err := h.At(leftIndex)

	if err == nil && leftElem > currentElem {
		largest = leftIndex
	} else {
		largest = i
	}

	// Check if the right element is greater than the
	// largest we have.
	rightElem, err := h.At(rightIndex)
	currentElem, _ = h.At(largest)

	if err == nil && rightElem > currentElem {
		largest = rightIndex
	}

	// If the current key is the largest item, then no further
	// work (swapping) is required since it's assumed the tree
	// down below already satisfies the max heap property.
	if largest == i {
		return
	}

	// If the largest is the index, then we have to move down
	h.Swap(i, largest)

	h.MaxHeapify(largest)
}

// MaxBuild takes an arbitrary slice on the heap and makes it a MaxHeap
// that indeed has the max heap property.
//
// This method has `O(n * log n)` complexity.
func (h *Heap) MaxBuild() {
	// The length of the heap will be what we currently have on the slice
	h.heapLength = len(h.data) // O(1)

	// At exactly the half, all the other remaining half are children,
	// thus the first iteration is trivial and afterwards the children
	// nodes on top are already compliant with the max heapify property.
	for i := h.heapLength/2 - 1; i >= 0; i-- { // n times
		h.MaxHeapify(i) // O(log n)
	}
}

// Maximum returns the biggest element of a max heap, which is the first
// element of the heap after is built.
//
// This method has `O(1)` complexity.
func (h *Heap) Maximum() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("unsufficient elements in array")
	}

	return h.At(0)
}

// ExtractMaximum extracts the maximum of the heap and then restarts
// the heap from the bottom to be rebuilt.
//
// `O(log n)` complexity.
func (h *Heap) ExtractMaximum() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("unsufficient elements in array")
	}

	// Take the element as a copy
	max, _ := h.At(0)

	// We swap with the latest element since this way the heap structure
	// is preserved and just the latest element must go all the way down again,
	// enabling a quick re-sort of the heap.
	h.Swap(0, len(h.data)-1)
	h.heapLength--

	// Rebuild the heapify process
	h.MaxHeapify(0)

	return max, nil
}

// IncreaseKey increases the value of a key at a particular position and then
// ensures that the max heap property is met doing a heapify process starting
// at that node.
func (h *Heap) IncreaseKey(i int, key int) error {
	currElem, err := h.At(i)
	if err != nil {
		return err
	}

	if currElem >= key {
		return fmt.Errorf("key %d is not bigger than current value %d at key %d", key, currElem, i)
	}

	// Set the value of the key
	err = h.Set(i, key)
	if err != nil {
		return err
	}

	// Update the heap
	for { // log n times at most
		if i == 0 {
			break
		}

		parent := h.Parent(i)

		h.Swap(parent, i)
		i = parent
	}

	return nil
}

// MaxInsert inserts a key and properly updates the heap with the max heap property.
//
// `O(log n)` complexity.
func (h *Heap) MaxInsert(key int) {
	// Update the length of the heap
	length := h.heapLength
	h.heapLength++

	// Start with key - 1 to have something to actually increase
	h.data = append(h.data, key-1)

	h.IncreaseKey(length, key)
}
