package main

import (
	"fmt"
)

func main() {
	fmt.Println("Binary Search ")

	arr := []int{1, 2, 4, 6, 9, 12, 23, 45, 66, 67, 72, 75, 77, 81, 83, 89, 90, 99}
	num := 66

	i := bs(arr, num)
	fmt.Println(i)

}

func bs(arr []int, num int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2
		fmt.Println(start, mid, end)
		if arr[mid] == num {
			return mid
		} else if arr[mid] < num {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}
	return -1 //if the number is not found
}
