package main

import (
	"fmt"
	"github.com/Saracomethstein/go_day_01/internal/go_day_01/fileutil"
	dbreader "github.com/Saracomethstein/go_day_01/internal/pkg/DBReader"
	"os"
)

func main() {
	filename, err := fileutil.GetInfo()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	var reader dbreader.DBReader
	reader, err = dbreader.Invert(reader, filename)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	recipes, err := reader.Read(filename)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
		return
	}

	if err := dbreader.PrintRecipes(recipes, filename); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error printing recipes: %s\n", err)
		return
	}
}
