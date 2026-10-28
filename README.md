Projet Groupie-Tracker

Par Pierre BASLEY, Alexis FOURNY et Léo GIFFARD.
-----------------------------------------------------

Groupie-Tracker est un projet de groupe dans le cadre du parcours de formation "développeur full-stack" de la zone01.

Le langage de programmation utilisé est golang pour le backend et l'HTML/CSS pour le frontend.

L'objectif du projet est de créer un site facile à prendre en main qui affiche les informations (nom, membres, dates de concert, etc...) de plusieurs artistes ou groupes de musique.

La source de ces données est une api qui peut être consulté par le lien ci-dessous:

Lien: https://groupietrackers.01edu.ai/

## Installation et téléchargement

Prérequis:
=> Go 1.26.5 ou supérieur (version déclarée dans go.mod)

Téléchargement via Git :

Commande : git clone https://github.com/Alex-Frny/groupie-tracker.git

Pas besoin d'installateur.

Projet pédagogique réalisé en groupe. Les crédits des trois auteurs sont conservés. Aucune licence de réutilisation supplémentaire n’est ajoutée à cette copie.

## Comment utiliser le programme ?

1. Lancer le serveur:
    Allez au dossier racine (/groupie-tracker), taper la commande: go run .
    Note: lors du lancement du serveur, un message apparaîtra "[date et heure du lancement serveur] Listening on :8080"
2. Ouvrir la page web:
    Ouvrez votre navigateur web et entrez dans la barre d'URL "http://localhost:8080/"
3. Usage:
    Il suffit de chercher sur la page d'accueil ou par la barre de recherche l'artiste ou groupe de musique et de cliquer sur le résultat pour consulter toutes les données disponibles sur l'artiste/groupe de musique en question.
