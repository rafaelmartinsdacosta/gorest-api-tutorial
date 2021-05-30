package main

import (
	"fmt"
	"net/http"

	"github.com/rafaelmartinsdacosta/gorest-api-tutorial/internal/comment"
	"github.com/rafaelmartinsdacosta/gorest-api-tutorial/internal/database"
	transportHTTP "github.com/rafaelmartinsdacosta/gorest-api-tutorial/internal/transport/http"
)

// App - the struct wich contains things like pointers
// to database connections
type App struct {
}

// Run - setup our application
func (app *App) Run() error {
	fmt.Println("setting up our app")

	var err error

	db, err := database.NewDatabase()

	if err != nil {
		return err
	}

	err = database.MigrateDB(db)

	if err != nil {
		return err
	}

	commentService := comment.NewSercice(db)

	handler := transportHTTP.Newhandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		return err
	}

	return nil
}

// main - entrypoint
func main() {
	fmt.Println("GO REST API Course")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
