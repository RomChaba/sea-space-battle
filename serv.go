package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Variable pour les Joueurs
var joueurs map[int]string
var joueurEnCours int

// Variable pour le jeux
var emplacement []string
var nbBateau map[int]int

// Variable pour la fonction de retour JSON
var resp map[string][]byte

type Sortie struct {
	Ok   bool
	Msg  string
	Data interface{}
}

func main() {
	var carteJ1 = [5][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	var carteJ2 = [5][5]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	// var i, j int

	// for i = 0; i < 5; i++ {
	// 	for j = 0; j < 5; j++ {
	// 		fmt.Printf("[%d][%d] = %d ", i, j, carte[i][j])
	// 	}
	// 	fmt.Print("\n")
	// }

	// Initialisation des variables
	joueurs = make(map[int]string)
	nbBateau = make(map[int]int)
	r := mux.NewRouter()
	joueur := r.PathPrefix("/joueur").Subrouter()
	jeux := r.PathPrefix("/jeux").Subrouter()

	// Gestion de la partie joueur
	joueur.HandleFunc("/liste", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if checkJoueur(joueurs[1]) && checkJoueur(joueurs[2]) {
			w = reponseFormatee(w, true, "Liste de tous les joueurs", joueurs)
		} else {
			w = reponseFormatee(w, false, "Il n'y a pas de joueurs.", nil)
		}

	})

	joueur.HandleFunc("/liste/{idJoueur}", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
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
		enableCors(&w)
		vars := mux.Vars(r)
		nomjoueur := vars["nomjoueur"]
		numJoueur := 1

		if joueurs[1] == "" {
			joueurs[1] = nomjoueur
		} else if joueurs[2] == "" {
			joueurs[2] = nomjoueur
			numJoueur = 2
		} else {
			w = reponseFormatee(w, false, "Les 2 joueurs sont déjà renseigner", nil)
			return
		}

		w = reponseFormatee(w, true, fmt.Sprintf("Ajout de %s qui est le joueur %d\n", joueurs[numJoueur], numJoueur), joueurs[numJoueur])
	})

	// Gestion de la partie jeux
	jeux.HandleFunc("/tour", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		//

		if joueurEnCours == 0 {
			joueurEnCours = 1
		}
		w = reponseFormatee(w, true, "Joueur en cours", joueurEnCours)
	})
	// Lieux du tir du joueur dont c'est le tour au format "00" → "44"
	jeux.HandleFunc("/tir/{empl}", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		vars := mux.Vars(r)
		empl := vars["empl"]

		emplacement = strings.Split(empl, "")
		iX, _ := strconv.Atoi(emplacement[0])
		iY, _ := strconv.Atoi(emplacement[1])

		if joueurEnCours != 1 {
			if (carteJ1[iX][iY] == 1) && carteJ1[iX][iY] != 3 {
				carteJ1[iX][iY] = 3
				joueurEnCours = 1
				w = reponseFormatee(w, true, "J1 Touché", carteJ1)
			} else {
				joueurEnCours = 1
				w = reponseFormatee(w, false, "Loupé", nil)
			}
		} else {
			if (carteJ2[iX][iY] == 2) && carteJ2[iX][iY] != 3 {
				carteJ2[iX][iY] = 3
				joueurEnCours = 2
				w = reponseFormatee(w, true, "J2 Touché", carteJ2)
			} else {
				joueurEnCours = 2
				w = reponseFormatee(w, false, "Loupé", nil)
			}
		}

	})
	// Lieux du tir du joueur dont c'est le tour au format "00" → "44"
	jeux.HandleFunc("/carte/{idJoueur}", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		vars := mux.Vars(r)
		idJoueur, _ := strconv.Atoi(vars["idJoueur"])

		if idJoueur == 1 {
			w = reponseFormatee(w, true, "Carte J1", carteJ1)
		} else {
			w = reponseFormatee(w, true, "Carte J2", carteJ2)
		}

	})
	//Fonction pour savoir a qui appartient une case
	// jeux.HandleFunc("/case/{empl}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	empl := vars["empl"]

	// 	emplacement = strings.Split(empl, "")

	// 	iX, _ := strconv.Atoi(emplacement[0])
	// 	iY, _ := strconv.Atoi(emplacement[1])

	// 	if carte[iX][iY] != 3 {
	// 		w = reponseFormatee(w, true, "La case appartient au joueur", joueurs[carte[iX][iY]])
	// 	} else {
	// 		w = reponseFormatee(w, false, "La case n'appartient a aucun joueur", nil)
	// 	}
	// })

	//Fonction pour qu'un joueur place un navire a un emplacement
	jeux.HandleFunc("/placer/{idJoueur}/{horiVet}/{empl}", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		vars := mux.Vars(r)
		idJoueur, _ := strconv.Atoi(vars["idJoueur"])
		horiVet := vars["horiVet"]
		empl := vars["empl"]
		emplacement = strings.Split(empl, "")

		iX, _ := strconv.Atoi(emplacement[0])
		iY, _ := strconv.Atoi(emplacement[1])

		// par default a false
		// var valide bool
		if idJoueur == 1 {
			if horiVet == "H" {
				if iY < 3 {
					carteJ1[iX][iY] = idJoueur
					carteJ1[iX][iY+1] = idJoueur
					carteJ1[iX][iY+2] = idJoueur
					nbBateau[idJoueur] = nbBateau[idJoueur] + 1
					w = reponseFormatee(w, true, "Emplacement OK", carteJ1)
				} else {
					w = reponseFormatee(w, false, "Emplacement Impossible", nil)
				}
			} else if horiVet == "V" {
				if iX < 3 {
					carteJ1[iX][iY] = idJoueur
					carteJ1[iX+1][iY] = idJoueur
					carteJ1[iX+2][iY] = idJoueur
					nbBateau[idJoueur] = nbBateau[idJoueur] + 1
					w = reponseFormatee(w, true, "Emplacement OK", carteJ1)
				} else {
					w = reponseFormatee(w, false, "Emplacement Impossible", nil)
				}
			}
		} else if idJoueur == 2 {
			if horiVet == "H" {
				if iY < 3 {
					carteJ2[iX][iY] = idJoueur
					carteJ2[iX][iY+1] = idJoueur
					carteJ2[iX][iY+2] = idJoueur
					nbBateau[idJoueur] = nbBateau[idJoueur] + 1
					w = reponseFormatee(w, true, "Emplacement OK", carteJ2)
				} else {
					w = reponseFormatee(w, false, "Emplacement Impossible", nil)
				}
			} else if horiVet == "V" {
				if iY < 3 {
					carteJ2[iX][iY] = idJoueur
					carteJ2[iX+1][iY] = idJoueur
					carteJ2[iX+2][iY] = idJoueur
					nbBateau[idJoueur] = nbBateau[idJoueur] + 1
					w = reponseFormatee(w, true, "Emplacement OK", carteJ2)
				} else {
					w = reponseFormatee(w, false, "Emplacement Impossible", nil)
				}
			}
		}

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
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
