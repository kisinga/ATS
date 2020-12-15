package user

type Interactor interface {
	GetUser()
}

type interactor struct {
	repository Repository
}

func NewIterator(repo Repository) Interactor {
	return &interactor{}
}

func (i *interactor) GetUser() {

}
