package main

import (
	"testing"
)

func TestGetAvailableLocationForKing(t *testing.T) {
	myDict := createBoard()
	piece := newPiece(1, K, 'e', 4)
	myDict[piece.GetCurrentLocation()] = *piece
	result := piece.getAvailableLocation(myDict)

	count := len(result)
	if count != 8 {
		t.Errorf("got %d.", count)
		t.Errorf("Total available Location is 8.")
	}

	if result[0] != "f4" {
		t.Error("Available location invalid.")
	}
}

func TestGetAvailableLocationForKing_EdgeCase(t *testing.T) {
	myDict := createBoard()
	piece := newPiece(1, K, 'h', 4)
	myDict[piece.GetCurrentLocation()] = *piece

	result := piece.getAvailableLocation(myDict)

	count := len(result)
	if count != 5 {
		t.Errorf("Got %d.", count)
		t.Error("Total available location is 5.")
	}
}

func TestGetAvailableLocationForKing_ExistingPieceAroundMainPiece(t *testing.T) {
	myDict := createBoard()
	piece := newPiece(1, K, 'e', 4)
	piece2 := newPiece(1, Q, 'f', 4)
	myDict[piece.GetCurrentLocation()] = *piece
	myDict[piece2.GetCurrentLocation()] = *piece2

	result := piece.getAvailableLocation(myDict)

	count := len(result)
	if count != 7 {
		t.Errorf("Got %d.", count)
		t.Error("Total available location is 7.")
	}

}
