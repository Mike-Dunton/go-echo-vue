// handlers/tasks.go
package handlers

import (
	"database/sql"
	"go-echo-vue/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

// H return arbitrary JSON in our response
type H map[string]interface{}

// GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		userID := sess.Values["user-id"].(string)
		if err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, models.GetTasks(db, userID))
	}
}

// PutTask endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task models.Task
		sess, err := session.Get("session", c)
		userID := sess.Values["user-id"].(string)
		// bind user data to our task struct
		c.Bind(&task)
		id, err := models.PutTask(db, task.Description, userID)
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}
	}
}

// DeleteTask endpoint
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		userID := sess.Values["user-id"].(string)
		id, _ := strconv.Atoi(c.Param("id"))
		_, err = models.DeleteTask(db, id, userID)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}

	}
}
