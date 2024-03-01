package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for k := '9'; k >= '0'; k-- {
				for m := '9'; m >= '0'; m-- {
					if (i*10 + j) > (k*10 + m) {
						if i == '0' && j == '1' && k == '0' && m == '0' {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(' ')
							z01.PrintRune(k)
							z01.PrintRune(m)

						} else {
							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(' ')
							z01.PrintRune(k)
							z01.PrintRune(m)
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
