package main

import "fmt"

func main() {
	fmt.Println("Quick Sort")

	var arr = []int{23, 21, 14, 57, 88, 3, 99, 54, 23, 9}
	fmt.Println("Unsorted Array: ", arr)
	fmt.Println("Quick Sort: ", quickSort(arr, 0, len(arr)-1))
}

func quickSort(arr []int, left int, right int) []int {
	if left < right {

		pviotIndex := pviot(arr, left, right)

		// left
		quickSort(arr, left, pviotIndex-1)

		// right
		quickSort(arr, pviotIndex+1, right)
	}

	return arr
}

// return the index of first elemenet after it was half sorted
func pviot(arr []int, start int, end int) int {
	pviot := arr[start]
	swapIndex := start

	for i := start + 1; i < len(arr); i++ {
		if pviot > arr[i] {
			swapIndex++
			arr[swapIndex], arr[i] = arr[i], arr[swapIndex]
		}
	}
	arr[start], arr[swapIndex] = arr[swapIndex], arr[start]

	return swapIndex
}
