package main

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

//	type Position struct {
//		file rune // Columns
//		rank int  // Rows
//	}
type Piece struct {
	pieceId   uuid.UUID
	playerId  int
	pieceType PieceType
	file      rune // Columns
	rank      int  // Rows
	// position  Position
	prevLoc string
	isDead  bool
}

func newPiece(id int, pieceType PieceType, newFile rune, newRank int) *Piece {
	p := Piece{pieceId: uuid.New(), playerId: id, pieceType: pieceType, file: newFile, rank: newRank, prevLoc: "", isDead: false}
	return &p
}
func (p Piece) GetCurrentLocation() string {
	return fmt.Sprintf("%c", p.file) + strconv.Itoa(p.rank)
}
func (p Piece) printPiece() {
	fmt.Printf("PieceID:%s, PlayerID: %d, Type: %s, Current Location:%c, %d \n", p.pieceId, p.playerId, p.pieceType, p.file, p.rank)
}

func getKingAvailableLocation(board map[string]Piece, file rune, rank int) []string {
	var availLoc []string

	newFile := file + 1
	newRank := rank

	// How to get the value for the
	if newFile <= 'h' && newFile > 'a' && newRank <= 8 && newRank > 0 {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}
	newFile = file - 1
	if newFile <= 'h' && newFile > 'a' && newRank <= 8 && newRank > 0 {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}
	newFile = file
	newRank = rank + 1
	if newFile <= 'h' && newFile > 'a' && newRank <= 8 && newRank > 0 {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}
	newFile = file
	newRank = rank - 1
	if newFile <= 'h' && newFile > 'a' && newRank <= 8 && newRank > 0 {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}
	newFile = file + 1
	newRank = rank + 1
	if newFile <= 'h' && newFile > 'a' && newRank <= 8 && newRank > 0 {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}
	newFile = file + 1
	newRank = rank - 1
	if newFile <= 'h' && newFile > 'a' && newRank <= 8 && newRank > 0 {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}
	newFile = file - 1
	newRank = rank + 1
	if newFile <= 'h' && newFile > 'a' && newRank <= 8 && newRank > 0 {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}
	newFile = file - 1
	newRank = rank - 1
	if newFile <= 'h' && newFile > 'a' && newRank <= 8 && newRank > 0 {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}
	return availLoc
}

func getQueenAvailableLocation(file rune, rank int) []string {
	var availLoc []string
	return availLoc
}
func getKnightAvailableLocation(file rune, rank int) []string {
	var availLoc []string
	return availLoc
}
func getBishopAvailableLocation(file rune, rank int) []string {
	var availLoc []string
	return availLoc
}

func (p Piece) getAvailableLocation(board map[string]Piece) []string {
	var availLoc []string
	if p.pieceType == K {
		availLoc = getKingAvailableLocation(board, p.file, p.rank)
	} else if p.pieceType == Q {
		availLoc = getQueenAvailableLocation(p.file, p.rank)
	} else if p.pieceType == N {
		availLoc = getKnightAvailableLocation(p.file, p.rank)
	} else if p.pieceType == B {
		availLoc = getBishopAvailableLocation(p.file, p.rank)
	} else {

	}
	return availLoc
}
func isAvailableMove(newFile rune, newRank int) bool {
	return true
}

func (p *Piece) updateLocation(newFile rune, newRank int) {

	if p.pieceType == K {
		isPossible := isAvailableMove(newFile, newRank)
		if isPossible == true {
			// TODO Update the current location.
		}
		p.file += 1
		p.file -= 1

	} else if p.pieceType == Q {

	} else if p.pieceType == B {

	} else if p.pieceType == N {

	} else if p.pieceType == R {

	} else if p.pieceType == P {

	} else {
		fmt.Println("Invali piece type.")
	}
}

type PieceType string

const (
	K = "King"
	Q = "Queen"
	B = "Bishop"
	N = "Knight"
	R = "Rook"
	P = "Pawn"
)
