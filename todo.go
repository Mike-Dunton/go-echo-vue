package main

import (
	"database/sql"
	"go-echo-vue/handlers"
	"os"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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

	cfg := mysql.Cfg(os.Getenv("TODO_SQL_CONNECTION_STRING"), os.Getenv("TODO_SQL_USER"), os.Getenv("TODO_SQL_PASSWORD"))
	cfg.DBName = os.Getenv("TODO_SQL_DATABASE")
	db, err := mysql.DialCfg(cfg)
	if err != nil {
		panic(err)
	}

	// create tasks db
	migrate(db)

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
	e.POST("/login/google", handlers.DoAuthorize(db, googleOauthConfig))
	e.Static("/", "client/recipe-book/dist/")
	// e.GET("/login", handlers.DoLogin(googleOauthConfig))

	private := e.Group("/auth")
	private.Use(handlers.AuthMiddleware())
	private.GET("/user", handlers.GetUser(googleOauthConfig))

	private.GET("/tasks", handlers.GetTasks(db))
	private.PUT("/tasks", handlers.PutTask(db))
	private.DELETE("/tasks/:id", handlers.DeleteTask(db))

	private.GET("/recipes", handlers.GetRecipes(db))
	private.GET("/recipes/:id", handlers.GetRecipe(db))

	// Start as a web server
	e.Start(":8000")
}

func migrate(db *sql.DB) {
	usersSql := `
	CREATE TABLE IF NOT EXISTS users(
        email VARCHAR(64) PRIMARY KEY NOT NULL
	);
	`
	tasksSql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
		user_id  VARCHAR(64) NOT NULL,
		description VARCHAR(64) NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(email) 
	);
    `

	_, err := db.Exec(usersSql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(tasksSql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
