package main

import "net/http"

// Fonction pour gérer une devinette de mot complet
func handleFullWordGuess(w http.ResponseWriter, guess string) {
	if len(guess) != len(word) {
		handleInvalidGuess(w, "Le mot que vous avez entré n'a pas la même longueur que le mot à deviner.")
		return
	}

	if guess == word {
		score++
		err := tmpl.ExecuteTemplate(w, "win.html", nil)
		if err != nil {
			http.Error(w, "Fail dans le chargement de la page :", http.StatusInternalServerError)
		}
		return
	}

	attempts--
	handleGameUpdate(w, "Le mot complet que vous avez deviné est incorrect !")
}
