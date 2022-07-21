package q852

import (
	"fmt"
)

func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1
	mid := left + (right-left)/2

	for !(arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1]) {
		if arr[mid] > arr[mid-1] {
			left = mid
		} else {
			right = mid
		}
		mid = left + (right-left)/2
	}

	return mid
}

func Test() {
	arr := []int{0, 10, 11, 4, 3, 5, 2}
	result := peakIndexInMountainArray(arr)
	fmt.Println(result)
}
