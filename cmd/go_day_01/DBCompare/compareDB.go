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
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	var reader dbreader.DBReader

	reader, err = dbreader.Invert(reader, oldFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	oldRecipe, err := reader.Read(oldFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	reader, err = dbreader.Invert(reader, newFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	newRecipe, err := reader.Read(newFile)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	dbcompare.CompareDatabases(oldRecipe, newRecipe)
}
