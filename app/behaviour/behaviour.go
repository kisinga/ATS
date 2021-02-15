package behaviour

import (
	"context"

	"github.com/kisinga/ATS/app/apiKey"
	"github.com/kisinga/ATS/app/behaviour/actions"
	"github.com/kisinga/ATS/app/behaviour/effects"
	"github.com/kisinga/ATS/app/behaviour/listeners"
	"github.com/kisinga/ATS/app/models"
)

type Behaviours struct {
	Token  *effects.TokenEffects
	APIKey *effects.APIKeyEffects
}

func New() *Behaviours {
	tokenActions := actions.NewTokenActions()
	APIKeyActions := actions.NewAPIKeyActions()
	tokenListeners := listeners.NewTokenListeners()
	APIKeyListeners := listeners.NewAPIKeyListeners()

	systemTokensListener := make(chan *models.Token)
	systemAPIKeyListener := make(chan *models.APIKey)

	tokenListeners.AddListener(context.Background(), systemTokensListener, listeners.TokenCreate)
	APIKeyListeners.AddListener(context.Background(), systemAPIKeyListener, listeners.KeyCreate)

	go effects.TransmitToken(systemTokensListener)

	go effects.UpdateCache(apiKey.UpdateCache, systemAPIKeyListener)

	return &Behaviours{
		Token:  effects.NewTokenEffects(tokenActions, tokenListeners),
		APIKey: effects.NewAPIKeyEffects(APIKeyActions, APIKeyListeners),
	}
}
