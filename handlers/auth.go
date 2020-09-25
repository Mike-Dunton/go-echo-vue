// handlers/tasks.go
package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-echo-vue/models"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

// DoAuthorize endpoint
func DoAuthorize(db *sql.DB, googleAuthConfig *oauth2.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			panic(err)
		}
		if c.QueryParam("state") != sess.Values["StateString"] {
			return c.JSON(http.StatusUnauthorized, nil)
		} else {
			user, token, err := models.OauthUser(googleAuthConfig, c.QueryParam("code"), c)
			if err != nil {
				c.Logger().Debug(fmt.Printf("OauthUser OauthUser %v \n", err))
				return c.JSON(http.StatusBadRequest, nil)
			}
			_, err = models.SaveUser(db, user)
			if err != nil {
				c.Logger().Debug(fmt.Printf("sqlExecError sqlExecError %v \n", err))
				return c.JSON(http.StatusBadRequest, nil)
			}
			sess.Values["user-id"] = user.Email
			marshalledToken, _ := json.Marshal(token)
			sess.Values["token"] = marshalledToken
			sess.Save(c.Request(), c.Response())

			return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/", c.Request().URL.Host))
			//return c.JSON(http.StatusOK, user)
		}
	}
}

// DoLogin endpoint
func DoLogin(googleAuthConfig *oauth2.Config) echo.HandlerFunc {
	stateString, _ := randToken()
	url := googleAuthConfig.AuthCodeURL(stateString)
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			panic(err)
		}
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   3600 * 7,
			HttpOnly: false,
			Secure:   false,
		}
		sess.Values["StateString"] = stateString
		sess.Save(c.Request(), c.Response())
		return c.JSON(http.StatusOK, H{
			"loginUrl": url,
		})
	}
}

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("session", c)
			if err != nil {
				return echo.ErrUnauthorized
			}
			token := new(oauth2.Token)
			c.Logger().Debug("Unmarshal Token")
			if sess.Values["token"] != nil && sess.Values["user-id"] != nil {
				json.Unmarshal(sess.Values["token"].([]byte), &token)
				c.Logger().Debug("Checking Token")
				if token != nil && token.Valid() {
					c.Logger().Debug("Valid Token")
					return next(c)
				} else {
					c.Logger().Debug("Invalid Token")
					return echo.ErrUnauthorized
				}
			} else {
				c.Logger().Debug("No Session.")
				return echo.ErrUnauthorized
			}
		}
	}
}

func randToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(b), nil
}
