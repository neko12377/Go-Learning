package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck number of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Club" {
		t.Errorf("Expected last card is Four of Club, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	errOfRemove := os.Remove("_deckTesting")
	if errOfRemove != nil {
		fmt.Println(errOfRemove)
	}

	deck := newDeck()
	errOfSaveFile := deck.saveToFile("_deckTesting")
	if errOfSaveFile != nil {
		t.Errorf("failed to add file")
	}

	loadedDeck := newDeckFromFile("_deckTesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}
	errOfRemove2 := os.Remove("_deckTesting")
	if errOfRemove2 != nil {
		t.Errorf("failed to remove file")
	}
}