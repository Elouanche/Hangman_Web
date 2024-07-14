package main

import "unicode"

// Cette fonction vérifie si le joueur à bien envoyé une lettre
func isAlpha(str string) bool {
	for _, char := range str {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}
