package logic

import (
	"errors"
	"strings"
)

type GameState struct {
	word string
	Guesses []string
	WordMap [][]string
	gameType string
	Win bool
	NumGuesses int
}

var Game GameState

// handleGuess accepts a string and compares it to the assigned game word
// returns a boolean indicating if the guess is correct or not
func handleGuess(guess string) (bool, error){
	if(len(guess) == 5){
		Game.Guesses = append(Game.Guesses, guess)
		Game.NumGuesses++
		wordMap := generateWordMap(guess)
		Game.WordMap[Game.NumGuesses - 1] = wordMap
	
		if guess == Game.word {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, errors.New("invalid guess length")
	}
}

// verifyLetter verifies a single letter in a guess
// accepts a split string (the game word), a single letter (to be verified), and an index(the letter's position in the guessed word)
// returns 2 booleans indicating the correctness of the guessed letters value and position
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

// generateWordMap creates a word map for the given guess
// accepts a split string (the guessed word)
// returns a slice containing 5 strings, each one representing a letter in the guessed word.  "O" indicates the letter is both correct and in the correct position, "I" indicates that the letter is correct but the position is wrong, and "X" indicates that the letter is incorrect
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