package main

import "fmt"

func main() {
	cards := newDeck()
	// hand, remainingCards := deal(cards, 5)

	// hand.printDeck()
	// remainingCards.printDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")
}

// func newCard() string {
// 	return "Five of Diamonds"
// }