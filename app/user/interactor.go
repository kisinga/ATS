package user

import "github.com/kisinga/ATS/app/models"

type Interactor interface {
	GetUser(email string) (*models.User, error)
	AddUser(user models.NewUser) (*models.User, error)
}

type interactor struct {
	repository Repository
}

func NewIterator(repo Repository) Interactor {
	return &interactor{
		repository: repo,
	}
}

func (i *interactor) GetUser(email string) (*models.User, error) {
	return i.repository.Read(email)
}

func (i *interactor) AddUser(user models.NewUser) (*models.User, error) {
	newUser := models.User{
		Email: user.Email,
	}
	return i.repository.Create(newUser)
}

// func (i *interactor) ValidLogin(email string, tokenID string) bool {
// 	return true
// }
