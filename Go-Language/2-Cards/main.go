package main

func main() {
	cards := deck{"Ace of Diamond", newCard()}
	cards = append(cards, newCard())
	cards.printDeck()
}

func newCard() string {
	return "Five of Diamonds"
}