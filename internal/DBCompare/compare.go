package dbcompare

import (
	"fmt"
	dbreader "github.com/Saracomethstein/go_day_01/internal/DBReader"
)

func CompareDatabases(oldRecipes, newRecipes []dbreader.Recipe) {
	oldMap := makeRecipeMap(oldRecipes)
	newMap := makeRecipeMap(newRecipes)

	compareCakes(oldMap, newMap)
	for name, oldRecipe := range oldMap {
		if newRecipe, exists := newMap[name]; exists {
			compareRecipeDetails(oldRecipe, newRecipe)
		}
	}
}

func makeRecipeMap(recipes []dbreader.Recipe) map[string]dbreader.Recipe {
	recipeMap := make(map[string]dbreader.Recipe)
	for _, recipe := range recipes {
		recipeMap[recipe.Name] = recipe
	}
	return recipeMap
}

func compareCakes(oldMap, newMap map[string]dbreader.Recipe) {
	for name := range oldMap {
		if _, exists := newMap[name]; !exists {
			fmt.Printf("REMOVED cake \"%s\"\n", name)
		}
	}
	for name := range newMap {
		if _, exists := oldMap[name]; !exists {
			fmt.Printf("ADDED cake \"%s\"\n", name)
		}
	}
}

func compareRecipeDetails(oldRecipe, newRecipe dbreader.Recipe) {
	if oldRecipe.Time != newRecipe.Time {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldRecipe.Name, newRecipe.Time, oldRecipe.Time)
	}

	oldIngredients := makeIngredientMap(oldRecipe.Ingredients)
	newIngredients := makeIngredientMap(newRecipe.Ingredients)

	for name, oldIng := range oldIngredients {
		if newIng, exists := newIngredients[name]; exists {
			compareIngredientDetails(oldRecipe.Name, oldIng, newIng)
		} else {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", name, oldRecipe.Name)
		}
	}
	for name := range newIngredients {
		if _, exists := oldIngredients[name]; !exists {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", name, newRecipe.Name)
		}
	}
}

func makeIngredientMap(ingredients []dbreader.Ingredient) map[string]dbreader.Ingredient {
	ingredientMap := make(map[string]dbreader.Ingredient)
	for _, ingredient := range ingredients {
		ingredientMap[ingredient.Name] = ingredient
	}
	return ingredientMap
}

func compareIngredientDetails(cakeName string, oldIng, newIng dbreader.Ingredient) {
	if oldIng.Count != newIng.Count {
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIng.Name, cakeName, newIng.Count, oldIng.Count)
	}
	if oldIng.Unit != newIng.Unit {
		if oldIng.Unit == "" {
			fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", newIng.Unit, oldIng.Name, cakeName)
		} else if newIng.Unit == "" {
			fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", oldIng.Unit, oldIng.Name, cakeName)
		} else {
			fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIng.Name, cakeName, newIng.Unit, oldIng.Unit)
		}
	}
}
