package q205

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	l := len(s)
	m1 := make(map[uint8]int)
	m2 := make(map[uint8]int)

	for i := 0; i < l; i++ {
		if _, ok := m1[s[i]]; !ok {
			m1[s[i]] = i
		}
		if _, ok := m2[t[i]]; !ok {
			m2[t[i]] = i
		}
		if m1[s[i]] != m2[t[i]] {
			return false
		}
	}

	return true
}

func Test() {
	s1 := "egg"
	s2 := "add"
	result := isIsomorphic(s1, s2)
	fmt.Println(result)
}
