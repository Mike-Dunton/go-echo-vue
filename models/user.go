// Package models db interface
package models

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

// User is a retrieved and authentiacted user.
type User struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Gender        string `json:"gender"`
}

// OauthUser gets info from the oauth endpoint
func OauthUser(googleAuthConfig *oauth2.Config, code string, c echo.Context) (user User, token *oauth2.Token, err error) {
	token, err = googleAuthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		c.Logger().Debug("Exchange Failed")
		return User{}, nil, err
	}
	client := googleAuthConfig.Client(oauth2.NoContext, token)
	userinfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.Logger().Debug("UserInfoGet Failed")
		return User{}, nil, err
	}
	defer userinfo.Body.Close()
	data, _ := ioutil.ReadAll(userinfo.Body)
	if err = json.Unmarshal(data, &user); err != nil {
		c.Logger().Debug("Unmarshal Successful")
		return user, token, nil
	}
	return User{}, nil, err
}

//SaveUser Saves user to DB
func SaveUser(db *sql.DB, user User) (int64, error) {
	sql := "INSERT IGNORE INTO users(email) VALUES(?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, sqlExecError := stmt.Exec(user.Email)
	if sqlExecError != nil {
		return 0, sqlExecError
	}

	return result.LastInsertId()
}
