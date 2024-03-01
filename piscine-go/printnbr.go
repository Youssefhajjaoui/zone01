package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	t := 1
	if n < 0 {
		z01.PrintRune('-')
		t = -1

	}
	if n != 0 {
		a := (n / 10) * t
		if a != 0 {
			PrintNbr(a)
		}
		b := (n % 10 * t) + '0'
		z01.PrintRune(rune(b))

	} else {
		z01.PrintRune('0')
	}
}
