package main

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	hand.printDeck()
	remainingCards.printDeck()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }