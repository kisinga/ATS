package token

import (
	"context"
	"fmt"

	"github.com/kisinga/ATS/app/communication/sms"
	"github.com/kisinga/ATS/app/domain/crudModels"
	"github.com/kisinga/ATS/app/domain/meter"
	"github.com/kisinga/ATS/app/models"
)

// Effects acts as a dependency injection struct
type effects struct {
	topics    *Topics
	listeners *Listeners
	deps      *Dependencies
}

type Effects interface {
	listen()
	Listeners() *Listeners
}

// RequiredRepos is the list of repositories required by this package to carry out the listed effects
type Dependencies struct {
	TokenRepository Repository
	MeterRepository meter.Repository
	SMS             sms.SMS
}

func NewEffects(deps *Dependencies, topics *Topics, listeners *Listeners) Effects {
	effects := effects{
		topics:    topics,
		listeners: listeners,
		deps:      deps,
	}
	effects.listen()
	return &effects
}

func (e *effects) Listeners() *Listeners {
	return e.listeners
}
func (e *effects) listen() {
	go func() {
		for {
			select {
			case token := <-e.topics.Created:
				// emit the value to all the listeners
				emitToListeners(e.listeners, TopicNames(crudModels.Create), &token)
				ctx := context.Background()
				meter, err := e.deps.MeterRepository.Read(ctx, token.MeterNumber)
				if err != nil {
					// Critical error, log
				}
				_, err = e.deps.SMS.Send(models.Text{Phone: meter.Phone, Message: "#" + token.String()})
				if err != nil {
					token.Status = models.StatusSent
					e.deps.TokenRepository.Update(ctx, token)
				}
				fmt.Println(token)
				break
			case Token := <-e.topics.Read:
				emitToListeners(e.listeners, TopicNames(crudModels.Read), &Token)
				fmt.Println(Token)
				break
			case Token := <-e.topics.Updated:
				emitToListeners(e.listeners, TopicNames(crudModels.Update), &Token)
				fmt.Println(Token)
				break
			case Token := <-e.topics.Deleted:
				emitToListeners(e.listeners, TopicNames(crudModels.Delete), &Token)
				fmt.Println(Token)
				break
			}
		}
	}()
}

func emitToListeners(listeners *Listeners, topic TopicNames, Token *models.Token) {
	listeners.mu.RLock()
	for _, listener := range listeners.list[topic] {
		go func(listener chan<- *models.Token) {
			listener <- Token
		}(listener)
	}
	listeners.mu.RUnlock()
}
