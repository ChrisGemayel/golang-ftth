package analyzer

import (
	"fmt"
	"golang-ftth/utils"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

// CountEntry représente le compte d'une paire spécifique (nom de fichier, message)
type CountEntry struct {
	Filename string
	Message  string
	Count    int
}

// fonction worker pour lire les fichiers CSV
func readCSVWorker(files <-chan string, entriesChan chan<- []utils.LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	for path := range files {
		entries, err := utils.ReadCSVFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "échec de la lecture du fichier CSV %s : %v\n", path, err)
			continue
		}
		entriesChan <- entries
	}
}

// fonction agrégateur pour compter les entrées de journal
func aggregateEntries(entriesChan <-chan []utils.LogEntry, counts map[string]map[string]int, mu *sync.Mutex, done chan<- struct{}) {
	defer close(done)
	for entries := range entriesChan {
		mu.Lock()
		for _, entry := range entries {
			dayHour := entry.Timestamp.Format("0201200615")
			if _, ok := counts[dayHour]; !ok {
				counts[dayHour] = make(map[string]int)
			}
			counts[dayHour][entry.Filename+"|"+entry.Message]++
		}
		mu.Unlock()
	}
}

// AnalyzeLogs analyse les journaux à partir des fichiers CSV dans le répertoire donné et renvoie les résultats sous forme de chaîne au format CSV
func AnalyzeLogs(dirPath string) (string, error) {
	var wg sync.WaitGroup
	files := make(chan string, 100)
	entriesChan := make(chan []utils.LogEntry, 100)
	counts := make(map[string]map[string]int)
	var mu sync.Mutex
	done := make(chan struct{})

	// Démarrer les workers pour lire les fichiers CSV
	numFileReaders := 4
	for i := 0; i < numFileReaders; i++ {
		wg.Add(1)
		go readCSVWorker(files, entriesChan, &wg)
	}

	// Démarrer une seule goroutine pour agréger les entrées de journal
	go aggregateEntries(entriesChan, counts, &mu, done)

	// Parcourir le répertoire et envoyer les chemins de fichiers au canal files
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".csv") {
			files <- path
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("échec de la lecture des fichiers CSV : %v", err)
	}
	close(files)

	// Attendre que tous les readCSVWorkers soient terminés
	wg.Wait()
	close(entriesChan)

	// Attendre que l'agrégation soit terminée
	<-done

	// Trier les clés pour un ordre de sortie prévisible
	var sortedKeys []string
	for dayHour := range counts {
		sortedKeys = append(sortedKeys, dayHour)
	}
	sort.Strings(sortedKeys)

	var results []string
	for _, dayHour := range sortedKeys {
		countMap := counts[dayHour]
		maxEntry := CountEntry{}
		for key, count := range countMap {
			parts := strings.Split(key, "|")
			filename, message := parts[0], parts[1]
			if count > maxEntry.Count {
				maxEntry = CountEntry{Filename: filename, Message: message, Count: count}
			}
		}
		day := dayHour[:8]
		hour := dayHour[8:]
		results = append(results, fmt.Sprintf("%s,%s,%s,%s", day, hour, maxEntry.Filename, maxEntry.Message))
	}

	return strings.Join(results, "\n"), nil
}
