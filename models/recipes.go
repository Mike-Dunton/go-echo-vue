// Package models db interface
package models

import (
	"database/sql"
	"errors"
)

// Recipe is a struct containing Task data
type Recipe struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Category      string   `json:"category"`
	Description   string   `json:"description"`
	FeaturedImage string   `json:"featuredImage"`
	Images        []string `json:"images"`
}

// GetRecipes returns a list a RecipeCollection
func GetRecipes(db *sql.DB, userID string) []Recipe {
	return localRecipes()
}

func GetRecipe(db *sql.DB, userID string, recipeID int) (Recipe, error) {
	selectedRecipe, found := find(localRecipes(), recipeID)
	if found {
		return selectedRecipe, nil
	}
	return Recipe{}, errors.New("No Recipe Found")
}

func find(a []Recipe, x int) (Recipe, bool) {
	for _, n := range a {
		if x == n.ID {
			return n, true
		}
	}
	return Recipe{}, false
}

func localRecipes() []Recipe {
	return []Recipe{
		{
			ID:            1234,
			Name:          "Meatballs",
			Category:      "Food",
			Description:   "Meaty Meatballs",
			FeaturedImage: "https://baconmockup.com/200/200",
			Images: []string{
				"https://baconmockup.com/200/200",
				"https://baconmockup.com/200/200",
				"https://baconmockup.com/200/200",
			},
		},
		{
			ID:            69420,
			Name:          "Japanese Old Fashioned ",
			Category:      "Drink",
			Description:   "Old Fashioned Made with Japanese Whisky",
			FeaturedImage: "https://placebeer.com/400/400",
			Images: []string{
				"https://placebeer.com/400/400",
				"https://placebeer.com/200/200",
				"https://placebeer.com/200/200",
			},
		},
	}
}
