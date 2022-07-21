package q1385

import "fmt"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	result := 0
	for i := 0; i < len(arr1); i++ {
	NEXT:
		for n := 0; n < len(arr2); n++ {
			x := arr1[i] - arr2[n]
			if x < 0 {
				x = -x
			}
			if x <= d {
				break NEXT
			}
			if n == len(arr2)-1 {
				result++
			}
		}
	}

	return result
}

func Test() {
	arr1 := []int{2, 1, 100, 3}
	arr2 := []int{-5, -2, 10, -3, 7}
	d := 6
	result := findTheDistanceValue(arr1, arr2, d)
	fmt.Println(result)
}
