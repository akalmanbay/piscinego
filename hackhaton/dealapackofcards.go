package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	fmt.Printf("Player 1: %v, %v, %v", deck[0], deck[1], deck[2])
	fmt.Printf("\nPlayer 2: %v, %v, %v", deck[3], deck[4], deck[5])
	fmt.Printf("\nPlayer 3: %v, %v, %v", deck[6], deck[7], deck[8])
	fmt.Printf("\nPlayer 4: %v, %v, %v\n", deck[9], deck[10], deck[11])
}
