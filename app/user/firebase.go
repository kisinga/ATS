package user

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func NewFirebase() (*firebase.App, error) {
	opt := option.WithCredentialsFile("../../firebase_admin.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
