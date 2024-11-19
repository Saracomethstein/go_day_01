package fileutil

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetInfo() (string, error) {
	filename := flag.String("f", "", "path to the data base file")
	flag.Parse()

	if *filename == "" {
		return "", fmt.Errorf("please provide a file path with -f flag")
	}

	if _, err := os.Stat(*filename); os.IsNotExist(err) {
		return "", fmt.Errorf("file dose not exist: %s", *filename)
	}

	return *filename, nil
}

func ParseToCompare() (string, string, error) {
	oldFile := flag.String("old", "", "path to the old data base file")
	newFile := flag.String("new", "", "path to the new data base file")
	flag.Parse()

	if *oldFile == "" || *newFile == "" {
		return "", "", fmt.Errorf("please provide a file path with -old or -new")
	}

	if _, err := os.Stat(*oldFile); os.IsNotExist(err) {
		return "", "", fmt.Errorf("old file dose not exist: %s", *oldFile)
	}

	if _, err := os.Stat(*newFile); os.IsNotExist(err) {
		return "", "", fmt.Errorf("new file dose not exist: %s", *newFile)
	}

	return *oldFile, *newFile, nil
}

func GetFormat(filename string) (string, error) {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".json":
		return "json", nil
	case ".xml":
		return "xml", nil
	default:
		return "", fmt.Errorf("unknown file format: %s", ext)
	}
}
