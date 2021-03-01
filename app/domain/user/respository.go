package user

import (
	"context"

	"github.com/kisinga/ATS/app/models"
	"github.com/kisinga/ATS/app/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Create(context.Context, models.User) (*models.User, error)
	Read(context.Context, string) (*models.User, error)
	ReadByID(context.Context, primitive.ObjectID) (*models.User, error)
	ReadMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.User, error)
	Update(ctx context.Context, email string, newUser models.User) (*models.User, error)
	Count(ctx context.Context) (int64, error)
}

// CrudChannels keeps a list of channels to which values are emitted when certain CRUD operations are performed
type CrudChannels struct {
	Created chan models.User
	Read    chan models.User
	Updated chan models.User
	Deleted chan models.User
}

// NewCrudChannels creates an instance of CrudChannels
func NewCrudChannels() *CrudChannels {
	return &CrudChannels{
		Created: make(chan models.User),
		Read:    make(chan models.User),
		Updated: make(chan models.User),
		Deleted: make(chan models.User),
	}
}
func NewRepository(database *storage.Database, topics *Topics) Repository {
	return &repository{database, topics}
}

type repository struct {
	db     *storage.Database
	topics *Topics
}

func (r repository) Create(ctx context.Context, user models.User) (*models.User, error) {
	res, err := r.db.Client.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	r.topics.Emit(r.topics.Created, user)
	return &user, nil
}

func (r repository) Read(ctx context.Context, email string) (*models.User, error) {
	user := models.User{}
	err := r.db.Client.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	r.topics.Emit(r.topics.Read, user)
	return &user, nil
}

func (r repository) Count(ctx context.Context) (int64, error) {
	return r.db.Client.Collection("users").CountDocuments(ctx, bson.M{})
}

func (r repository) ReadByID(ctx context.Context, ID primitive.ObjectID) (*models.User, error) {
	user := models.User{}
	err := r.db.Client.Collection("users").FindOne(ctx, bson.M{"_id": ID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	r.topics.Emit(r.topics.Read, user)
	return &user, nil
}

func (r repository) ReadMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.User, error) {
	users := []*models.User{}
	DataCursor, dataErr := r.db.Client.Collection("users").Find(ctx,
		bson.M{"_id": bson.M{"$lt": after}},
		&options.FindOptions{Limit: limit, Sort: bson.M{"_id": -1}})

	if dataErr != nil {
		return nil, dataErr
	}
	for DataCursor.Next(ctx) {
		elem := models.User{}
		err := DataCursor.Decode(&elem)
		if err != nil {
			return nil, err
		}
		users = append(users, &elem)
	}
	for _, user := range users {
		r.topics.Emit(r.topics.Read, *user)
	}
	return users, nil
}
func (r repository) Update(ctx context.Context, email string, newUser models.User) (*models.User, error) {
	user := models.User{}
	after := options.After
	updateOptions := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	err := r.db.Client.Collection("users").FindOneAndUpdate(ctx, bson.M{"email": email}, bson.M{"$set": newUser},
		&updateOptions).Decode(&user)
	if err != nil {
		return nil, err
	}
	r.topics.Emit(r.topics.Updated, user)
	return &user, nil
}
