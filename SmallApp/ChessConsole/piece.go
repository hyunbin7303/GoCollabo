package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Piece struct {
	pieceId    uuid.UUID
	playerId   int
	pieceType  PieceType
	currentLoc string
	prevLoc    string
	isDead     bool
}

func newPiece(playerId int, pieceType PieceType) Piece {
	p := Piece{pieceId: uuid.New(), playerId: playerId, pieceType: pieceType, currentLoc: "", prevLoc: "", isDead: false}
	return p
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

type PieceAction interface {
	getLocation() string
	UpdateLocation(Piece) error
}

func (p Piece) printPiece() {
	fmt.Printf("PieceID:%s, Type: %s, Current Location:%s", p.pieceId, p.pieceType, p.currentLoc)
}

// How to define rules for each pieces.
