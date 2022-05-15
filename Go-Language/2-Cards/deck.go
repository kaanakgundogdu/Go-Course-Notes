package main

import "fmt"

//Create new  type of deck
//whic is a slic of decks

type deck []string

func newDeck() deck  {
	cards := deck{}

	cardSuits := []string {"Spades","Hearts","Diamonds","Clubs"}
	cardValues := []string {"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Joker","Queen","King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value +" of "+ suit)
		}
	}
	return cards
}

func (d deck) printDeck(){
	for i, card := range d {
		fmt.Println(i,card)
	}
}

func deal(d deck, handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}
