package main

import (
	"groupie/handler"
	"log"
	"net/http"
	"time"
)

const (
	PORT             = ":8080"
	URL_API_ARTIST   = "https://groupietrackers.01edu.ai/api/artists"
	URL_API_RELATION = "https://groupietrackers.01edu.ai/api/relation"
	REFRESH_INTERVAL = 30 * time.Second
)

func main() {
	// On crée une variable JsonArtist, une Struct de type Artist (voir artistStruct.go dans le dossier models)
	// pour decoder les données JSON et les intégrer dans une variable go (json Artist)
	// pour éviter de solliciter l'API. C'est un "cache".
	jsonArtists, err := handler.JsonAllArtist(URL_API_ARTIST)
	if err != nil {
		log.Fatal("Erreur dans le chargement des données de l'API")
	}
	jsonDatesLocations, err := handler.JsonAllDatesLocations(URL_API_RELATION)
	if err != nil {
		log.Fatal("Erreur dans le chargement des données de l'API")
	}

	//Configuration d'une goroutine avec le package time pour créer un système de rafraichissement des données de l'API.
	go func() {
		ticker := time.NewTicker(REFRESH_INTERVAL)
		defer ticker.Stop()
		for range ticker.C {
			jsonArtists, err = handler.JsonAllArtist(URL_API_ARTIST)
			if err != nil {
				log.Fatal("Erreur dans le chargement des données de l'API")
			}
			jsonDatesLocations, err = handler.JsonAllDatesLocations(URL_API_RELATION)
			if err != nil {
				log.Fatal("Erreur dans le chargement des données de l'API")
			}
			log.Println("Rafraichissement des données de l'API...")
		}
	}()

	// On crée et on initie les gestionnaires http du projet.
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", handler.IndexHandler(&jsonArtists))
	http.HandleFunc("/search", handler.SearchHandler(&jsonArtists))
	http.HandleFunc("/artist/", handler.ArtistHandler(&jsonArtists, &jsonDatesLocations))

	log.Println("Listening on ", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))

}
