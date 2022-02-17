package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Radix Sort")

	var arr = []int{23, 211, 1564, 57, 6, 88, 35, 99, 1, 554, 213, 98}
	fmt.Println("Unsorted Array: ", arr)

	// fmt.Println("Get Digit: ", getDigit(153, 1))
	// fmt.Println("Digit Count: ", digitCount(15233))

	fmt.Println("Radix Sort: ", radixSort(arr))

}

func radixSort(arr []int) []int {

	maxDigitsCount := mostDigits(arr)

	for i := 1; i < maxDigitsCount; i++ {

		var digitBuckets [10][]int //array of 10 empty slices
		for j := 0; j < len(arr); j++ {
			digit := getDigit(arr[j], i)
			digitBuckets[digit] = append(digitBuckets[digit], arr[j])

		}

		arr = append(arr, digitBuckets[i]...)
	}

	return arr
}

// get the number at a specific digit
func getDigit(num int, digit int) int {
	r := num % int(math.Pow(10, float64(digit)))
	return r / int(math.Pow(10, float64(digit-1)))
}

// how many digits are in a number
func digitCount(num int) int {
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}

	return count
}

// finding which number has the most digits in an array
func mostDigits(arr []int) int {
	maxDigits := 0
	for i := 0; i < len(arr); i++ {
		maxDigits = int(math.Max(float64(maxDigits), float64(digitCount(arr[i]))))

	}
	return maxDigits
}
