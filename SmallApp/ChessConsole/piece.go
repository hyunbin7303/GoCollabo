package main

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type Piece struct {
	pieceId   uuid.UUID
	playerId  int
	pieceType PieceType
	file      rune // Columns
	rank      int  // Rows
	prevLoc   string
	isDead    bool
}

func newPiece(id int, pieceType PieceType, newFile rune, newRank int) *Piece {
	p := Piece{pieceId: uuid.New(), playerId: id, pieceType: pieceType, file: newFile, rank: newRank, prevLoc: "", isDead: false}
	return &p
}
func (p Piece) GetCurrentLocation() (file rune, rank int) {
	return p.file, p.rank
}
func (p Piece) printPiece() {
	fmt.Printf("PieceID:%s, PlayerID: %d, Type: %s, Current Location:%c, %d \n", p.pieceId, p.playerId, p.pieceType, p.file, p.rank)
}
func (p Piece) printAvailableLocation() {
	// TODO In the future.
}

// TODO : Get Available location based on the current loc.
func (p Piece) getAvailableLocation() []string {
	var availLoc []string
	if p.pieceType == K {
		// TODO Total 8 directions.
		newFile := p.file + 1
		newRank := p.rank
		// TODO Validation
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))

		newFile = p.file - 1
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))

		newFile = p.file
		newRank = p.rank + 1
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))

		newFile = p.file
		newRank = p.rank - 1
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))

		fmt.Println(availLoc)
	}
	// Need to return List?
	return nil
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
