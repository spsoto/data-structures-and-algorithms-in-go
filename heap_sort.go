package main

// HeapSort sorts an array int using the heap sort algorithm. This usually
// employs `O(n)` computational time.
func HeapSort(arr []int) {
	h := Heap{data: arr, heapLength: len(arr)}

	// Build Max Heap
	h.MaxBuild()

	// Now start iterating
	for i := len(arr) - 1; i > 0; i-- {
		// Exchange. Note that one of the lowest values of the heap
		// is now on top.
		h.Swap(0, i)

		// Decrease heap size
		h.heapLength--

		// Restore heapify priority from the origin
		h.MaxHeapify(0)
	}
}
