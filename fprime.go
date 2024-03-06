package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	b := os.Args[1]
	target, _ := strconv.Atoi(b)
	str := ""
	for i := 2; i <= target; i++ {
		for target%i == 0 {
			if i != target {
				str += strconv.Itoa(i) + "*"
			} else {
				str += strconv.Itoa(i)
			}
			target = target / i
		}
	}
	fmt.Println(str)
}
