package main

import (
	"fmt"
	"strconv"
)

func createBoard() map[string]Piece {

	myDict := make(map[string]Piece)
	for i := 8; i >= 1; i-- {
		for j := 'a'; j <= 'h'; j++ {
			temp := fmt.Sprintf("%c", j) + strconv.Itoa(i)
			myDict[temp] = Piece{isDead: true}
		}
	}
	return myDict
}

func main() {

	// basic board setup
	myDict := createBoard()

	// Player setup
	// manually setup the pieces location.
	// player1, err := newPlayer(1)
	// player2, err := newPlayer(2)
	// fmt.Println(player1.pieces)
	// fmt.Println(player2.pieces)

	player1_k := newPiece(1, K, 'e', 8)
	player1_q := newPiece(1, Q, 'd', 8)
	player1_b := newPiece(1, B, 'c', 8)
	player1_b2 := newPiece(1, B, 'f', 8)
	myDict["d8"] = *player1_q
	myDict["e8"] = *player1_k
	myDict["c8"] = *player1_b
	myDict["f8"] = *player1_b2

	player2_k := newPiece(2, K, 'd', 1)
	player2_q := newPiece(2, Q, 'e', 1)
	myDict["d1"] = *player2_q
	myDict["e1"] = *player2_k

	fmt.Println("[Player1]Pick the piece")
	var input string
	fmt.Scan(&input)
	fmt.Println("[Player1] You entered Piece location:", input)
	piece, ok := myDict[input]
	if ok {
		piece.printPiece()
		availableMap := piece.getAvailableLocation(myDict)
		fmt.Println("Get Available Location:")
		fmt.Println(availableMap)
		if piece.playerId == 1 {
			fmt.Println("Enter movement location: ")
			_, err := fmt.Scan(&input)
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

	// fmt.Println("[Player2]Pick the piece")
	// _, err = fmt.Scan(&input)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("[Player2] You entered Piece location:", input)
	// piece, ok = myDict[input]
	// if ok {
	// 	piece.printPiece()
	// }
}
