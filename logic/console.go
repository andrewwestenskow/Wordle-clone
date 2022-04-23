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

func consoleTurn(){
	var guess string
	fmt.Print("Enter your guess: ")
	fmt.Scan(&guess)

	isCorrect := handleGuess(guess)

	printWordMap()

	if isCorrect {
		fmt.Printf("You did it!  It only took you %d guesses!", Game.numGuesses)
	} else if Game.numGuesses < 5 {
		consoleTurn()
	} else {
		fmt.Printf("You lost, better luck next time.  Your word was %s", Game.word)
	}
}