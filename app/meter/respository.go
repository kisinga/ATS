package meter

import "github.com/kisinga/ATS/app/storage"

type Repository interface {
	Create()
	Read()
	Update()
	Delete()
}

type repository struct {
	db *storage.Database
}

func NewRepository(database *storage.Database) Repository {
	return &repository{db: database}
}
func (r repository) Create() {
}
func (r repository) Read() {
}
func (r repository) ReadMany() {
}
func (r repository) Update() {
}
func (r repository) Delete() {
}
