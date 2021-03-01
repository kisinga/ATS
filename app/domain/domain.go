package domain

import (
	"github.com/kisinga/ATS/app/domain/apiKey"
	"github.com/kisinga/ATS/app/domain/meter"
	"github.com/kisinga/ATS/app/domain/token"
	"github.com/kisinga/ATS/app/domain/user"
	"github.com/kisinga/ATS/app/storage"
)

type effects struct {
	apiKey apiKey.Effects
	user   user.Effects
	meter  meter.Effects
	token  token.Effects
}
type Listeners struct {
	Meter  *meter.Listeners
	User   *user.Listeners
	Token  *token.Listeners
	APIKey *apiKey.Listeners
}
type Domain struct {
	Meter     meter.Interactor
	User      user.Interactor
	Token     token.Interactor
	APIKey    apiKey.Interactor
	effects   effects
	Listeners Listeners
}

func New(db *storage.Database) *Domain {
	apiKeyTopics := apiKey.NewTopics(apiKey.NewCrudChannels())
	apiKeyRepo := apiKey.NewRepository(db, apiKeyTopics)
	apiKeyListeners := apiKey.NewListeners()

	meterTopics := meter.NewTopics(meter.NewCrudChannels())
	meterRepo := meter.NewRepository(db, meterTopics)
	meterListeners := meter.NewListeners()

	tokenTopics := token.NewTopics(token.NewCrudChannels())
	tokenRepo := token.NewRepository(db, tokenTopics)
	tokenListeners := token.NewListeners()

	userTopics := user.NewTopics(user.NewCrudChannels())
	userRepo := user.NewRepository(db, userTopics)
	userListeners := user.NewListeners()

	return &Domain{
		APIKey: apiKey.NewIterator(apiKeyRepo),
		effects: effects{
			apiKey: apiKey.NewEffects(&apiKey.RequiredRepos{APIKey: apiKeyRepo}, apiKeyTopics, apiKeyListeners),
			meter:  meter.NewEffects(&meter.RequiredRepos{Meter: meterRepo}, meterTopics, meterListeners),
			token:  token.NewEffects(&token.RequiredRepos{Token: tokenRepo}, tokenTopics, tokenListeners),
			user:   user.NewEffects(&user.RequiredRepos{User: userRepo}, userTopics, userListeners),
		},
	}

}
