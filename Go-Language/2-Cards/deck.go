package main

import "fmt"

//Create new  type of deck
//whic is a slic of decks

type deck []string

func (d deck) printDeck(){
	for i, card := range d {
		fmt.Println(i,card)
	}
}