package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kisinga/ATS/app"
	"github.com/kisinga/ATS/app/storage"
)

var prod bool = false

func main() {
	if os.Getenv("prod") == "true" {
		fmt.Println("We are in production!! Yeah")
		prod = true
	}
	db := storage.New()
	// CORS is enabled only in prod profile
	app := app.NewApp(db, prod)
	err := app.Serve(getPort())
	log.Println("Error", err)
}

// Get the Port from the environment so we can run on Heroku
func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
