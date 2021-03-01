package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/kisinga/ATS/app/models"
)

type Interactor interface {
	GetUser(context.Context, string) (*models.User, error)
	GetUserByID(context.Context, primitive.ObjectID) (*models.User, error)
	GetMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.User, error)
	AddUser(ctx context.Context, user models.NewUser, creatorEmail string) (*models.User, error)
	UpdateUser(ctx context.Context, email string, newUser models.User) (*models.User, error)
	Count(ctx context.Context) (int64, error)
}

type interactor struct {
	repository Repository
}

func NewIterator(repo Repository) Interactor {
	return &interactor{
		repository: repo,
	}
}
func (i *interactor) Count(ctx context.Context) (int64, error) {
	return i.repository.Count(ctx)
}
func (i *interactor) GetUser(ctx context.Context, email string) (*models.User, error) {
	return i.repository.Read(ctx, email)
}

func (i *interactor) GetUserByID(ctx context.Context, ID primitive.ObjectID) (*models.User, error) {
	return i.repository.ReadByID(ctx, ID)
}

func (i *interactor) AddUser(ctx context.Context, user models.NewUser, creatorEmail string) (*models.User, error) {
	creator, err := i.GetUser(ctx, creatorEmail)
	if err != nil {
		return nil, err
	}
	newUser := models.User{
		Email:     user.Email,
		CreatedBy: creator.ID,
		ID:        primitive.NewObjectID(),
		Name:      user.Name,
		Active:    true,
	}
	return i.repository.Create(ctx, newUser)
}

func (i *interactor) GetMany(ctx context.Context, after primitive.ObjectID, limit *int64) ([]*models.User, error) {
	return i.repository.ReadMany(ctx, after, limit)
}

func (i *interactor) UpdateUser(ctx context.Context, email string, newUser models.User) (*models.User, error) {
	return i.repository.Update(ctx, email, newUser)
}
