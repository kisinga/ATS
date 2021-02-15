package registry

import (
	"github.com/kisinga/ATS/app/apiKey"
	"github.com/kisinga/ATS/app/behaviour"
	"github.com/kisinga/ATS/app/meter"
	"github.com/kisinga/ATS/app/storage"
	"github.com/kisinga/ATS/app/token"
	"github.com/kisinga/ATS/app/user"
)

type Domain struct {
	Meter  meter.Interactor
	User   user.Interactor
	Token  token.Interactor
	APIKey apiKey.Interactor
}

func NewDomain(db *storage.Database) *Domain {
	behaviour := behaviour.New()
	meterRepo := meter.NewRepository(db)
	userRepo := user.NewRepository(db)
	apikeyRepo := apiKey.NewRepository(db)
	tokenRepo := token.NewRepository(db, behaviour.Token.Actions)

	return &Domain{
		Meter:  meter.NewIterator(meterRepo),
		User:   user.NewIterator(userRepo),
		Token:  token.NewIterator(tokenRepo, meterRepo),
		APIKey: apiKey.NewIterator(apikeyRepo),
	}
}
