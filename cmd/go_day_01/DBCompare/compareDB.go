package main

import (
	"fmt"
	dbcompare "github.com/Saracomethstein/go_day_01/internal/DBCompare"
	dbreader "github.com/Saracomethstein/go_day_01/internal/DBReader"
	"github.com/Saracomethstein/go_day_01/internal/file"
	"os"
)

func main() {
	oldFile, newFile, err := file.ParseToCompare()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	var reader dbreader.DBReader

	reader, err = dbreader.Invert(reader, oldFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	oldRecipe, err := reader.Read(oldFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	reader, err = dbreader.Invert(reader, newFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	newRecipe, err := reader.Read(newFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	dbcompare.CompareDatabases(oldRecipe, newRecipe)
}
