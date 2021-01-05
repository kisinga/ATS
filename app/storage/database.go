package storage

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Database struct {
}

func New() (*Database, *firebase.App, error) {
	firebase, err := newFirebase()
	if err != nil {
		return nil, nil, err
	}
	return &Database{}, firebase, nil
}

func newFirebase() (*firebase.App, error) {
	opt := option.WithCredentialsFile("../../firebase_admin.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
