package main

import "net/http"

/* Cette fonction permet d'accéder aux requêtes d'accueil et de les gérer graçe à un formulaire*/
func handlerAcceuil(w http.ResponseWriter, r *http.Request) {
	page = "Acceuil"
	data := struct {
		Word string
		Page string
	}{
		Word: WordHide,
		Page: page,
	}
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "Fail dans le chargement de la page :", http.StatusInternalServerError)
		return
	}
}
