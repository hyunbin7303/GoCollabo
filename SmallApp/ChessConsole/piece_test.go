package main

import (
	"fmt"
	"testing"
)

func TestPlayerIdIsNot1or2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Player ID should be 1 or 2.", r)
		}
	}()
	newPiece(3, K, 'e', 4)
}

// func TestIsValidLoctiontesting(t *testing.T) {
// 	myDict := createBoard()
// 	newFile := 10
// 	newRank := 1
// 	piece := newPiece(1, K, 'e', 4)
// 	check := piece.isValidLoc(myDict, 1, newFile, newRank)

// 	if check {

// 	}
// }

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
func TestGetAvailableLocationForKing_SurroundedBySameTeam_NoAvailableSpot(t *testing.T) {
	myDict := createBoard()
	king := newPiece(1, K, 'e', 8)
	queen := newPiece(1, Q, 'd', 8)
	pawn1 := newPiece(1, P, 'd', 7)
	pawn2 := newPiece(1, P, 'e', 7)
	pawn3 := newPiece(1, P, 'f', 7)
	bishop := newPiece(1, B, 'f', 8)

	myDict[king.GetCurrentLocation()] = *king
	myDict[queen.GetCurrentLocation()] = *queen
	myDict[pawn1.GetCurrentLocation()] = *pawn1
	myDict[pawn2.GetCurrentLocation()] = *pawn2
	myDict[pawn3.GetCurrentLocation()] = *pawn3
	myDict[bishop.GetCurrentLocation()] = *bishop

	result := king.getAvailableLocation(myDict)

	count := len(result)
	if count != 0 {
		t.Errorf("Got %d.", count)
	}
}

func TestGetQueenAvailableLocation_OriginalLoc(t *testing.T) {
	myDict := createBoard()
	piece := newPiece(1, Q, 'd', 8)
	myDict[piece.GetCurrentLocation()] = *piece

	result := piece.getQueenAvailableLocation(myDict)

	count := len(result)
	if count != 21 {
		t.Errorf("Got %d.", count)
	}
}

func TestGetQueenAvailableLocation_OrignalLoc_AllPieceExist(t *testing.T) {
	myDict := createBoard()
	piece := newPiece(1, Q, 'd', 8)
	piece_k := newPiece(1, K, 'e', 8)
	piece_b := newPiece(1, B, 'c', 8)
	piece_p1 := newPiece(1, P, 'c', 7)
	piece_p2 := newPiece(1, P, 'd', 7)
	piece_p3 := newPiece(1, P, 'e', 7)

	myDict[piece.GetCurrentLocation()] = *piece
	myDict[piece_k.GetCurrentLocation()] = *piece_k
	myDict[piece_b.GetCurrentLocation()] = *piece_b
	myDict[piece_p1.GetCurrentLocation()] = *piece_p1
	myDict[piece_p2.GetCurrentLocation()] = *piece_p2
	myDict[piece_p3.GetCurrentLocation()] = *piece_p3

	result := piece.getQueenAvailableLocation(myDict)

	count := len(result)
	if count != 0 {
		t.Errorf("Got %d.", count)
	}
}

func TestGetQueenAvailableLocation_EnemyExistInAvailableLoc(t *testing.T) {
	myDict := createBoard()
	piece := newPiece(1, Q, 'd', 8)
	piece_enemy_k := newPiece(2, K, 'e', 8)

	myDict[piece.GetCurrentLocation()] = *piece
	myDict[piece_enemy_k.GetCurrentLocation()] = *piece_enemy_k

	result := piece.getQueenAvailableLocation(myDict)
	count := len(result)
	if count != 18 {
		t.Errorf("Got %d.", count)
	}

}
