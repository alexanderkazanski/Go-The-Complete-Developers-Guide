package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length to be 16 recieved %v", len(d))
	}

	if d[0] != "Hearts of Ace" {
		t.Errorf("Recieved %v Expected Hearts of Two", d[0])
	}

	if d[len(d)-1] != "Spades of Four" {
		t.Errorf("Expected Spades of Three recieved %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromDisk("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected a length of 16 recieved %v", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
