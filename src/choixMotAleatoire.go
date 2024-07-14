package main

import (
	"io/ioutil"
	"strings"
)

// Lit les mots à partir d'un fichier et les retourne sous forme de slice.
func readWords(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	words := strings.Split(string(data), "\n")
	return words, nil
}

// Choisi un mot aléatoire à partir de la liste de mots et l'affiche.
func ChooseWord(words []string) {
	// Utilise la fonction chooseRandomWord pour obtenir un mot aléatoire.
	word = chooseRandomWord(words)

	// Vérifie et élimine le caractère de nouvelle ligne (\n) s'il est présent à la fin du mot.
	if strings.HasSuffix(word, "\n") {
		word = strings.TrimRight(word, "\n")
	}

	// Enlève les espaces entre les mots
	word = strings.TrimSpace(word)
}
