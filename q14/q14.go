package q14

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	str := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.HasPrefix(strs[i], str) == false {
			str = str[0 : len(str)-1]
		}
	}

	return str
}

func Test() {
	result := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(result)
}
