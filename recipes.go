package main

import (
	"fmt"
	"github.com/mdimec4/allrecipes"
)

func main() {
	recipe, error := allrecipes.GetRecipe("11772") // Package allrecepies exports function
	fmt.Println(recipe.Ingredients)                // Exported names must start with uppercase.
}
