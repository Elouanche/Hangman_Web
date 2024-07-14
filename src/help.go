package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// cette fonction permets d'afficher les aides essentiels pour le jeu
func help() {
	helpFlag := flag.Bool("h", false, "Affiche les explications")

	flag.Parse()

	// Vérification de l'option -h
	if *helpFlag {
		printHelp()
		os.Exit(0)
	}
}

// Fonction qui permets de clear le cmd
func ClearConsole() {
	var cmd *exec.Cmd

	// Choix de la commande en fonction du système d'exploitation
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "linux", "darwin":
		cmd = exec.Command("clear")
	default:
		fmt.Println("Système d'exploitation non pris en charge pour le nettoyage de la console.")
		return
	}

	// Exécution de la commande
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// fonction qui permets d'afficher les aides
func printHelp() {
	ClearConsole()
	fmt.Println("Bienvenue dans Hangman Web - Le Jeu du Pendu en Ligne. Hangman Web est un jeu du Pendu conçu pour être joué sur Internet.")
	fmt.Println("\nCe jeu a été créé dans le but d'apprendre à utiliser le langage de programmation Go et de mettre en œuvre un modèle client-serveur.")

	fmt.Println("\nPour jouer au jeu, suivez ces étapes :")
	fmt.Println("1. Dans votre terminal, exécutez :\n    make build\n  puis\n   make run\n   ou bien\n    .\\Hangman_Web.exe")
	fmt.Println("2. Ouvrez votre navigateur web et rendez-vous à l'adresse 'localhost:8080'")
	fmt.Println("Si vous avez compilé et que vous faites des modifications vous pouvez tout de même faire :\n make restart")

	fmt.Println("\nConditions d'utilisation :")
	fmt.Println("Après avoir exécuté make build, ne modifiez pas le fichier index.html et ne supprimez pas les photos, cela pourrait causer des problèmes.")
	fmt.Println("Vous pouvez ajouter des mots dans les fichiers words.txt pour enrichir le jeu.")
	fmt.Println("Avant de modifier le programme, exécutez 'make clean'.")

	fmt.Println("\nObjectif du Jeu :")
	fmt.Println("L'objectif est de deviner le mot caché en sélectionnant des lettres. Le jeu propose trois niveaux de difficulté pour tester vos compétences.")

	fmt.Println("\nNiveaux de Difficulté :")
	fmt.Println("Facile : Les mots à deviner sont courts et simples.")
	fmt.Println("Moyen : Les mots sont un peu plus complexes.")
	fmt.Println("Difficile : Les mots sont plus longs et plus difficiles.")

	fmt.Println("\nApprentissage de Go et Modèle Client-Serveur :")
	fmt.Println("Hangman Web a été développé comme un projet d'apprentissage pour maîtriser le langage de programmation Go. De plus, le jeu illustre l'implémentation d'un modèle client-serveur, où le serveur gère le jeu et le client interagit avec le joueur.")

	fmt.Println("\nAmusez-vous bien et bonne chance pour deviner les mots cachés ! 🎉✨")
}
