package logic

import (
	"errors"
	"strings"
)

type GameState struct {
	word string
	Guesses []string
	wordMap [][]string
	gameType string
	Win bool
	numGuesses int
}

var Game GameState

// handleGuess accepts a string and compares it to the assigned game word
// returns a boolean indicating if the guess is correct or not
func handleGuess(guess string) (bool, error){
	if(len(guess) == 5){
		Game.Guesses = append(Game.Guesses, guess)
		Game.numGuesses++
	
		if guess == Game.word {
			return true, nil
		} else {
			wordMap := generateWordMap(guess)
			Game.wordMap[Game.numGuesses - 1] = wordMap
			return false, nil
		}
	} else {
		return false, errors.New("invalid guess length")
	}
}

func verifyLetter(word []string, letter string, index int) (bool, bool) {
	if word[index] == letter {
		return true, true
	}

	for i, l := range word {
		if l == letter && i!= index {
			return true, false
		} else {
			continue
		}
	}
	return false, false
}

func generateWordMap(guess string) []string {
	wordSlice := strings.Split(Game.word, "")
	guessSlice := strings.Split(guess, "")

	for i, l := range guessSlice {
		correctLetter, correctIndex := verifyLetter(wordSlice, l, i)

		if correctLetter && correctIndex {
			guessSlice[i] = "O"
		} else if correctLetter && !correctIndex {
			guessSlice[i] = "I"
		} else {
			guessSlice[i] = "X"
		}
	}
	return guessSlice
}