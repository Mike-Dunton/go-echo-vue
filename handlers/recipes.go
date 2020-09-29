// handlers/recipes.go
package handlers

import (
	"database/sql"
	"go-echo-vue/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// GetRecipes endpoint
func GetRecipes(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("uid").(string)
		return c.JSON(http.StatusOK, models.GetRecipes(db, userID))
	}
}

func GetRecipe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		recipeID, _ := strconv.Atoi(c.Param("id"))
		userID := c.Get("uid").(string)
		recipe, err := models.GetRecipe(db, userID, recipeID)
		if err != nil {
			return c.JSON(http.StatusNotFound, H{})
		}
		return c.JSON(http.StatusOK, recipe)
	}
}
