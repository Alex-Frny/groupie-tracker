package handler

import (
	"groupie/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// Gestionnaire http pour gérer l'accueil du site internet
// IndexHandler retourne un http.HandlerFunc qui utilise les données préchargées pour éviter les requêtes superflues à l'API.
func IndexHandler(artists *[]models.Artist) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			http.Error(w, "404 - Page introuvable", http.StatusNotFound)
			return
		}

		if r.Method != http.MethodGet {
			http.Error(w, "405 - Méthode non autorisée", http.StatusMethodNotAllowed)
			return
		}

		t, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, artists)
		if err != nil {
			http.Error(w, "Erreur d'affichage", http.StatusInternalServerError)
			return
		}
	}
}

// Gestionnaire http qui gère la recherche de l'utilisateur et affiche les résultats sur la page.
func SearchHandler(artists *[]models.Artist) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			t, err := template.ParseFiles("./templates/search.html")
			if err != nil {
				http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
				return
			}

			err = r.ParseForm()
			if err != nil {
				http.Error(w, "400 Error", http.StatusBadRequest)
			}

			// On récupère la recherche utilisateur pour la comparer aux données contenues dans la liste
			// préchargée. Pour chaque correspondance (que ce soit une sous-chaîne de caractère ou un nom),
			// la fonction retourne une slice Artist qui va constituer la liste des resultats.
			urlSearch := r.URL.RawQuery
			search := strings.TrimPrefix(urlSearch, "name=")
			data := jsonSearch(search, *artists)

			//Si aucune correspondance n'est trouvée, on charge "noArtist.html" pour informer l'utilisateur.
			if len(data) == 0 {
				t := template.Must(template.ParseFiles("./templates/noArtist.html"))

				err = t.Execute(w, nil)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else {
				err = t.Execute(w, data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		} else {
			http.Error(w, "Erreur 405: méthode non autorisé", http.StatusMethodNotAllowed)
		}
	}
}

// Gestionnaire http qui gère les pages d'information sur l'artiste ou groupe de musique choisi par l'utilisateur.
func ArtistHandler(artists *[]models.Artist, datesLocations *[]models.Relation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./templates/artist.html")
		if err != nil {
			http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
			return
		}

		err = r.ParseForm()
		if err != nil {
			http.Error(w, "400 Error", http.StatusBadRequest)
		}

		var artistSearch models.Artist
		urlSearch := r.URL.RawQuery
		searchID := strings.TrimPrefix(urlSearch, "id=")

		for i := 0; i < len(*artists); i++ {
			if strconv.Itoa((*artists)[i].ID) == searchID {
				artistSearch = (*artists)[i]
			}
		}

		//Dans le cas où l'utilisateur chercherait à trouver un artiste dont l'ID
		//n'est pas connu des données json, on envoie la page "noArtist.html"
		if artistSearch.ID <= 0 || artistSearch.ID > len(*artists) {
			t, err = template.ParseFiles("./templates/noArtist.html")
			if err != nil {
				http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
				return
			}

			err = t.Execute(w, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {

			//Implanter une fonction qui va compléter une variable et faire le lien avec
			var artistDetails models.ArtistDetails

			artistDetails.ID = artistSearch.ID
			artistDetails.Artist = artistSearch
			artistDetails.Relation = (*datesLocations)[artistSearch.ID-1]

			err = t.Execute(w, artistDetails)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
