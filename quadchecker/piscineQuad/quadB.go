package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		if i == 1 {
			z01.PrintRune('/')
		} else if i == y {
			z01.PrintRune('\\')
		} else {
			z01.PrintRune('*')
		}
		if x != 1 {
			for j := 1; j <= x-2; j++ {
				if i == 1 || i == y {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}

			if i == 1 {
				z01.PrintRune('\\')
			} else if i == y {
				z01.PrintRune('/')
			} else {
				z01.PrintRune('*')
			}
		}
		z01.PrintRune('\n')
	}
}
func main() {
	args := os.Args
	if len(args) != 3 {
		return
	} else {
		b, _ := strconv.Atoi(args[1])
		j, _ := strconv.Atoi(args[2])
		QuadB(b, j)
	}

}
