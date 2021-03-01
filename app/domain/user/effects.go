package user

import (
	"fmt"

	"github.com/kisinga/ATS/app/domain/crudModels"
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

// Dependencies is the list of dependencies required by this package to carry out the listed effects
type Dependencies struct {
	UserRepository Repository
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
			case User := <-e.topics.Created:
				// emit the value to all the listeners
				emitToListeners(e.listeners, TopicNames(crudModels.Create), &User)
				fmt.Println(User)
				break
			case User := <-e.topics.Read:
				emitToListeners(e.listeners, TopicNames(crudModels.Read), &User)
				fmt.Println(User)
				break
			case User := <-e.topics.Updated:
				emitToListeners(e.listeners, TopicNames(crudModels.Update), &User)
				fmt.Println(User)
				break
			case User := <-e.topics.Deleted:
				emitToListeners(e.listeners, TopicNames(crudModels.Delete), &User)
				fmt.Println(User)
				break
			}
		}
	}()
}

func emitToListeners(listeners *Listeners, topic TopicNames, User *models.User) {
	listeners.mu.RLock()
	for _, listener := range listeners.list[topic] {
		go func(listener chan<- *models.User) {
			listener <- User
		}(listener)
	}
	listeners.mu.RUnlock()
}
