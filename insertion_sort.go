package main

// InsertionSort performs an insertion sort over the slice arr
func InsertionSort(arr []int) {
	for i, element := range arr[1:] {
		j := i
		for ; j >= 0 && arr[j] > element; j-- {
			arr[j+1] = arr[j]
		}

		arr[j+1] = element
	}
}
