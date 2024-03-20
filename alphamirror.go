package main

import (
	"fmt"
	"os"
)

func main() {
	arr := os.Args[1:]
	for _, v := range arr {
		fmt.Printf(convert(v))
		fmt.Printf("\n")
	}
}
func convert(s string) string {
	result := ""
	for _, t := range s {
		if t >= 'A' && t <= 'Z' {
			t = rune(90 - (t - 65))
			result += string(t)
		} else if t >= 'a' && t <= 'z' {
			t = rune(122 - (t - 97))
			result += string(t)
		} else {
			result += string(t)
		}
	}
	return result
}
