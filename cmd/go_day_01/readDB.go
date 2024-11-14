package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	dbreader "github.com/Saracomethstein/go_day_01/internal/DBReader"
)

func main() {
	var reader dbreader.DBReader

	filename, err := getFile()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	format, err := getFileFormat(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if format == "json" {
		reader = dbreader.JSONReader{}
	} else {
		reader = dbreader.XMLReader{}
	}

	recipes, err := reader.Read(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if err := dbreader.PrintRecipes(recipes, format); err != nil {
		fmt.Println("Error printing recipes:", err)
	}
}

func getFile() (string, error) {
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

func getFileFormat(filename string) (string, error) {
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
