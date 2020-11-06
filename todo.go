package main

import (
	"context"
	"go-echo-vue/handlers"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"cloud.google.com/go/firestore"
)

var (
	googleOauthConfig *oauth2.Config
)

// todo.go
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"openid",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	ctx := context.Background()
	dbClient := createClient(ctx)
	defer dbClient.Close()

	// Create a new instance of Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))
	if _, debug := os.LookupEnv("DEBUG"); debug {
		e.Logger.SetLevel(log.DEBUG)
	}

	e.File("/login/google", "client/recipe-book/dist/index.html")
	e.POST("/login/google", handlers.DoAuthorize(ctx, dbClient, googleOauthConfig))
	e.Static("/", "client/recipe-book/dist/")

	private := e.Group("/auth")
	private.Use(handlers.AuthMiddleware())
	private.GET("/user", handlers.GetUser(ctx, dbClient))

	private.GET("/recipes", handlers.GetRecipes(ctx, dbClient))
	private.GET("/recipes/:id", handlers.GetRecipe(ctx, dbClient))

	// Start as a web server
	e.Start(":8000")
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("GOOGLE_PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}
