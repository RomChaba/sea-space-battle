package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Variable pour les Joueurs
var joueurs map[int]string
var joueurEnCours int

// Variable pour la fonction de retour JSON
var resp map[string][]byte

func main() {
	// Initialisation des variables
	joueurs = make(map[int]string)
	r := mux.NewRouter()
	joueur := r.PathPrefix("/joueur").Subrouter()
	jeux := r.PathPrefix("/jeux").Subrouter()

	// Gestion de la partie joueur
	joueur.HandleFunc("/liste", func(w http.ResponseWriter, r *http.Request) {

		if checkJoueur(joueurs[1]) && checkJoueur(joueurs[2]) {
			w = reponseFormatee(w, true, "Liste de tous les joueurs", joueurs)
		} else {
			w = reponseFormatee(w, false, "Il n'y a pas de joueurs.", nil)
		}

	})

	joueur.HandleFunc("/liste/{idJoueur}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idJoueur := vars["idJoueur"]
		i, err := strconv.Atoi(idJoueur)

		if err != nil {
			w = reponseFormatee(w, false, "Pb conversion id joueur", joueurs)
			return
		}

		if checkJoueur(joueurs[i]) {
			temp := make(map[int]string)
			temp[i] = joueurs[i]

			w = reponseFormatee(w, true, "Nom du joueur", temp[i])

		} else {
			w = reponseFormatee(w, false, "Le joueur n'existe pas !", nil)
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
			w = reponseFormatee(w, true, "Les 2 joueurs sont déjà renseigner", nil)
			return
		}

		w = reponseFormatee(w, true, fmt.Sprintf("Ajout de %s qui est le joueur %d\n", joueurs[numJoueur], numJoueur), joueurs[numJoueur])
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
		w = reponseFormatee(w, true, "TEST MAP", mapD)
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

type Sortie struct {
	Ok   bool
	Msg  string
	Data interface{}
}

func reponseFormatee(w http.ResponseWriter, ok bool, message string, dataIn interface{}) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")

	resp := Sortie{
		Ok:   ok,
		Msg:  message,
		Data: dataIn,
	}

	ret, _ := json.Marshal(resp)
	fmt.Fprintf(w, "%s", ret)

	return w
}
