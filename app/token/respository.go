package token

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
