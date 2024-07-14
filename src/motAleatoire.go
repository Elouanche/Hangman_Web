package main

import (
	"fmt"
	"math/rand"
)

// Permets de choisir un mot aléatoire
func chooseRandomWord(words []string) string {
    // Vérifie si la slice de mots est vide.
    if len(words) == 0 {
        // Affiche un message d'erreur si la slice est vide.
        fmt.Println("Erreur: le fichier words est vide.")
        // Retourne une chaîne vide ou toute autre logique pour gérer le cas où il n'y a pas de mots.
        return ""
    }

    // Génère un index aléatoire en utilisant la fonction Intn du package rand.
    randIndex := rand.Intn(len(words))
    
    // Sélectionne le mot correspondant à l'index aléatoire.
    word := words[randIndex]
    
    // Retourne le mot sélectionné.
    return word
}