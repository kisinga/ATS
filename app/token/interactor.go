package token

type Interactor interface {
	GetToken()
}

type interactor struct {
	repository Repository
}

func NewIterator(repo Repository) Interactor {
	return &interactor{}
}

func (i *interactor) GetToken() {

}
