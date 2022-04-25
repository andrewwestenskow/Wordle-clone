package logic

import (
	"fmt"
	"strings"
)

// printWordMap prints the current wordmap to the console
// returns nothing
func printWordMap(){
	for i := len(Game.WordMap) - 1; i>=0; i-- {
		if Game.NumGuesses >= i + 1 {
			fmt.Printf("%d. %[2]s %[3]s\n", i + 1, Game.WordMap[i], strings.Split(Game.Guesses[i], ""))
			} else {
			fmt.Printf("%d. %s\n", i + 1, Game.WordMap[i])
		}
	}
}

// playAgain asks the player if they would like to play again
// should only accept "y" or "n" as input
// if "n" is input, the process exits
// returns nothing
func playAgain(){
	var ans string

	fmt.Scan(&ans)

	if ans == "y" {
		InitializeGame("console")
	} else if ans == "n" {
		fmt.Println("Thanks for playing!")
	} else {
		playAgain()
	}
}

// consoleTurn handles one round of the game in the console
// asks the player for their guess and passes the resulting input off for verification logic
// called recursively if there are remaining guesses available
// returns nothing
func consoleTurn(){
	var guess string
	fmt.Print("Enter your guess: ")
	fmt.Scan(&guess)

	isCorrect, err := handleGuess(guess)

	if err != nil {	
		fmt.Println("You don't count so well, try again using a 5 letter word")
		consoleTurn()
		} else {
		printWordMap()

		if isCorrect {
			fmt.Printf("You did it!  It only took you %d guesses!\n", Game.NumGuesses)
			playAgain()
		} else if Game.NumGuesses < 5 {
			consoleTurn()
		} else {
			fmt.Printf("You lost, better luck next time.  Your word was %s\n", Game.word)
			playAgain()
		}
	}

}