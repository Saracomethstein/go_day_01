package main

import (
	"fmt"
	"os"

	dbreader "github.com/Saracomethstein/go_day_01/internal/DBReader"
	file "github.com/Saracomethstein/go_day_01/internal/file"
)

func main() {
	filename, err := file.GetInfo()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	var reader dbreader.DBReader
	reader, err = dbreader.Invert(reader, filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	recipes, err := reader.Read(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if err := dbreader.PrintRecipes(recipes, filename); err != nil {
		fmt.Println("Error printing recipes:", err)
	}
}
