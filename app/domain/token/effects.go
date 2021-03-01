package token

import (
	"fmt"

	"github.com/kisinga/ATS/app/domain/crudModels"
	"github.com/kisinga/ATS/app/models"
)

// Effects acts as a dependency injection struct
type effects struct {
	topics    *Topics
	listeners *Listeners
	repos     *RequiredRepos
}

type Effects interface {
	listen()
	Listeners() *Listeners
}

// RequiredRepos is the list of repositories required by this package to carry out the listed effects
type RequiredRepos struct {
	Token Repository
}

func NewEffects(repos *RequiredRepos, topics *Topics, listeners *Listeners) Effects {
	effects := effects{
		topics:    topics,
		listeners: listeners,
		repos:     repos,
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
			case Token := <-e.topics.Created:
				// emit the value to all the listeners
				emitToListeners(e.listeners, TopicNames(crudModels.Create), &Token)
				fmt.Println(Token)
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
