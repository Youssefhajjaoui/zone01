package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	st := os.Args[1:]
	sr := ""
	for i := 0; i < len(st); i++ {
		sr += st[i]
	}

	for i := 0; i < len(st); i++ {
		for j := i + 1; j < len(st); j++ {
			if st[i] > st[j] {
				st[i], st[j] = st[j], st[i]
			}
		}
	}
	for i := 0; i < len(st); i++ {
		v := st[i]
		rs := []rune(v)

		for i := 0; i < len(rs); i++ {
			z01.PrintRune(rs[i])
		}
		z01.PrintRune('\n')
	}
}
