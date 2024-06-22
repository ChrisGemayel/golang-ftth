package handler

import (
	"golang-ftth/analyzer"
	"net/http"
	"os"
	"path/filepath"
)

func AnalysisHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") != "text/csv" {
		http.Error(w, "En-tête Accept invalide", http.StatusNotAcceptable)
		return
	}

	result, err := analyzer.AnalyzeLogs("./data")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Créer le répertoire de sortie s'il n'existe pas
	outputDir := "./output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		http.Error(w, "Échec de la création du répertoire de sortie", http.StatusInternalServerError)
		return
	}

	// Spécifier le chemin du fichier de sortie
	outputPath := filepath.Join(outputDir, "analysis_result.csv")

	// Écrire le résultat dans le fichier de sortie spécifié
	file, err := os.Create(outputPath)
	if err != nil {
		http.Error(w, "Échec de la création du fichier de sortie", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = file.WriteString(result)
	if err != nil {
		http.Error(w, "Échec de l'écriture du résultat dans le fichier", http.StatusInternalServerError)
		return
	}

	// Définir les en-têtes de réponse
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=analysis_result.csv")
	w.Write([]byte(result))
}
