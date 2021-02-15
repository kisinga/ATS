package effects

import (
	"github.com/kisinga/ATS/app/behaviour/actions"
	"github.com/kisinga/ATS/app/behaviour/listeners"
	"github.com/kisinga/ATS/app/models"
)

// TokenEffects are all the posible operations performed whenever a specific action even occurs
type APIKeyEffects struct {
	Actions   *actions.APIKeyActions
	Listeners *listeners.APIKeyListeners
}

type CacheUpdater func(channel chan *models.APIKey)

// New creates an instance of TokenEffects given a pointer to the actions and listeners
func NewAPIKeyEffects(actions *actions.APIKeyActions, listeners *listeners.APIKeyListeners) *APIKeyEffects {
	effects := APIKeyEffects{
		Actions:   actions,
		Listeners: listeners,
	}
	go effects.listenNewApiKey()
	return &effects
}

func (e APIKeyEffects) listenNewApiKey() {
	for {
		select {
		case key := <-e.Actions.GetCreate():
			e.Listeners.Mu.Lock()
			for _, listener := range e.Listeners.Create {
				go func(l chan<- *models.APIKey) {
					l <- key
				}(listener)
			}
			e.Listeners.Mu.Unlock()
		}
	}
}

// UpdateCache a special system listener that keeps the cache updated with the latest API Key
func UpdateCache(updater CacheUpdater, channel chan *models.APIKey) {
	updater(channel)
}
