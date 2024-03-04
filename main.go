package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGen() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(50)
	return randomNumber
}

func getPlayerNumber() int {
	var playerNumber int
	fmt.Println("Please input a number here:")
	fmt.Println("Let the number be greater than 0 but less than 50, for it pleases the gods...")
	fmt.Scan(&playerNumber)
	fmt.Printf("Your number is... %d \n", playerNumber)
	return playerNumber
}

func compareGuesses(randomNumber int, playerNumber int) {
	fmt.Println("Let us compare the numbers...")

	if randomNumber == playerNumber {
		fmt.Println("You and The Machine have tied. The gods have decided.")
	} else if randomNumber > playerNumber {
		fmt.Println("The Machine has won the duel. The gods have decided.")
	} else if randomNumber < playerNumber {
		fmt.Println("The newcomer has won the duel. The gods have decided.")
	}
}

func main() {
	var randomNumber int = randomNumberGen()
	var playerNumber int = getPlayerNumber()
	compareGuesses(randomNumber, playerNumber)
}