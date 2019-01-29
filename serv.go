package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

// Variable pour les Joueurs
var joueurs map[int]string
var joueurEnCours int

func main() {
	// Initialisation des variables
	joueurs = make(map[int]string)
	r := mux.NewRouter()
	joueur := r.PathPrefix("/joueur").Subrouter()
	jeux := r.PathPrefix("/jeux").Subrouter()

	// Gestion de la partie joueur

	joueur.HandleFunc("/liste", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Liste de tous les joueurs: \n")
		if checkJoueur(joueurs[1]) && checkJoueur(joueurs[2]) {

			var keys []int
			for k := range joueurs {
				keys = append(keys, k)
			}
			sort.Ints(keys)
			for _, k := range keys {
				fmt.Fprintf(w, "Joueur N°%d : %s\n", k, joueurs[k])
			}

			// for key, value := range joueurs {
			// 	fmt.Fprintf(w, "Joueur N°%d : %s\n", key, value)
			// }
			fmt.Fprintf(w, "JSON: \n")
			fmt.Fprintf(w, "%s", toJSON(joueurs))
		} else {
			fmt.Fprintf(w, "Il n'y a pas de joueurs.\n")
		}

	})

	joueur.HandleFunc("/liste/{idJoueur}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idJoueur := vars["idJoueur"]
		i, err := strconv.Atoi(idJoueur)

		if err != nil {
			fmt.Fprintf(w, "Pb conversion id joueur\n")
			return
		}

		if checkJoueur(joueurs[i]) {
			fmt.Fprintf(w, "Nom du joueur: %s\n", idJoueur)
			fmt.Fprintf(w, "JSON :\n")
			temp := make(map[int]string)
			temp[i] = joueurs[i]
			fmt.Fprintf(w, "%s\n", toJSON(temp))
		} else {
			fmt.Fprintf(w, "Le joueur n'existe pas !\n")
		}

	})
	// Fonction pour la création d'un joueur
	joueur.HandleFunc("/name/{nomjoueur}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nomjoueur := vars["nomjoueur"]
		numJoueur := 1

		if joueurs[1] == "" {
			joueurs[1] = nomjoueur
		} else if joueurs[2] == "" {
			joueurs[2] = nomjoueur
			numJoueur = 2
		} else {
			fmt.Fprintf(w, "Les 2 joueurs sont déjà renseigner")
			return
		}

		fmt.Fprintf(w, "Ajout de %s qui est le joueur %d\n", joueurs[numJoueur], numJoueur)
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

func checkJoueur(joueur string) bool {
	if joueur == "" {
		return false
	} else {
		return true
	}
}

func toJSON(j map[int]string) []byte {
	joueursJSON, _ := json.Marshal(j)
	return joueursJSON
}
