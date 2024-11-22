package main

import (
	"fmt"
	"github.com/Saracomethstein/go_day_01/internal/go_day_01/fileutil"
	dbcompare "github.com/Saracomethstein/go_day_01/internal/pkg/DBCompare"
	dbreader "github.com/Saracomethstein/go_day_01/internal/pkg/DBReader"
	"os"
)

func main() {
	oldFile, newFile, err := fileutil.ParseToCompare()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	var reader dbreader.DBReader

	reader, err = dbreader.Invert(reader, oldFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	oldRecipe, err := reader.Read(oldFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	reader, err = dbreader.Invert(reader, newFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	newRecipe, err := reader.Read(newFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	dbcompare.CompareDatabases(oldRecipe, newRecipe)
}
