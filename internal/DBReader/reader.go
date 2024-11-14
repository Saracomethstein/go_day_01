package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type RecipeFile struct {
	Cakes []Recipe `json:"cake" xml:"cake"`
}

type Recipe struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
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

	var recipeFile RecipeFile
	err = xml.NewDecoder(file).Decode(&recipeFile)

	return recipeFile.Cakes, err
}

func PrintRecipes(recipes []Recipe, format string) error {
	var output []byte
	var err error

	recipeFile := RecipeFile{Cakes: recipes}
	if format == "json" {
		output, err = json.MarshalIndent(recipeFile, "", "    ")
	} else if format == "xml" {
		output, err = xml.MarshalIndent(recipeFile, "", "    ")
	}

	if err != nil {
		return err
	}

	fmt.Println(string(output))
	return nil
}
