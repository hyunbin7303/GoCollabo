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
	if !(id == 1 || id == 2) {
		panic("Player ID should be 1 or 2.")
	}
	p := Piece{pieceId: uuid.New(), playerId: id, pieceType: pieceType, file: newFile, rank: newRank, prevLoc: "", isDead: false}
	return &p
}
func (p Piece) GetCurrentLocation() string {
	return fmt.Sprintf("%c", p.file) + strconv.Itoa(p.rank)
}
func (p Piece) printPiece() {
	fmt.Printf("PieceID:%s, PlayerID: %d, Type: %s, Current Location:%c, %d \n", p.pieceId, p.playerId, p.pieceType, p.file, p.rank)
}

//	func isValidLoc(board map[string]Piece, playerId int, newFile rune, newRank int) bool {
//		if newFile <= 'h' && newFile >= 'a' && newRank <= 8 && newRank > 0 {
//			checkLoc := fmt.Sprintf("%c", newFile) + strconv.Itoa(newRank)
//			if board[checkLoc].playerId != playerId {
//				return true
//			}
//		}
//		return false
//	}
func isValidLoc(board map[string]Piece, playerId int, newFile rune, newRank int) (bool, Piece) {
	if newFile <= 'h' && newFile >= 'a' && newRank <= 8 && newRank > 0 {
		checkLoc := fmt.Sprintf("%c", newFile) + strconv.Itoa(newRank)
		piece := board[checkLoc]
		if piece.playerId != playerId {
			return true, piece
		}
	}
	return false, Piece{isDead: true}
}

func (p *Piece) getKingAvailableLocation(board map[string]Piece) []string {
	var availLoc []string

	newFile := p.file + 1
	newRank := p.rank
	isValidLocResult, checkPiece := isValidLoc(board, p.playerId, newFile, newRank)
	if isValidLocResult && checkPiece.playerId != p.playerId {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}

	newFile = p.file - 1
	isValidLocResult, checkPiece = isValidLoc(board, p.playerId, newFile, newRank)
	if isValidLocResult && checkPiece.playerId != p.playerId {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}

	newFile = p.file
	newRank = p.rank + 1
	isValidLocResult, checkPiece = isValidLoc(board, p.playerId, newFile, newRank)
	if isValidLocResult && checkPiece.playerId != p.playerId {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}

	newFile = p.file
	newRank = p.rank - 1
	isValidLocResult, checkPiece = isValidLoc(board, p.playerId, newFile, newRank)
	if isValidLocResult && checkPiece.playerId != p.playerId {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}

	newFile = p.file + 1
	newRank = p.rank + 1
	isValidLocResult, checkPiece = isValidLoc(board, p.playerId, newFile, newRank)
	if isValidLocResult && checkPiece.playerId != p.playerId {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}

	newFile = p.file + 1
	newRank = p.rank - 1
	isValidLocResult, checkPiece = isValidLoc(board, p.playerId, newFile, newRank)
	if isValidLocResult && checkPiece.playerId != p.playerId {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}

	newFile = p.file - 1
	newRank = p.rank + 1
	isValidLocResult, checkPiece = isValidLoc(board, p.playerId, newFile, newRank)
	if isValidLocResult && checkPiece.playerId != p.playerId {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}

	newFile = p.file - 1
	newRank = p.rank - 1
	isValidLocResult, checkPiece = isValidLoc(board, p.playerId, newFile, newRank)
	if isValidLocResult && checkPiece.playerId != p.playerId {
		availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
	}
	return availLoc
}

func updateAvailableLoc() {

}

func (p *Piece) getQueenAvailableLocation(board map[string]Piece) []string {
	var availLoc []string
	newFile := p.file
	newRank := p.rank
	// Moving to right direction
	for true {
		newFile += 1
		isValidLocResult, checkPiece := isValidLoc(board, p.playerId, newFile, newRank)
		if isValidLocResult {
			availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
			if !checkPiece.isDead && checkPiece.playerId != p.playerId {
				break
			}
			continue
		}
		break
	}

	// Moving to left direction
	newFile = p.file
	for true {
		newFile -= 1
		isValidLocResult, checkPiece := isValidLoc(board, p.playerId, newFile, newRank)
		if isValidLocResult {
			availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
			if !checkPiece.isDead && checkPiece.playerId != p.playerId {
				break
			}
			continue
		}
		break
	}
	// moving to below direction
	newFile = p.file
	newRank = p.rank
	for true {
		newRank -= 1
		isValidLocResult, checkPiece := isValidLoc(board, p.playerId, newFile, newRank)
		if isValidLocResult {
			availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
			if !checkPiece.isDead && checkPiece.playerId != p.playerId {
				break
			}
			continue
		}
		break
	}

	// Moving to upper direction
	newFile = p.file
	newRank = p.rank
	for true {
		newRank += 1
		isValidLocResult, checkPiece := isValidLoc(board, p.playerId, newFile, newRank)
		if isValidLocResult {
			availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
			if !checkPiece.isDead && checkPiece.playerId != p.playerId {
				break
			}
			continue
		}
		break
	}

	// Moving to upper right direction
	newFile = p.file
	newRank = p.rank
	for true {
		newRank += 1
		newFile += 1
		isValidLocResult, checkPiece := isValidLoc(board, p.playerId, newFile, newRank)
		if isValidLocResult {
			availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
			if !checkPiece.isDead && checkPiece.playerId != p.playerId {
				break
			}
			continue
		}
		break
	}

	// Moving to upper left direction
	newFile = p.file
	newRank = p.rank
	for true {
		newRank += 1
		newFile -= 1
		isValidLocResult, checkPiece := isValidLoc(board, p.playerId, newFile, newRank)
		if isValidLocResult {
			availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
			if !checkPiece.isDead && checkPiece.playerId != p.playerId {
				break
			}
			continue
		}
		break
	}

	// Moving to below left direction
	newFile = p.file
	newRank = p.rank
	for true {
		newRank -= 1
		newFile -= 1
		isValidLocResult, checkPiece := isValidLoc(board, p.playerId, newFile, newRank)
		if isValidLocResult {
			availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
			if !checkPiece.isDead && checkPiece.playerId != p.playerId {
				break
			}
			continue
		}
		break
	}

	// Moving to below right direction
	newFile = p.file
	newRank = p.rank
	for true {
		newRank -= 1
		newFile += 1
		isValidLocResult, checkPiece := isValidLoc(board, p.playerId, newFile, newRank)
		if isValidLocResult {
			availLoc = append(availLoc, fmt.Sprintf("%c", newFile)+strconv.Itoa(newRank))
			if !checkPiece.isDead && checkPiece.playerId != p.playerId {
				break
			}
			continue
		}
		break
	}

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

// TODO
func (p Piece) getAvailableLocation(board map[string]Piece) []string {
	var availLoc []string
	if p.pieceType == K {
		availLoc = p.getKingAvailableLocation(board)
	} else if p.pieceType == Q {
		availLoc = p.getQueenAvailableLocation(board)
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
