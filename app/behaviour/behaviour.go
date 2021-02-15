package behaviour

import (
	"context"

	"github.com/kisinga/ATS/app/behaviour/actions"
	"github.com/kisinga/ATS/app/behaviour/effects"
	"github.com/kisinga/ATS/app/behaviour/listeners"
	"github.com/kisinga/ATS/app/models"
)

type Behaviours struct {
	Token *effects.TokenEffects
}

func New() *Behaviours {
	actions := actions.New()
	myListeners := listeners.New()

	systemListener := make(chan *models.Token)
	myListeners.AddListener(context.Background(), systemListener, listeners.Create)

	go effects.TransmitToken(systemListener)

	return &Behaviours{
		Token: effects.New(actions, myListeners),
	}
}
