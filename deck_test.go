package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 , but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if l := d[len(d)-1]; l != "Ace of Spades" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", l)
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deck_test")
	d := newDeck()
	d.saveToFile("_deck_test")

	loadedDeck := newDeckFromFile("_deck_test")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_deck_test")
}
