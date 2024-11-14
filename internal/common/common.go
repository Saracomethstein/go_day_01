package file

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetName() (string, error) {
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

func InvertFormat(format string) string {
	if format == "json" {
		return "xml"
	}
	return "json"
}
