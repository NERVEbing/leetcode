package q6

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	groupLen := 2*numRows - 2
	numRowsIndex := numRows - 1

	strArr := make([]string, numRows)

	for i := 0; i < len(s); i++ {
		winIndex := i % groupLen
		if winIndex > numRowsIndex {
			winIndex = numRowsIndex - (winIndex - numRowsIndex)
		}
		strArr[winIndex] += string(s[i])
	}

	return strings.Join(strArr, "")
}

func Test() {
	str := "LEETCODEISHIRING"
	numRows := 4
	result := convert(str, numRows)
	fmt.Println(result)
}
