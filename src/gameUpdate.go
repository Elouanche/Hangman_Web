package main

import "net/http"

// Fonction pour gérer la mise à jour du jeu
func handleGameUpdate(w http.ResponseWriter, additionalMessage string) {
	if allLettersGuessed(guessedLetters) {
		score++
		err := tmpl.ExecuteTemplate(w, "win.html", nil)
		if err != nil {
			http.Error(w, "Fail dans le chargement de la page :", http.StatusInternalServerError)
		}
		return
	}

	if attempts <= 0 {
		err := tmpl.ExecuteTemplate(w, "defeat.html", nil)
		if err != nil {
			http.Error(w, "Fail dans le chargement de la page :", http.StatusInternalServerError)
		}
		return
	}

	WordHide = rewriting(word, guessedLetters)
	message := additionalMessage
	handleInvalidGuess(w, message)
}
