package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var (
	err error
	tmpl           *template.Template
	page           = "la"
	Ms             = ""
	score          int
	attempts       int
	guess          = "0"
	lastGuess      = "8"
	allguess       string
	word           string
	WordHide       = ""
	guessedLetters []bool
	httpServer     *http.Server
	hintLetter     string
)

func main() {
	help()
	route(
		handlerAcceuil,
		handlerHangman,
		handlerMenu,
		handlerEasy,
		handlerMedium,
		handlerHard,
	)
}

func init() {
	tmpl, err = template.ParseGlob("./static/html/*.html")
	if err != nil {
		panic(err)
	}
}

// Cette fonction récupère et imprime le texte du formulaire pour le traitement ultérieur.
func handleFormData(r *http.Request) {
	text := strings.ToLower(r.FormValue("Ms"))
	fmt.Println(text)
}

