package user

import (
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

func NewRepository(database *storage.Database) Repository {
	return &repository{database}
}

type repository struct {
	db *storage.Database
}

func (r repository) Create(models.User) (*models.User, error) {
	return r.db.Client.Database(env).Collection("posts")
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
