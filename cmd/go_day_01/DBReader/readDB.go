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
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	var reader dbreader.DBReader
	reader, err = dbreader.Invert(reader, filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	recipes, err := reader.Read(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
		return
	}

	if err := dbreader.PrintRecipes(recipes, filename); err != nil {
		fmt.Fprintf(os.Stderr, "Error printing recipes: %s\n", err)
		return
	}
}
