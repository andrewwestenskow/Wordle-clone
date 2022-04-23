package main

import (
	"fmt"

	"github.com/andrewwestenskow/Wordle-clone/logic"
)

func askForGameType() string{
	var gameType string
	
	fmt.Println("Would you like to play in the browser or the console?")
	fmt.Println("[1] Console")
	fmt.Println("[2] Browser")
	fmt.Scan(&gameType)
	
	if gameType == "1" {
		return "console"
	} else if gameType == "2" {
		return "browser"
	} else {
		fmt.Println("You have chosen poorly")
		panic("Incorrect choice")
	}
}
		
func main() {
	fmt.Println("Welcome to wordle!")

	gameType := askForGameType()

	logic.InitializeGame(gameType)
}