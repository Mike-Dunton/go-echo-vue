// handlers/recipes.go
package handlers

import (
	"context"
	"go-echo-vue/models"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo"
)

// GetRecipes endpoint
func GetRecipes(ctx context.Context, db *firestore.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("uid").(string)
		return c.JSON(http.StatusOK, models.GetRecipes(ctx, db, userID))
	}
}

func GetRecipe(ctx context.Context, db *firestore.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		recipeID, _ := strconv.Atoi(c.Param("id"))
		userID := c.Get("uid").(string)
		recipe, err := models.GetRecipe(ctx, db, userID, recipeID)
		if err != nil {
			return c.JSON(http.StatusNotFound, H{})
		}
		return c.JSON(http.StatusOK, recipe)
	}
}
