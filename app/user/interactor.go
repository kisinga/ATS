package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kisinga/ATS/app/models"
)

type Interactor interface {
	GetUser(context.Context, string) (*models.User, error)
	GetMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.User, error)
	AddUser(context.Context, models.NewUser) (*models.User, error)
}

type interactor struct {
	repository Repository
}

func NewIterator(repo Repository) Interactor {
	return &interactor{
		repository: repo,
	}
}

func (i *interactor) GetUser(ctx context.Context, email string) (*models.User, error) {
	return i.repository.Read(ctx, email)
}

func (i *interactor) AddUser(ctx context.Context, user models.NewUser) (*models.User, error) {
	newUser := models.User{
		Email: user.Email,
	}
	return i.repository.Create(ctx, newUser)
}

func (i *interactor) GetMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.User, error) {
	return i.repository.ReadMany(ctx, after, limit)
}
