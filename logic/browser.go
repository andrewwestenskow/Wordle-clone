package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type guessBody struct {
	Guess string
}

type secret struct {
	Word string
}

// serveGameState sends the current game state
// accepts no parameters
// serves the current game state
func serveGameState(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Game)
}


// handler serves up the boilerplate html stored in the views folder
// serves the html file in that folder
func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./views/index.html")
}

// resetGame handles resetting the game state
// serves a blank game state
func resetGame(w http.ResponseWriter, r *http.Request){
	state := initializeGameState("browser")
	Game = state
	serveGameState(w, r)
}

// guess accepts a guessed word on the request body and verifies it
// serves an updated game state
func guess(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var b guessBody
	err := decoder.Decode(&b)
	if err != nil {
		fmt.Println("There was an error parsing your request")
	}
	

	handleGuess(b.Guess)

	serveGameState(w, r)
}

// getSecretWord sends the secret game word
// serves a string
func getSecretWord(w http.ResponseWriter, r *http.Request){
	word := Game.word

	secret := secret{Word: word}
	json.NewEncoder(w).Encode(secret)
}

// initializeBrowserGame sets up all of the necessary endpoints and listens for incoming requests
// serves nothing
func initializeBrowserGame(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/game", serveGameState)
	http.HandleFunc("/guess", guess)
	http.HandleFunc("/reset", resetGame)
	http.HandleFunc("/word", getSecretWord)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
