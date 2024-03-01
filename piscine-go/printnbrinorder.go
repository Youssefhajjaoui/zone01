package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var tab []int
	if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		a := n % 10
		tab = append(tab, a)
		n = n / 10
	}

	for i := 0; i < len(tab); i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[i] > tab[j] {
				tab[i], tab[j] = tab[j], tab[i]
			}
		}
	}

	for i := 0; i < len(tab); i++ {
		z01.PrintRune(rune(tab[i] + '0'))
	}
}
