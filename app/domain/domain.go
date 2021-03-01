package domain

import (
	"github.com/kisinga/ATS/app/domain/apiKey"
	"github.com/kisinga/ATS/app/domain/meter"
	"github.com/kisinga/ATS/app/domain/token"
	"github.com/kisinga/ATS/app/domain/user"
	"github.com/kisinga/ATS/app/storage"
)

type Listeners struct {
	Meter  *meter.Listeners
	User   *user.Listeners
	Token  *token.Listeners
	APIKey *apiKey.Listeners
}
type Domain struct {
	Meter  meter.Interactor
	User   user.Interactor
	Token  token.Interactor
	APIKey apiKey.Interactor
}

func New(db *storage.Database) *Domain {
	apiKeyTopics := apiKey.NewTopics(apiKey.NewCrudChannels())
	apiKeyRepo := apiKey.NewRepository(db, apiKeyTopics)
	apiKeyListeners := apiKey.NewListeners()
	apiKeyEffects := apiKey.NewEffects(&apiKey.RequiredRepos{APIKey: apiKeyRepo}, apiKeyTopics, apiKeyListeners)

	meterTopics := meter.NewTopics(meter.NewCrudChannels())
	meterRepo := meter.NewRepository(db, meterTopics)
	meterListeners := meter.NewListeners()
	meterEffects := meter.NewEffects(&meter.RequiredRepos{Meter: meterRepo}, meterTopics, meterListeners)

	tokenTopics := token.NewTopics(token.NewCrudChannels())
	tokenRepo := token.NewRepository(db, tokenTopics)
	tokenListeners := token.NewListeners()
	tokenEffects := token.NewEffects(&token.RequiredRepos{Token: tokenRepo}, tokenTopics, tokenListeners)

	userTopics := user.NewTopics(user.NewCrudChannels())
	userRepo := user.NewRepository(db, userTopics)
	userListeners := user.NewListeners()
	userEffects := user.NewEffects(&user.RequiredRepos{User: userRepo}, userTopics, userListeners)

	return &Domain{
		APIKey: apiKey.NewIterator(apiKeyRepo, apiKeyEffects),
		Meter:  meter.NewIterator(meterRepo, meterEffects),
		Token:  token.NewIterator(tokenRepo, tokenEffects),
		User:   user.NewIterator(userRepo, userEffects),
	}

}
