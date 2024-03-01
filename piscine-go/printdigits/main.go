package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for z := 0; z <= 9; z++ {
		a := rune(z) + '0'
		z01.PrintRune(a)
	}
	z01.PrintRune('\n')
}
