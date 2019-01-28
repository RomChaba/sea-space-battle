package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	joueur1 := map[string]int("Rifi": 1)
	joueur1 := map[string]int("Fifi": 2)

	joueurEnCours := 1

	// Gestion de la partie joueur

	joueur := r.PathPrefix("/joueur").Subrouter()
	jeux := r.PathPrefix("/jeux").Subrouter()

	joueur.HandleFunc("/liste", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Liste de tous les joueurs.\n")
	})
	joueur.HandleFunc("/liste/{idJoueur}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idJoueur := vars["idJoueur"]

		fmt.Fprintf(w, "id du joueur: %s\n", idJoueur)
	})
	joueur.HandleFunc("/name/{nomjoueur}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nomjoueur := vars["nomjoueur"]

		fmt.Fprintf(w, "Ajout un joueur avec le nom : %s\n", nomjoueur)
	})

	// Gestion de la partie jeux
	jeux.HandleFunc("/tour", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Return l'id du joueur dont c'est le tour.\n")
	})
	// Lieux du tir du joueur dont c'est le tour au format "B3" "E10"
	jeux.HandleFunc("/tir/{empl}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		empl := vars["empl"]
		fmt.Fprintf(w, "Le joueur [Nom du jour] à tiré ici %s.\n", empl)
	})
	//Fonction pour savoir a qui appartient une case
	jeux.HandleFunc("/case/{empl}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		empl := vars["empl"]
		fmt.Fprintf(w, "La case %s appartient à [NomDuJoueur].\n", empl)
	})
	//Fonction pour qu'un joueur place un navire a un emplacement
	jeux.HandleFunc("/placer/{idJoueur}/{horiVet}/{empl}", func(w http.ResponseWriter, r *http.Request) {
		// vars := mux.Vars(r)
		// idJoueur := vars["idJoueur"]
		// horiVet := vars["horiVet"]
		// empl := vars["empl"]

		mapD := map[string]int{"apple": 5, "lettuce": 7, "lettuces": 7, "lesttuce": 7, "lettucce": 7, "leettuce": 7}
		mapB, _ := json.Marshal(mapD)
		w.Header().Set("Content-Type", "text/JSON")
		fmt.Fprintf(w, "%s", mapB)
		// fmt.Fprintf(w, "Le joueur %s a placé son navire en %s a l'emplacement %s.\n", idJoueur, horiVet, empl)

	})

	http.ListenAndServe(":56700", r)
}
