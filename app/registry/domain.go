package registry

import (
	firebase "firebase.google.com/go"
	"github.com/kisinga/ATS/app/meter"
	"github.com/kisinga/ATS/app/storage"
	"github.com/kisinga/ATS/app/token"
	"github.com/kisinga/ATS/app/user"
)

type Domain struct {
	Meter meter.Interactor
	User  user.Interactor
	Token token.Interactor
}

func NewDomain(db *storage.Database, firebase *firebase.App) *Domain {
	meterRepo := meter.NewRepository(db)
	userRepo := user.NewRepository(db, firebase)
	tokenRepo := token.NewRepository(db)
	return &Domain{
		Meter: meter.NewIterator(meterRepo),
		User:  user.NewIterator(userRepo),
		Token: token.NewIterator(tokenRepo),
	}
}
