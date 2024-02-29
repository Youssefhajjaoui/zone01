package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	rs := []rune(os.Args[0])
	for i := range rs {
		if i > 1 {
			z01.PrintRune(rs[i])
		}
	}
	z01.PrintRune('\n')
}
