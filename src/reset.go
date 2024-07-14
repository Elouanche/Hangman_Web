package main

import "math/rand"

// Cette fonction réinitialise les variables du jeu, réinitialise les lettres devinées, sélectionne aléatoirement une lettre à deviner, et met à jour la variable cachée `WordHide`.
func resetGameVariables() {
	for i := range guessedLetters {
		guessedLetters[i] = false
	}
	Ms = ""
	WordHide = ""
	guess = "0"
	attempts = 10
	allguess = ""
	guessedLetters = make([]bool, len(word))

	// Sélectionner une lettre aléatoire comme indice
	letterRandom := rand.Intn(len(word))
	hintLetter = string(word[letterRandom])
	guessedLetters = initializeGuessedLetters(word, hintLetter)

	WordHide = rewriting(word, guessedLetters)
}

// Fonction pour initialiser les lettres devinées avec toutes les occurrences de l'indice
func initializeGuessedLetters(word, hintLetter string) []bool {
	guessedLetters := make([]bool, len(word))
	for i, char := range word {
		if string(char) == hintLetter {
			guessedLetters[i] = true
		}
	}
	return guessedLetters
}
