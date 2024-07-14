package main

// Vérifiez si la lettre est déjà révélée en tant qu'indice
func isHintLetter(guess string) bool {
	for i, char := range word {
		if guess == string(char) && guessedLetters[i] {
			return true
		}
	}
	return false
}
