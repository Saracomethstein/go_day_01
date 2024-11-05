package main

import (
	"fmt"

	dbreader "github.com/Saracomethstein/go_day_01/internal/DBReader"
)

func main() {
	var reader dbreader.DBReader
	filename := "recipes.json" // Укажите путь к файлу
	format := "json"           // Выберите формат вывода

	// Выберите JSON или XML reader
	if format == "json" {
		reader = dbreader.JSONReader{}
	} else {
		reader = dbreader.XMLReader{}
	}

	recipes, err := reader.Read(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if err := dbreader.PrintRecipes(recipes, format); err != nil {
		fmt.Println("Error printing recipes:", err)
	}
}
