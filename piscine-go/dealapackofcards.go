package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	count := 0
	for i := 0; i < 12; i += 3 {
		count++
		fmt.Printf("Player %v: %v, %v, %v\n", count, deck[i+0], deck[i+1], deck[i+2])
	}
}
