package main

// Merge merges two arrays from an specific point
func merge(arr []int, j int) {
	if len(arr) <= 1 {
		return
	}

	leftArr := make([]int, len(arr[:j]))
	rightArr := make([]int, len(arr[j:]))
	copy(leftArr, arr[:j])
	copy(rightArr, arr[j:])

	i, j := 0, 0

	for k := 0; k < len(arr); k++ {
		if j >= len(rightArr) {
			// Out of bounds of the right array. Proceed copying i
			arr[k] = leftArr[i]
			i++
		} else if i >= len(leftArr) || leftArr[i] > rightArr[j] {
			// Out of bounds of the left array. Proceed copying j
			arr[k] = rightArr[j]
			j++
		} else {
			// Otherwise just copy i
			arr[k] = leftArr[i]
			i++
		}

		// Note that len(leftArr) + len(rightArr) = len(arr) so with this
		// approach there should be no more out of bounds exceptions.
	}
}

// MergeSort performs merge sort strategy over an int slice.
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2

	MergeSort(arr[:mid])
	MergeSort(arr[mid:])
	merge(arr, mid)
}
