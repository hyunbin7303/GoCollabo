package main

import (
	"fmt"
	"strconv"
)

type Player struct {
	playerId  int
	numOfTurn int
}

func main() {
	playingSet := 0

	// basic board setup
	myDict := make(map[string]Piece)
	for i := 8; i >= 1; i-- {
		for j := 'a'; j <= 'h'; j++ {
			temp := fmt.Sprintf("%c", j) + strconv.Itoa(i)
			myDict[temp] = Piece{}
		}
	}

	// Player setup
	player1 := Player{playerId: 1, numOfTurn: 0}
	player2 := Player{playerId: 2, numOfTurn: 0}
	// TODO : player Pieces setup.

	// fmt.Println(myDict)
	fmt.Println(playingSet)
	fmt.Println(player1)
	fmt.Println(player2)

	// pieceTesting := Piece{}
	// p := Piece.x

	// Game Engine setup
	// while loop for playing game until one user loses (checkmate or king died)
}
