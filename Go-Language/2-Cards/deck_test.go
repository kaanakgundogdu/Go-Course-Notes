package main

import (
	"os"
	"testing"
)


func TestNewDeck(t *testing.T) {
	deckForTest := newDeck()
	wantedLength:=52
	if len(deckForTest) !=wantedLength {
		t.Errorf("Expected deck length of %d but got %d",wantedLength ,len(deckForTest))
	}

	if deckForTest[0]!="Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v" , deckForTest[0])
	}

	if deckForTest[len(deckForTest) -1]!="King of Clubs" {
		t.Errorf("Expected last card of King of Clubs but got %v" , deckForTest[len(deckForTest) -1])
	}
	
}


func TestSaveToDeckandNewDeckFromFile(t *testing.T)  {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck:= newDeckFromFile(("_decktesting"))

	if len(loadedDeck)!=52 {
		t.Errorf("Expected 52 cards in deck but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}





