package apiKey

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
}

// RequiredRepos is the list of repositories required by this package to carry out the listed effects
type RequiredRepos struct {
	APIKey Repository
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

func (e *effects) listen() {
	go func() {
		for {
			select {
			case APIKey := <-e.topics.Created:
				// emit the value to all the listeners
				emitToListeners(e.listeners, TopicNames(crudModels.Create), &APIKey)
				fmt.Println(APIKey)
				break
			case APIKey := <-e.topics.Read:
				emitToListeners(e.listeners, TopicNames(crudModels.Read), &APIKey)
				fmt.Println(APIKey)
				break
			case APIKey := <-e.topics.Updated:
				emitToListeners(e.listeners, TopicNames(crudModels.Update), &APIKey)
				fmt.Println(APIKey)
				break
			case APIKey := <-e.topics.Deleted:
				emitToListeners(e.listeners, TopicNames(crudModels.Delete), &APIKey)
				fmt.Println(APIKey)
				break
			}
		}
	}()
}

func emitToListeners(listeners *Listeners, topic TopicNames, APIKey *models.APIKey) {
	listeners.mu.RLock()
	for _, listener := range listeners.list[topic] {
		go func(listener chan<- *models.APIKey) {
			listener <- APIKey
		}(listener)
	}
	listeners.mu.RUnlock()
}
