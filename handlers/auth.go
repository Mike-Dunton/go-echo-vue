// handlers/tasks.go
package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-echo-vue/models"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

// H return arbitrary JSON in our response
type H map[string]interface{}

// DoAuthorize endpoint
func DoAuthorize(ctx context.Context, db *firestore.Client, googleAuthConfig *oauth2.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			panic(err)
		}
		loginAttempt := new(models.LoginAttempt)
		if err = c.Bind(loginAttempt); err != nil {
			c.Logger().Debug(fmt.Printf("Error Binding Request to LoginAttempt Struct %v \n", err))
			return c.JSON(http.StatusBadRequest, nil)
		}
		user, token, err := models.ExchangeCodeForToken(googleAuthConfig, loginAttempt.Code, c)
		if err != nil {
			c.Logger().Debug(fmt.Printf("OauthUser OauthUser %v \n", err))
			return c.JSON(http.StatusBadRequest, nil)
		}
		user.LastLogin = time.Now().Format(time.RFC3339)
		err = models.SaveUser(ctx, db, user)
		if err != nil {
			c.Logger().Debug(fmt.Printf("sqlExecError sqlExecError %v \n", err))
			return c.JSON(http.StatusBadRequest, nil)
		}
		sess.Values["user-id"] = user.Sub
		marshalledToken, _ := json.Marshal(token)
		sess.Values["token"] = marshalledToken
		sess.Save(c.Request(), c.Response())

		//return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/", c.Request().URL.Host))
		c.Response().Header().Set("token-type", token.Type())
		c.Response().Header().Set("access-token", token.AccessToken)
		c.Response().Header().Set("expiry", token.Expiry.Format(time.RFC3339))
		c.Response().Header().Set("uid", user.Email)
		return c.JSON(http.StatusOK, H{
			"data": user,
		})
	}
}

func GetUser(ctx context.Context, db *firestore.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			panic(err)
		}
		userID := sess.Values["user-id"].(string)
		user, err := models.GetUser(ctx, db, userID)
		if err != nil {
			c.Logger().Debug(fmt.Printf("GetUser OauthUser %v \n", err))
			return c.JSON(http.StatusBadRequest, nil)
		}
		c.Logger().Debug("GetUserGoogle Successful")
		return c.JSON(http.StatusOK, H{
			"data": user,
		})
	}
}

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := new(oauth2.Token)
			sessionToken := new(oauth2.Token)
			c.Logger().Debug("Checking Token")
			sess, err := session.Get("session", c)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(sess.Values["token"].([]byte), &sessionToken)
			if err != nil {
				panic(err)
			}
			token.TokenType = c.Request().Header.Get("token-type")
			token.AccessToken = c.Request().Header.Get("access-token")
			token.Expiry, _ = time.Parse(time.RFC3339, c.Request().Header.Get("expiry"))
			if sessionToken.AccessToken != token.AccessToken {
				return echo.ErrUnauthorized
			}
			c.Request().Header.Get("uid")
			if token != nil && token.Valid() {
				c.Logger().Debug("Valid Token")
				c.Set("token", token)
				c.Set("uid", c.Request().Header.Get("uid"))
				return next(c)
			} else {
				c.Logger().Debug("Invalid Token")
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
