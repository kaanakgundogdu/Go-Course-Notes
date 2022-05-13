package main

import "fmt"

func main() {
	// card := newCard()
	cards := []string{
		newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")
	fmt.Println((cards))

	for i, card := range cards {
		fmt.Println(i,card)
	} 
}

func newCard() string {
	return "Five of Diamonds"
}