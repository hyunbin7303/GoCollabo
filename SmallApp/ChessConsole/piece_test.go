package main

import (
	"os"
	"testing"
)

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newPiece(1, K, 'e', 8)
	result := deck.getAvailableLocation()
	t.Log(result)

	os.Remove("_decktesting")
}
