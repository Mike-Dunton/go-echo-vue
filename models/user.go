// Package models db interface
package models

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/oauth2"
)

// User is a retrieved and authenticated user.
type User struct {
	Sub           string `json:"sub" mapstructure:"sub"`
	Name          string `json:"name" mapstructure:"name"`
	GivenName     string `json:"given_name" mapstructure:"given_name"`
	FamilyName    string `json:"family_name" mapstructure:"family_name"`
	Profile       string `json:"profile" mapstructure:"profile,omitempty"`
	Picture       string `json:"picture" mapstructure:"picture"`
	Email         string `json:"email" mapstructure:"email"`
	EmailVerified string `json:"email_verified" mapstructure:"email_verified,omitempty"`
	Gender        string `json:"gender" mapstructure:"gender,omitempty"`
	LastLogin     string `json:"lastLogin" mapstructure:"lastLogin"`
}

type LoginAttempt struct {
	Code string `json:"code"`
}

// OauthUser gets info from the oauth endpoint
func ExchangeCodeForToken(googleAuthConfig *oauth2.Config, code string, c echo.Context) (user User, token *oauth2.Token, err error) {
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

func GetUser(ctx context.Context, db *firestore.Client, UserID string) (user User, err error) {
	userDoc, err := db.Collection("users").Doc(UserID).Get(ctx)
	if err != nil {
		return user, err
	}
	if err = userDoc.DataTo(&user); err != nil {
		return user, fmt.Errorf("doc.DataTo: %v", err)
	}

	return user, nil
}

//SaveUser Saves user to DB
func SaveUser(ctx context.Context, db *firestore.Client, user User) error {
	var userDoc map[string]interface{}
	err := mapstructure.Decode(user, &userDoc)
	if err != nil {
		return err
	}
	_, err = db.Collection("users").Doc(user.Sub).Set(ctx, userDoc, firestore.MergeAll)
	if err != nil {
		return err
	}

	return nil
}
