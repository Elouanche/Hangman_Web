package main

import "strings"

// Cette fonction vérifie si un mot donné est présent dans une chaîne de mots.

func containsGuess(allguess string, guess string) bool {
	mots := strings.Split(allguess, "-")

	// Parcourt les mots pour trouver une correspondance.
	for _, mot := range mots {
		// Vérifie si le mot actuel correspond au mot recherché.
		if mot == guess {
			return true
		}
	}
	return false
}
