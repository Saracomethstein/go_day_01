package fscompare

import (
	"bufio"
	"fmt"
	"os"
)

func LoadFile(filename string) (map[string]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

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

func Compare(oldSet, newSet map[string]bool) {
	for file := range newSet {
		if !oldSet[file] {
			fmt.Printf("ADDED %s\n", file)
		}
	}

	for file := range oldSet {
		if !newSet[file] {
			fmt.Printf("REMOVED %s\n", file)
		}
	}
}
