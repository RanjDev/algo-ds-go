package main

import "fmt"

func main() {
	fmt.Println("Merge Sort")

	var arr = []int{23, 21, 14, 57, 88, 3, 99, 54, 23, 9}
	fmt.Println("Unsorted Array: ", arr)
	fmt.Println("Merge Sort", mergeSort(arr))

}

func mergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	// split the array into 2 pieces
	mid := len(arr) / 2
	leftSide := arr[:mid]
	rightSide := arr[mid:]

	left := mergeSort(leftSide)
	right := mergeSort(rightSide)

	// fmt.Println("left", left, " --- right", right)

	return merge(left, right)
}

func merge(arr1 []int, arr2 []int) []int {
	var arr []int
	// fmt.Println("arr1", arr1, " --- arr2", arr2)

	index1 := 0 //index for the first array
	index2 := 0 //index for the second array

	for index1 < len(arr1) && index2 < len(arr2) {
		if arr2[index2] > arr1[index1] {
			arr = append(arr, arr1[index1])
			index1++
		} else {
			arr = append(arr, arr2[index2])
			index2++
		}
	}

	/*
		when 1 array is larger than the other it means that all
		the rest of that array that is not compared to the other
		array is sorted, so we will just add it to the merged array
	*/
	for index1 < len(arr1) {
		arr = append(arr, arr1[index1])
		index1++
	}
	for index2 < len(arr2) {
		arr = append(arr, arr2[index2])
		index2++
	}

	// fmt.Println("arr", arr)
	return arr //the merged array
}
