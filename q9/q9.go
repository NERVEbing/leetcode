package q9

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	num := 0
	for x > num {
		num = num*10 + x%10
		x /= 10
	}

	return x == num || x == num/10
}

func Test() {
	result := isPalindrome(34543)
	fmt.Println(result)
}
