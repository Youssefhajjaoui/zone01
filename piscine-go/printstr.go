package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	phrase := []rune(s)
	for i := 0; i <= len(phrase)-1; i++ {
		z01.PrintRune(phrase[i])
	}
}
