package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kisinga/ATS/app"
	"github.com/kisinga/ATS/app/communication/sms"
	"github.com/kisinga/ATS/app/storage"
)

var prod bool = false
var live bool = false

func main() {
	if os.Getenv("prod") == "true" {
		fmt.Println("We are in production!! Yeah")
		prod = true
	}
	if os.Getenv("live") == "true" {
		fmt.Println("Not local server")
		live = true
	}
	// prod = true
	// I know that this should be in a config somewhere, but I'll put it here for now
	dburi := "mongodb+srv://backend:0SLbeeQ1Z0gg@cluster0.zq04m.mongodb.net/prod?retryWrites=true&w=majority"

	db, firebase, err := storage.New(context.Background(), dburi, prod, live)
	if err != nil {
		log.Fatalln("Error", err)
	}

	sms := sms.NewSMS(sms.AfricasTalking{
		URI:      "https://api.africastalking.com/version1/messaging",
		Username: "atske",
		Key:      "40e4a22c0a93284f930443076b5dba7cf7f287098ee90baa4fb101c582de3994",
	})
	err = app.Serve(db, sms, firebase, getPort(), prod)
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
