package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// LogEntry represents a log entry
type LogEntry struct {
	Timestamp time.Time
	Filename  string
	Message   string
}

// readCSVFile reads a CSV file and returns its entries
func ReadCSVFile(filePath string) ([]LogEntry, error) {
	var entries []LogEntry

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %v", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Custom logic to split line into fields
		fields := strings.Split(line, ",")

		// Ensure there are at least 3 fields (timestamp, filename, message)
		if len(fields) < 3 {
			continue // Skip lines with fewer than 3 fields
		}

		// Join excess fields back into the message field
		message := strings.Join(fields[2:], ",")

		timestamp, err := time.Parse(time.RFC3339, fields[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse timestamp in CSV file %s: %v", filePath, err)
		}

		entry := LogEntry{
			Timestamp: timestamp,
			Filename:  fields[1],
			Message:   message,
		}
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading CSV file %s: %v", filePath, err)
	}

	return entries, nil
}
