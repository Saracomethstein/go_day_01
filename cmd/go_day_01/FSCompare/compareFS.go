package main

import (
	"bufio"
	"fmt"
	"os"
)

// test version //
func main() {
	oldFile := "snapshot1.txt"
	newFile := "snapshot2.txt"

	oldSet, err := loadFileToSet(oldFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %s\n", oldFile, err)
		return
	}

	newSet, err := loadFileToSet(newFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %s\n", newFile, err)
		return
	}

	for file := range newSet {
		if !oldSet[file] {
			fmt.Printf("ADDED %s\n", file)
		}
	}

	// Find removed files
	for file := range oldSet {
		if !newSet[file] {
			fmt.Printf("REMOVED %s\n", file)
		}
	}
}

func loadFileToSet(filename string) (map[string]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	set := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		set[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return set, nil
}
