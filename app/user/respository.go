package user

import (
	firebase "firebase.google.com/go"
	"github.com/kisinga/ATS/app/models"
	"github.com/kisinga/ATS/app/storage"
)

type Repository interface {
	Create(models.User) (*models.User, error)
	Read(email string) (*models.User, error)
	ReadMany(start string, stop string) ([]models.User, error)
	Update(models.User) (*models.User, error)
	Delete(email string) (*models.User, error)
}

type repository struct {
	db       *storage.Database
	firebase *firebase.App
}
