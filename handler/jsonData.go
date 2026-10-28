package handler

import (
	"encoding/json"
	"groupie/models"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// Fonction qui permet de récupérer les données de l'API (section artist) pour les stocker dans une variable.
func JsonAllArtist(url string) ([]models.Artist, error) {
	//Variable qui permet d'initialiser un nouveau client http pour faire les requêtes.
	var client = http.Client{
		Timeout: time.Second * 10,
	}

	//On crée et on envoie une requête GET à l'url entrée par l'utilisateur.
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, Err := io.ReadAll(resp.Body)
	if Err != nil {
		log.Println("Problème dans la lecture du corps de la réponse")
	}

	var result []models.Artist
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error from the decodage of the api/artists")
		return nil, err
	}
	return result, nil
}

// Fonction qui permet de récupérer les correspondances entre la recherche utilisateur et les données contenues
// dans la liste artists.
func jsonSearch(artistName string, artists []models.Artist) []models.Artist {
	var artistSearch []models.Artist

	if len(artistName) == 0 {
		log.Println("Veuillez mettre le nom correct d'un artiste ou d'un groupe")
		return artistSearch
	}

	artistName = strings.ReplaceAll(artistName, "+", " ")

	for i := 0; i < len(artists); i++ {
		if strings.Contains(strings.ToLower(artists[i].Name), strings.ToLower(artistName)) == true {
			artistSearch = append(artistSearch, artists[i])
		}
	}
	return artistSearch
}

// Fonction qui permet de récupérer les données de l'API (section relation) pour les stocker dans une variable.
func JsonAllDatesLocations(url string) ([]models.Relation, error) {
	//Variable qui permet d'initialiser un nouveau client http pour faire les requêtes.
	var client = http.Client{
		Timeout: time.Second * 10,
	}

	//On crée et on envoie une requête GET à l'url entrée par l'utilisateur.
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, Err := io.ReadAll(resp.Body)
	if Err != nil {
		log.Println("Problème dans la lecture du corps de la réponse")
	}

	var result models.RelationResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error from the decodage of the api/relation")
		return nil, err
	}
	return result.Index, nil
}
