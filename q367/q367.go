package q367

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	left := 0
	right := num
	for left <= right {
		mid := left + (right-left)/2
		s := mid * mid
		switch {
		case s > num:
			right = mid - 1
		case s < num:
			left = mid + 1
		default:
			return true
		}
	}

	return false
}

//func isPerfectSquare(num int) bool {
//	n := 1
//	for num > 0 {
//		num -= n
//		n += 2
//	}
//
//	return num == 0
//}

func Test() {
	num := 256
	result := isPerfectSquare(num)
	fmt.Println(result)
}
