package user

import (
	"context"

	"github.com/kisinga/ATS/app/models"
	"github.com/kisinga/ATS/app/storage"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	Create(context.Context, models.User) (*models.User, error)
	Read(context.Context, string) (*models.User, error)
	ReadMany(context.Context, string, string) ([]models.User, error)
	Update(context.Context, models.User) (*models.User, error)
	Delete(context.Context, string) (*models.User, error)
}

func NewRepository(database *storage.Database) Repository {
	return &repository{database}
}

type repository struct {
	db *storage.Database
}

func (r repository) Create(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}
func (r repository) Read(ctx context.Context, email string) (*models.User, error) {
	user := models.User{}
	return &user, r.db.Client.Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user)
}
func (r repository) ReadMany(ctx context.Context, after string, limit int) ([]models.User, error) {
	return nil, nil
}
func (r repository) Update(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}
func (r repository) Delete(ctx context.Context, email string) (*models.User, error) {
	return nil, nil
}
