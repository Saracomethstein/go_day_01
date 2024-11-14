package main

import (
	"fmt"
	"os"

	dbreader "github.com/Saracomethstein/go_day_01/internal/DBReader"
	file "github.com/Saracomethstein/go_day_01/internal/common"
)

func main() {
	var reader dbreader.DBReader

	filename, err := file.GetName()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	format, err := file.GetFormat(filename)
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

	format = file.InvertFormat(format)

	if err := dbreader.PrintRecipes(recipes, format); err != nil {
		fmt.Println("Error printing recipes:", err)
	}
}
