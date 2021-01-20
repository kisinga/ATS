package storage

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
)

type Database struct {
	Client *mongo.Database
}

func New(ctx context.Context, prod bool, live bool) (*Database, *firebase.App, error) {
	firebase, err := newFirebase(live)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Connecting to Db........")
	var env string
	if prod {
		env = "prod"
	} else {
		env = "test"
	}
	// I know that this should be in a config somewhere, but I'll put it here for now
	uri := "mongodb+srv://backend:0SLbeeQ1Z0gg@cluster0.zq04m.mongodb.net/prod?retryWrites=true&w=majority"

	clientOptions := options.Client().ApplyURI(uri)
	clientvar, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}
	err = clientvar.Ping(ctx, nil)
	go func() {
		<-ctx.Done()
		clientvar.Disconnect(ctx)
	}()
	if err != nil {
		fmt.Println("Error connecting to Db........", err)
		return nil, nil, err
	}
	fmt.Println("Connected to MongoDB!")

	return &Database{
		Client: clientvar.Database(env),
	}, firebase, nil
}

func newFirebase(live bool) (*firebase.App, error) {
	var opt option.ClientOption
	if live {
		opt = option.WithCredentialsFile("./firebase_admin.json")
	} else {
		opt = option.WithCredentialsFile("../firebase_admin.json")
	}
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
