package main

import "fmt"

func main() {
	cards := newDeck()
	cards = append(cards, "Six of Spades")
	hand, remaining := deal(cards, 5)
	fmt.Println("hand = ")
	hand.print()
	fmt.Println("remaining = ")
	remaining.print()
	fmt.Println("cards = ")
	cards.print()
}
