package main

import (
	"fmt"
	"github.com/Saracomethstein/go_day_01/internal/go_day_01/fileutil"
	fscompare "github.com/Saracomethstein/go_day_01/internal/pkg/FSCompare"
	"os"
)

func main() {
	oldFile, newFile, err := fileutil.ParseToCompare()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}

	oldSet, err := fscompare.LoadFile(oldFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}

	newSet, err := fscompare.LoadFile(newFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}

	fscompare.Compare(oldSet, newSet)
}
