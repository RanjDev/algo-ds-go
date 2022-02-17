package main

import "fmt"

func main() {
	fmt.Println("Sorting")

	var arr = []int{23, 21, 14, 57, 88, 3, 99, 54, 23, 9}
	fmt.Println("Unsorted Array: ", arr)

	// fmt.Println("Bubble Sort: ", bubbleSort(arr))
	// fmt.Println("Selection Sort: ", selectionSort(arr))
	// fmt.Println("Insertion Sort: ", insertionSort(arr))
}

func bubbleSort(arr []int) []int {

	for i := len(arr) - 1; i >= 0; i-- {
		noSwap := true
		for j := 0; j < i; j++ {

			// fmt.Println(arr, arr[j], arr[j+1])
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] //swap
				// fmt.Println(arr)
				noSwap = false
			}

		}
		if noSwap {
			return arr
		}
	}
	return arr
}

// better than Bubble because it has fewer swaps
func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {

			if arr[j] < arr[min] {
				min = j
			}

		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i] //swap
		}
	}

	return arr
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		//start from the second element and
		// compare it to the element before it.
		currentVal := arr[i]
		index := i
		for j := i - 1; j >= 0 && arr[j] > currentVal; j-- {
			// we compare i with elements before it.
			arr[j+1] = arr[j]
			index = j
		}
		arr[index] = currentVal
	}
	return arr
}
