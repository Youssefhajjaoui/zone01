package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printNum(x int) {
	a := '0'
	b := '0'
	if x < 0 {
		z01.PrintRune('-')
	}
	for i := 0; i < x/10; i++ {
		a++
	}
	z01.PrintRune(a)
	for j := 0; j < x%10; j++ {
		b++
	}
	z01.PrintRune(b)
}

func printstring(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

func main() {
	points := &point{}

	setPoint(points)
	printstring("x = ")
	printNum(points.x)
	printstring(", y = ")
	printNum(points.y)
	printstring("\n")
}
