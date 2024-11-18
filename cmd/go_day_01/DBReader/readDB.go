package main

import (
	"fmt"
	"os"

	dbreader "github.com/Saracomethstein/go_day_01/internal/DBReader"
	file "github.com/Saracomethstein/go_day_01/internal/file"
)

func main() {
	filename, format, err := file.GetInfo()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	var reader dbreader.DBReader
	switch format {
	case "json":
		reader = dbreader.JSONReader{}
	case "xml":
		reader = dbreader.XMLReader{}
	default:
		fmt.Printf("Unsupported format: %s\n", format)
		os.Exit(1)
	}

	recipes, err := reader.Read(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	format = file.InvertFormat(format)
	if err := dbreader.PrintRecipes(recipes, format); err != nil {
		fmt.Println("Error printing recipes:", err)
	}
}
