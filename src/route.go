package main

import (
	"fmt"
	"log"
	"net/http"
)

// Permet de faire les routes vers les differentes pages.
func route(h1, h2, h3, h4, h5, h6 http.HandlerFunc) {
	http.HandleFunc("/", h1)
	http.HandleFunc("/hangman", h2)
	http.HandleFunc("/Menu", h3)
	http.HandleFunc("/Easy", h4)
	http.HandleFunc("/Medium", h5)
	http.HandleFunc("/Hard", h6)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
