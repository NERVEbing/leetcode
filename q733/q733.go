package q733

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	var f func(int, int)
	flag := image[sr][sc]

	f = func(x int, y int) {
		if image[x][y] == color {
			return
		}
		image[x][y] = color

		if x-1 >= 0 {
			if image[x-1][y] == flag {
				f(x-1, y)
			}
		}

		if x+1 <= len(image)-1 {
			if image[x+1][y] == flag {
				f(x+1, y)
			}
		}

		if y-1 >= 0 {
			if image[x][y-1] == flag {
				f(x, y-1)
			}
		}

		if x <= len(image)-1 && y+1 <= len(image[x])-1 {
			if image[x][y+1] == flag {
				f(x, y+1)
			}
		}
	}

	f(sr, sc)

	return image
}

func Test() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	color := 2

	result := floodFill(image, sr, sc, color)
	fmt.Println(result)
}
