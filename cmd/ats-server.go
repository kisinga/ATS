package main

import (
	"context"
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
	db, firebase, err := storage.New(context.Background())
	if err != nil {
		log.Fatalln("Error", err)
	}
	err = app.NewApp(db, firebase, getPort(), prod)
	if err != nil {
		log.Fatalln("Error", err)
	}
}

// Get the Port from the environment so we can run on Heroku
func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4242"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
