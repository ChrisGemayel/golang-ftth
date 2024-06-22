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

// CountEntry represents the count of a specific (filename, message) pair
type CountEntry struct {
	Filename string
	Message  string
	Count    int
}

// AnalyzeLogs analyzes the logs from CSV files in the given directory and returns the results as a CSV format string
func AnalyzeLogs(dirPath string) (string, error) {
	var allEntries []utils.LogEntry

	// Parallel processing of CSV files
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".csv") {
			entries, err := utils.ReadCSVFile(path)
			if err != nil {
				return fmt.Errorf("failed to read CSV file %s: %v", path, err)
			}
			allEntries = append(allEntries, entries...)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to read CSV files: %v", err)
	}

	// Use a map with mutex for concurrent access
	counts := make(map[string]map[string]int)
	var mu sync.Mutex // Mutex for synchronization

	// Process entries concurrently
	for _, entry := range allEntries {
		dayHour := entry.Timestamp.Format("0201200615")
		mu.Lock()
		if _, ok := counts[dayHour]; !ok {
			counts[dayHour] = make(map[string]int)
		}
		counts[dayHour][entry.Filename+"|"+entry.Message]++
		mu.Unlock()
	}

	// Sort keys for predictable output order
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
