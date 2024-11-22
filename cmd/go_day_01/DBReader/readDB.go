package main

import (
	"github.com/Saracomethstein/go_day_01/internal/go_day_01/fileutil"
	"github.com/Saracomethstein/go_day_01/internal/go_day_01/handling"
	"github.com/Saracomethstein/go_day_01/internal/go_day_01/read"
	dbreader "github.com/Saracomethstein/go_day_01/internal/pkg/DBReader"
)

func main() {
	filename, err := fileutil.GetInfo()
	if err != nil {
		handling.Error(err)
		return
	}

	var reader dbreader.DBReader
	recipes, err := read.Read(reader, filename)
	if err != nil {
		handling.Error(err)
		return
	}

	if err := dbreader.PrintRecipes(recipes, filename); err != nil {
		handling.Error(err)
		return
	}
}
