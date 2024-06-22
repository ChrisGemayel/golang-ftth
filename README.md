# golang-ftth

## Description

`golang-ftth` est un projet en Go conçu pour analyser des fichiers de journaux au format CSV et générer un rapport agrégé également au format CSV. Ce projet met en évidence l'analyse et l'agrégation de journaux en utilisant des techniques de parallélisation pour améliorer les performances.

## Structure du Projet

golang-ftth/
│
├── Makefile
├── main.go
├── handler/
│ └── analysis_handler.go
├── utils/
│ └── log_utils.go
└── analyzer/
└── log_analyzer.go


## Fonctionnalités

- Lire et analyser les fichiers CSV de journaux
- Agréger les journaux par jour et heure
- Identifier le couple (fichier, message) avec le plus grand nombre d'occurrences pour chaque intervalle temporel
- Fournir une API HTTP pour récupérer les résultats d'analyse au format CSV

## Installation

1. Clonez le dépôt :
    ```sh
    git clone https://github.com/votre-utilisateur/golang-ftth.git
    cd golang-ftth
    ```

2. Assurez-vous d'avoir Go installé (version 1.16 ou supérieure) :
    ```sh
    go version
    ```

3. Placez vos fichiers CSV de journaux dans le dossier `./data`.

## Utilisation

1. Pour exécuter le projet, utilisez `make` :
    ```sh
    make run
    ```

2. Le serveur HTTP écoutera sur le port `15442`.

3. Pour récupérer le rapport d'analyse au format CSV, utilisez la commande `curl` :
    ```sh
    curl -H "Accept: text/csv" http://localhost:15442/analysis -o report.csv
    ```

## Structure des Dossiers

- `main.go` : Point d'entrée principal de l'application.
- `handler/` : Contient le gestionnaire HTTP pour l'analyse des journaux.
  - `analysis_handler.go` : Gère les requêtes HTTP pour analyser les journaux.
- `utils/` : Contient les utilitaires utilisés dans le projet.
  - `log_utils.go` : Fournit des fonctions utilitaires pour manipuler les journaux.
- `analyzer/` : Contient la logique d'analyse des journaux.
  - `log_analyzer.go` : Contient les fonctions pour lire, analyser et agréger les journaux.

## Exemple de Fichier CSV d'Entrée

2000-01-01T00:00:00+00:00,main.go,Listening on http://localhost:8080
2000-01-01T00:00:10+00:00,main.go,Fatal error: restarting
2000-01-01T00:00:11+00:00,main.go,Listening on http://localhost:8080

## Sortie Attendue

Pour les journaux d'entrée ci-dessus, la sortie agrégée sera :

01012000,00,main.go,Listening on http://localhost:8080

## Licence

Ce projet est sous licence MIT. Voir le fichier `LICENSE` pour plus de détails.

## Auteurs

- **El Gemayel** - *Initial work* - [ChrisGemayel](https://github.com/ChrisGemayel)
