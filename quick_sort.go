package main

func partition(arr []int, p int, r int) int {
	pivot, i := arr[r], p-1

	// Run over the entire array, with the exception of
	// the pivot
	for j, element := range arr[p:r] {
		if element <= pivot {
			i++ // keep current element on the left side

			// exchange element i with the position occupied
			// note the offset by p units to compensate that
			// j starts at 0 on the range.
			arr[i], arr[p+j] = arr[p+j], arr[i]
		}
	}

	arr[i+1], arr[r] = arr[r], arr[i+1]

	// Return the final position of the pivot
	return i + 1
}

// QuickSort performs the quick sort algorithm over an array
// of integers.
//
// Has `O(n^2)` complexity but usually does the sorting on
// $\Theta(n * log n)$ and much faster than heap sort because
// of the constant costs.
func QuickSort(arr []int, p int, r int) {
	// Base case: element of length 1
	if p < r {
		q := partition(arr, p, r)
		QuickSort(arr, p, q-1)
		QuickSort(arr, q+1, r)
	}
}
