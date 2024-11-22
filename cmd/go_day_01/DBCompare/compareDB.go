package main

import (
	"github.com/Saracomethstein/go_day_01/internal/go_day_01/fileutil"
	"github.com/Saracomethstein/go_day_01/internal/go_day_01/handling"
	"github.com/Saracomethstein/go_day_01/internal/go_day_01/read"
	dbcompare "github.com/Saracomethstein/go_day_01/internal/pkg/DBCompare"
	dbreader "github.com/Saracomethstein/go_day_01/internal/pkg/DBReader"
)

func main() {
	oldFile, newFile, err := fileutil.ParseToCompare()
	if err != nil {
		handling.Error(err)
		return
	}
	var reader dbreader.DBReader
	oldRecipe, err := read.Read(reader, oldFile)
	if err != nil {
		handling.Error(err)
		return
	}

	newRecipe, err := read.Read(reader, newFile)
	if err != nil {
		handling.Error(err)
		return
	}

	dbcompare.CompareDatabases(oldRecipe, newRecipe)
}
