package main

import (
	"net/http"
	"strings"
)

// Fonction pour jouer au Hangman
func playHangman(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	guess := strings.ToLower(strings.TrimSpace(r.FormValue("letter")))

	if guess == "" {
		handleInvalidGuess(w, "Vous devez entrer une lettre ou un mot.")
		return
	}

	if len(guess) > 1 {
		handleFullWordGuess(w, guess)
		return
	}

	handleSingleLetterGuess(w, guess)
}
