package meter

type Interactor interface {
}

type interactor struct {
	repository Repository
}

func NewIterator(repo Repository) Interactor {
	return &interactor{}
}
