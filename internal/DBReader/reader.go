package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Recipe struct {
	Name        string   `json:"name" xml:"name"`
	Ingredients []string `json:"ingredients" xml:"ingredients>ingredient"`
	Time        int      `json:"stovetime" xml:"time"`
}

type DBReader interface {
	Read(filename string) ([]Recipe, error)
}

type JSONReader struct{}

type XMLReader struct{}

func (r JSONReader) Read(filename string) ([]Recipe, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var recipes []Recipe
	err = json.NewDecoder(file).Decode(&recipes)

	return recipes, err
}

func (r XMLReader) Read(filename string) ([]Recipe, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var recipes []Recipe
	err = xml.NewDecoder(file).Decode(&recipes)

	return recipes, err
}

func PrintRecipes(recipes []Recipe, format string) error {
	var data []byte
	var err error
	if format == "json" {
		data, err = json.MarshalIndent(recipes, "", "    ")
	} else if format == "xml" {
		data, err = xml.MarshalIndent(recipes, "", "    ")
	} else {
		return fmt.Errorf("unknown format: %s", format)
	}
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
