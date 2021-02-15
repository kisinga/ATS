package effects

import (
	"fmt"

	"github.com/kisinga/ATS/app/behaviour/actions"
	"github.com/kisinga/ATS/app/behaviour/listeners"
	"github.com/kisinga/ATS/app/models"
)

// TokenEffects are all the posible operations performed whenever a specific action even occurs
type TokenEffects struct {
	Actions   *actions.TokenActions
	listeners *listeners.TokenListeners
}

// New creates an instance of TokenEffects given a pointer to the actions and listeners
func New(actions *actions.TokenActions, listeners *listeners.TokenListeners) *TokenEffects {
	effects := TokenEffects{
		Actions:   actions,
		listeners: listeners,
	}
	go effects.listen()
	return &effects
}

func (e TokenEffects) listen() {
	for {
		select {
		case key := <-e.Actions.GetCreate():
			e.listeners.Mu.Lock()
			for _, listener := range e.listeners.Create {
				go func(l chan<- *models.Token) {
					l <- key
				}(listener)
			}
			e.listeners.Mu.Unlock()
		}
	}
}

// TransmitToken a special system listener that sends the token downstram to the CIU whenever a new token is created
func TransmitToken(channel <-chan *models.Token) {
	for {
		select {
		case token := <-channel:
			fmt.Println(token)
			// TransmitToken(*token)
		}
	}
}
