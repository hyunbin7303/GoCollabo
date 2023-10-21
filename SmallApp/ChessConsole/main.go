package main

import (
	"fmt"
	"strconv"
)

func main() {

	// basic board setup
	myDict := make(map[string]Piece)
	for i := 8; i >= 1; i-- {
		for j := 'a'; j <= 'h'; j++ {
			temp := fmt.Sprintf("%c", j) + strconv.Itoa(i)
			myDict[temp] = Piece{}
		}
	}

	// Player setup
	// manually setup the pieces location.
	player1 := Player{playerId: 1, numOfTurn: 0}
	player2 := Player{playerId: 2, numOfTurn: 0}

	player1_k := newPiece(1, K, 'e', 8)
	myDict["e8"] = *player1_k
	// fmt.Println(myDict)

	fmt.Println(player1)
	fmt.Println(player2)

	fmt.Println("[Player1]Pick the piece")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Need to verify if Piece Exists in that tile.
	fmt.Println("[Player1] You entered Piece location:", input)
	piece, ok := myDict[input]
	if ok {
		piece.printPiece()
		if piece.playerId == 1 {
			fmt.Println("Enter movement location: ")
			_, err = fmt.Scan(&input)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			_, ok := myDict[input]
			if ok {
				fmt.Println("Avaliable Location.")

			} else {
				fmt.Print("Not available. Please try again.")
			}

		} else {
		}

	} else {
		fmt.Println("Please check the tile")
	}

	// TODO Set up the King's location

	// Game Engine setup
	// while loop for playing game until one user loses (checkmate or king died)
}
