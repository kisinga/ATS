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

func NewRepository(database *storage.Database, firebase *firebase.App) Repository {
	return &repository{database, firebase}
}

type repository struct {
	db       *storage.Database
	firebase *firebase.App
}

func (r repository) Create(models.User) (*models.User, error) {
	return nil, nil
}
func (r repository) Read(email string) (*models.User, error) {
	return nil, nil
}
func (r repository) ReadMany(start string, stop string) ([]models.User, error) {
	return nil, nil
}
func (r repository) Update(models.User) (*models.User, error) {
	return nil, nil
}
func (r repository) Delete(email string) (*models.User, error) {
	return nil, nil
}
