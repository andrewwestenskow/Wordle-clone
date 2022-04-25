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

func serveGameState(w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Game)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./views/index.html")
}

func getGameState(w http.ResponseWriter, r *http.Request){
	serveGameState(w)
}

func resetGame(w http.ResponseWriter, r *http.Request){
	state := initializeGameState("browser")
	Game = state
	serveGameState(w)
}

func guess(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var b guessBody
	err := decoder.Decode(&b)
	if err != nil {
		fmt.Println("There was an error parsing your request")
	}
	

	handleGuess(b.Guess)

	serveGameState(w)
}

func getSecretWord(w http.ResponseWriter, r *http.Request){
	word := Game.word

	secret := secret{Word: word}
	json.NewEncoder(w).Encode(secret)
}

func initializeBrowserGame(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/game", getGameState)
	http.HandleFunc("/guess", guess)
	http.HandleFunc("/reset", resetGame)
	http.HandleFunc("/word", getSecretWord)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
