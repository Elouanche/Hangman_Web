package main

import "net/http"

// Cette fonction sert à gérer la page du menu en définissant la page sur "Menu", préparant des données structurées pour le modèle, exécutant le modèle, et traitant les données de formulaire.
func handlerMenu(w http.ResponseWriter, r *http.Request) {
	page = "Menu"
	data := struct {
		Word string
		Page string
	}{
		Word: WordHide,
		Page: page,
	}
	err = tmpl.ExecuteTemplate(w, "menu.html", data)
	if err != nil {
		http.Error(w, "Fail dans le chargement de la page :", http.StatusInternalServerError)
		return
	}
}

// Ces fonctions configurent le jeu avec des mots spécifiques provenant de fichiers différents en fonction du niveau de difficulté. Permets donc d'avoir le bonus des difficultés
func handlerEasy(w http.ResponseWriter, r *http.Request) {
	page = "Hangman"
	configureGame("./ressources/words.txt", w, r)
}

func handlerMedium(w http.ResponseWriter, r *http.Request) {
	page = "Hangman"
	configureGame("./ressources/words2.txt", w, r)
}

func handlerHard(w http.ResponseWriter, r *http.Request) {
	page = "Hangman"
	configureGame("./ressources/words3.txt", w, r)
}
