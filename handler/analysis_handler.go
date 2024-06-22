package handler

import (
	"golang-ftth/analyzer"
	"net/http"
	"os"
	"path/filepath"
)

func AnalysisHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") != "text/csv" {
		http.Error(w, "Invalid Accept header", http.StatusNotAcceptable)
		return
	}

	result, err := analyzer.AnalyzeLogs("./data")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create the output directory if it doesn't exist
	outputDir := "./output"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		http.Error(w, "Failed to create output directory", http.StatusInternalServerError)
		return
	}

	// Specify the output file path
	outputPath := filepath.Join(outputDir, "analysis_result.csv")

	// Write result to the specified output file
	file, err := os.Create(outputPath)
	if err != nil {
		http.Error(w, "Failed to create output file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = file.WriteString(result)
	if err != nil {
		http.Error(w, "Failed to write result to file", http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=analysis_result.csv")

	// Serve the file content as response
	http.ServeFile(w, r, outputPath)
}
