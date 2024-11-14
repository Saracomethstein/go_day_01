package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type RecipeFile struct {
	Cakes []Recipe `json:"cake"`
}

type Recipe struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"time"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>ingredient"`
}

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"ingredient_name"`
	Count string `json:"ingredient_count" xml:"ingredient_count"`
	Unit  string `json:"ingredient_unit,omitempty" xml:"ingredient_unit,omitempty"`
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

	var recipeFile RecipeFile
	err = json.NewDecoder(file).Decode(&recipeFile)

	if err != nil {
		return nil, err
	}

	return recipeFile.Cakes, nil
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
