package main

import (
	"net/http"
)

// Fonction pour gérer la page Hangman
func handlerHangman(w http.ResponseWriter, r *http.Request) {
	page = "Hangman"
	if r.Method == http.MethodPost {
		playHangman(w, r)
		return
	}
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
		Message:  "",
	}
	err = tmpl.ExecuteTemplate(w, "hangman.html", data)
	if err != nil {
		http.Error(w, "Fail dans le chargement de la page :", http.StatusInternalServerError)
		return
	}
}

