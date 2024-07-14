package main

import "net/http"

// Fonction pour gérer une devinette de lettre individuelle
func handleSingleLetterGuess(w http.ResponseWriter, guess string) {
	if !isAlpha(guess) {
		handleInvalidGuess(w, "Vous ne pouvez pas faire ça !")
		return
	}

	if containsGuess(allguess, guess) {
		handleInvalidGuess(w, "Cette lettre est déjà essayé !")
		return
	}

	if isHintLetter(guess) {
		handleInvalidGuess(w, "Cette lettre est déjà donnée comme indice !")
		return
	}

	allguess += guess + "-"
	found := false

	for i, char := range word {
		if guess == string(char) {
			guessedLetters[i] = true
			found = true
		}
	}

	if !found {
		attempts--
		handleInvalidGuess(w, "Mauvaise lettre !")
		return
	}

	handleGameUpdate(w, "")
}
