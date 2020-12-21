package main

import "fmt"

func main() {
	cards := newDeck()
	cards = append(cards, "Six of Spades")
	// hand, remaining := deal(cards, 5)

	cards.saveToFile("out.txt")
	fmt.Println(cards.toString())
}
