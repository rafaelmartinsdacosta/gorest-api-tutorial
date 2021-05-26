package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/rafaelmartinsdacosta/gorest-api-tutorial/internal/transport/http"
)

// App - the struct wich contains things like pointers
// to database connections
type App struct {
}

// Run - setup our application
func (app *App) Run() error {
	fmt.Println("setting up our app")

	handler := transportHTTP.Newhandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8081", handler.Router); err != nil {
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
