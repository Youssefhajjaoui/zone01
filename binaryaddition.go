package main

import "fmt"

func main() {
	fmt.Println(BinaryAddition(1, 1))
	fmt.Println(BinaryAddition(1, 2))
	fmt.Println(BinaryAddition(1, 3))
	fmt.Println(BinaryAddition(2, 1))
	fmt.Println(BinaryAddition(2, 2))
	fmt.Println(BinaryAddition(1, 16))
}
func BinaryAddition(a int, b int) []int {
	x := a + b
	tab := []int{}
	for x > 0 {
		tab = append(tab, x%2)
		x = x / 2
	}
	for i := 0; i < len(tab)/2; i++ {
		tab[i], tab[len(tab)-1-i] = tab[len(tab)-1-i], tab[i]
	}
	return tab
}
