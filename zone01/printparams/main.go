package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	rs := []string(os.Args)
	for i := 1; i < len(rs); i++ {
		pritr(rs[i])
		z01.PrintRune('\n')
	}
}

func pritr(s string) {
	tab := []rune(s)
	for i := 0; i < len(tab); i++ {
		z01.PrintRune(tab[i])
	}
}
