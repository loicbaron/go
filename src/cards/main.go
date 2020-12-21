package main

import "fmt"

func main() {
	cards := newDeck()
	cards = append(cards, "Six of Spades")
	// hand, remaining := deal(cards, 5)

	fmt.Println(cards.toString())
}
