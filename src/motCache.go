package main

// recris le mots en chacahnt les lettre non trouver
func rewriting(word string, guessedLetters []bool) string {
	// Variable pour stocker le mot caché en cours de construction.
	WordHide := ""
	// Parcours chaque caractère du mot à deviner (word).
	for i, char := range word {
		// Vérifie si la lettre à la position i a été devinée.
		if guessedLetters[i] {
			// Si la lettre a été devinée, ajoute la lettre à la variable WordHide.
			WordHide += string(char)
		} else {
			// Si la lettre n'a pas été devinée, ajoute un tiret bas (_) suivi d'un espace à WordHide.
			WordHide += "_ "
		}
	}
	// Retourne le mot caché construit.
	return WordHide
}
