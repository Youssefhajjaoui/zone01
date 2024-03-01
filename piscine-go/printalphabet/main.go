package main

import "github.com/01-edu/z01"

func main() {
	for z := 'a'; z <= 'z'; z++ {
		z01.PrintRune(z)
	}
	z01.PrintRune('\n')
}
