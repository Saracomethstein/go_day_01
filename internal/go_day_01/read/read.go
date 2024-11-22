package read

import (
	dbreader "github.com/Saracomethstein/go_day_01/internal/pkg/DBReader"
)

func Read(reader dbreader.DBReader, filename string) ([]dbreader.Recipe, error) {
	reader, err := dbreader.Invert(reader, filename)
	if err != nil {
		return nil, err
	}

	recipes, err := reader.Read(filename)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}
