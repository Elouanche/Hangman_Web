package main

//ce programme permets de garder en mémoire les lettres déjà deviné
func allLettersGuessed(guessedLetters []bool) bool {
	for _, guessed := range guessedLetters {
		if !guessed {
			return false
		}
	}
	return true
}
