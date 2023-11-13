package main

import "fmt"

type Player struct {
	playerId  int
	numOfTurn int
	pieces    []Piece
}

func newPlayer(id int) (Player, error) {
	if id == 1 {
		newPieces, err := getPiecesForPlayer1()
		if err != nil {
			return Player{}, fmt.Errorf("could not construct new Player: %v", err)
		}
		return Player{
			playerId: id,
			pieces:   newPieces,
		}, nil
	} else {
		newPieces, err := getPiecesForPlayer2()
		if err != nil {
			return Player{}, fmt.Errorf("could not construct new Player: %v", err)
		}
		return Player{
			playerId: id,
			pieces:   newPieces,
		}, nil
	}
}

func getPiecesForPlayer1() ([]Piece, error) {
	player_k := newPiece(1, K, 'e', 8)
	player_q := newPiece(1, Q, 'd', 8)

	pieces := []Piece{
		*player_k,
		*player_q,
	}
	return pieces, nil
}
func getPiecesForPlayer2() ([]Piece, error) {
	player_k := newPiece(2, K, 'd', 1)
	player_q := newPiece(2, Q, 'e', 1)

	pieces := []Piece{
		*player_k,
		*player_q,
	}
	return pieces, nil
}
