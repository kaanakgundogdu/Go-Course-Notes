package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join(d,",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName,[]byte(d.toString()),0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		//log error & and return call to newdec
		//log error & quit program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	
	return deck(s)
}

