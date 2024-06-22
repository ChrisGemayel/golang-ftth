package main

import (
	"golang-ftth/handler"
	"log"
	"net/http"
)

/*

J'ai quelques regrets :

- Je voulais que le résultat soit trié, mais cela prenait trop de temps et d'efforts.

- J'ai remarqué qu'il y avait des virgules dans la troisième ligne du CSV, mais c'était un peu tard, donc je regrette de ne pas l'avoir remarqué plus tôt.

- J'aurais aussi aimé paralléliser l'ensemble du projet pour le rendre aussi rapide que possible, mais je ne voulais pas dépasser les 2 heures conseillées.

- Dans la dernière version avec les workers, la sortie est également placée dans le répertoire racine, cela pourrait être corrigé mais prend un peu de temps
et je ne veux pas envoyer le Repo trop tard. Je préfère être un homme de parole en le faisant samedi en 2 heures plutôt que de prendre plus de temps et ou de prolonger jusqu'à dimanche.

- Pour le cas d'erreure j'ai preferer juste laisser comme tel car je trouve dommage de retourner une erreure, ceci montre mon application n'est pas assez robuste
pour pouvoir traiter un nombre egal de nomfichier.go,message sur la meme heure et donc la mon code il retourne la premiere qu'il trouve en mode random.

- J'aurais pu ajouter des tests unitaires mais ca ne respecterait plus les 2 heures

- J'aurais aussi pu ajouter que la personne pouvait envoyer le fishier CSV par cette API ou une autre et dans le premier cas ou ce faire directement analyser
ou dans le deuxieme cas, qu'elle soit deposer dans ./data et une autre API ferait l'analyse.



*/

func main() {
	// Définition du gestionnaire pour l'endpoint "/analysis"
	http.HandleFunc("/analysis", handler.AnalysisHandler)

	// Démarrage du serveur HTTP sur le port 15442
	log.Println("Serveur en écoute sur le port 15442")
	log.Fatal(http.ListenAndServe(":15442", nil))
}
