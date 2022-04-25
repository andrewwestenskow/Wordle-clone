package logic

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/andrewwestenskow/Wordle-clone/constants"
)

// getRandomWord gets a word from the available bank
// returns a string (the game word)
func getRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(constants.Words))
	return constants.Words[index]
}

// initializeGameState creates a new blank game state with an empty word map and a new game word
func initializeGameState(gameType string) GameState {
	word := getRandomWord()

	state := GameState{
		word: word,
		Guesses: []string{},
		WordMap: [][]string{
			{"_","_","_","_","_"},
			{"_","_","_","_","_"},
			{"_","_","_","_","_"},
			{"_","_","_","_","_"},
			{"_","_","_","_","_"},
		},
		Win: false,
		gameType: gameType,
		NumGuesses: 0,
	}

	return state
}

// InitializeGame creates a new game state, assigns it to the global variable and initiates the game based on the type of game passed in "console" or "browser"
func InitializeGame(gameType string) {
	Game = initializeGameState(gameType)

	if(Game.gameType == "console"){
		fmt.Println("I'm thinking of a 5 letter word......")
		consoleTurn()
	} else if Game.gameType == "browser" {
		fmt.Println("To play, open http://localhost:4000")
		initializeBrowserGame()
	}

}