package logic

import (
	"fmt"
	"strings"
)

func printWordMap(){
	for i := len(Game.wordMap) - 1; i>=0; i-- {
		if Game.numGuesses >= i + 1 {
			fmt.Printf("%d. %[2]s %[3]s\n", i + 1, Game.wordMap[i], strings.Split(Game.Guesses[i], ""))
			} else {
			fmt.Printf("%d. %s\n", i + 1, Game.wordMap[i])
		}
	}
}

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
			fmt.Printf("You did it!  It only took you %d guesses!\n", Game.numGuesses)
			playAgain()
		} else if Game.numGuesses < 5 {
			consoleTurn()
		} else {
			fmt.Printf("You lost, better luck next time.  Your word was %s\n", Game.word)
			playAgain()
		}
	}

}