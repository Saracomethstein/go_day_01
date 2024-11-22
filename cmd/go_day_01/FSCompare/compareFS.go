package main

import (
	"github.com/Saracomethstein/go_day_01/internal/go_day_01/fileutil"
	"github.com/Saracomethstein/go_day_01/internal/go_day_01/handling"
	fscompare "github.com/Saracomethstein/go_day_01/internal/pkg/FSCompare"
)

func main() {
	oldFile, newFile, err := fileutil.ParseToCompare()
	if err != nil {
		handling.Error(err)
		return
	}

	oldSet, err := fscompare.LoadFile(oldFile)
	if err != nil {
		handling.Error(err)
		return
	}

	newSet, err := fscompare.LoadFile(newFile)
	if err != nil {
		handling.Error(err)
		return
	}

	fscompare.Compare(oldSet, newSet)
}
