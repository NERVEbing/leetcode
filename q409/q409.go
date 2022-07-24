package q409

import (
	"fmt"
)

func longestPalindrome(s string) int {
	m := make(map[uint8]int)
	c := 0

	for i := 0; i < len(s); i++ {
		_, ok := m[s[i]]
		if ok {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
		if m[s[i]]%2 == 0 {
			c += 2
			m[s[i]] = 0
		}
	}

	if c < len(s) {
		c++
	}

	return c
}

func Test() {
	s := "banana"
	result := longestPalindrome(s)
	fmt.Println(result)
}
