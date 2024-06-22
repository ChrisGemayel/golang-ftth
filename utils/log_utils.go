package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// LogEntry représente une entrée de journal
type LogEntry struct {
	Timestamp time.Time
	Filename  string
	Message   string
}

// readCSVFile lit un fichier CSV et renvoie ses entrées
func ReadCSVFile(filePath string) ([]LogEntry, error) {
	var entries []LogEntry

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("échec de l'ouverture du fichier %s : %v", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Logique personnalisée pour diviser la ligne en champs
		fields := strings.Split(line, ",")

		// S'assurer qu'il y a au moins 3 champs (timestamp, filename, message)
		if len(fields) < 3 {
			continue // Ignorer les lignes avec moins de 3 champs
		}

		// Rejoindre les champs excédentaires dans le champ message
		message := strings.Join(fields[2:], ",")

		timestamp, err := time.Parse(time.RFC3339, fields[0])
		if err != nil {
			return nil, fmt.Errorf("échec de l'analyse du timestamp dans le fichier CSV %s : %v", filePath, err)
		}

		entry := LogEntry{
			Timestamp: timestamp,
			Filename:  fields[1],
			Message:   message,
		}
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("erreur de lecture du fichier CSV %s : %v", filePath, err)
	}

	return entries, nil
}
