package main

import "net/http"

// Fonction pour gérer une devinette invalide
func handleInvalidGuess(w http.ResponseWriter, message string) {
	data := struct {
		Word     string
		Allguess string
		Score    int
		Page     string
		Attempts int
		Message  string
	}{
		Word:     WordHide,
		Allguess: allguess,
		Score:    score,
		Page:     page,
		Attempts: attempts,
		Message:  message,
	}
	err := tmpl.ExecuteTemplate(w, "hangman.html", data)
	if err != nil {
		http.Error(w, "Fail dans le chargement de la page :", http.StatusInternalServerError)
	}
}
