package main

import (
	"fmt"
	"net/http"
)

// Cette fonction configure le jeu en chargeant les mots à partir d'un fichier spécifié, choisit un mot aléatoire, réinitialise les variables de jeu, et exécute le modèle.
func configureGame(filePath string, w http.ResponseWriter, r *http.Request) {
	words, err := readWords(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	ChooseWord(words)
	resetGameVariables()
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
		Message:  "", // Initialement vide
	}
	
	err = tmpl.ExecuteTemplate(w, "hangman.html", data)
	if err != nil {
		http.Error(w, "Fail dans le chargement de la page :", http.StatusInternalServerError)
		return
	}
}
