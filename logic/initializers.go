package logic

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/andrewwestenskow/Wordle-clone/constants"
)

func getRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(constants.Words))
	return constants.Words[index]
}

func initializeGameState(gameType string) GameState {
	word := getRandomWord()

	state := GameState{
		word: word,
		Guesses: []string{},
		wordMap: [][]string{
			{"_","_","_","_","_"},
			{"_","_","_","_","_"},
			{"_","_","_","_","_"},
			{"_","_","_","_","_"},
			{"_","_","_","_","_"},
		},
		Win: false,
		gameType: gameType,
		numGuesses: 0,
	}

	return state
}

func InitializeGame(gameType string) {
	Game = initializeGameState(gameType)

	if(Game.gameType == "console"){
		fmt.Println("I'm thinking of a 5 letter word......")
		consoleTurn()
	} else if Game.gameType == "browser" {
		fmt.Println("Coming soon.....")
	}

}