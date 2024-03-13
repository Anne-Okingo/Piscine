package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	i := 0
	for j := 0; j < 4; j++ {
		fmt.Printf("Player %v: %v, %v, %v\n", j+1, deck[i], deck[i+1], deck[i+2])
		i = i + 3
	}

}
