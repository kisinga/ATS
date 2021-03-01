package domain

import (
	"github.com/kisinga/ATS/app/communication/sms"
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

func New(db *storage.Database, sms sms.SMS) *Domain {

	apiKeyTopics := apiKey.NewTopics(apiKey.NewCrudChannels())
	apiKeyRepo := apiKey.NewRepository(db, apiKeyTopics)
	apiKeyListeners := apiKey.NewListeners()
	apiKeyEffects := apiKey.NewEffects(&apiKey.Dependencies{APIKeyRepository: apiKeyRepo}, apiKeyTopics, apiKeyListeners)

	meterTopics := meter.NewTopics(meter.NewCrudChannels())
	meterRepo := meter.NewRepository(db, meterTopics)
	meterListeners := meter.NewListeners()
	meterEffects := meter.NewEffects(&meter.Dependencies{MeterRepository: meterRepo}, meterTopics, meterListeners)

	tokenTopics := token.NewTopics(token.NewCrudChannels())
	tokenRepo := token.NewRepository(db, tokenTopics)
	tokenListeners := token.NewListeners()
	tokenEffects := token.NewEffects(&token.Dependencies{TokenRepository: tokenRepo, MeterRepository: meterRepo, SMS: sms}, tokenTopics, tokenListeners)

	userTopics := user.NewTopics(user.NewCrudChannels())
	userRepo := user.NewRepository(db, userTopics)
	userListeners := user.NewListeners()
	userEffects := user.NewEffects(&user.Dependencies{UserRepository: userRepo}, userTopics, userListeners)

	return &Domain{
		APIKey: apiKey.NewIterator(apiKeyRepo, apiKeyEffects),
		Meter:  meter.NewIterator(meterRepo, meterEffects),
		Token:  token.NewIterator(tokenRepo, tokenEffects),
		User:   user.NewIterator(userRepo, userEffects),
	}

}
