package meter

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
	MeterRepository Repository
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
			case Meter := <-e.topics.Created:
				// emit the value to all the listeners
				emitToListeners(e.listeners, TopicNames(crudModels.Create), &Meter)
				fmt.Println(Meter)
				break
			case Meter := <-e.topics.Read:
				emitToListeners(e.listeners, TopicNames(crudModels.Read), &Meter)
				fmt.Println(Meter)
				break
			case Meter := <-e.topics.Updated:
				emitToListeners(e.listeners, TopicNames(crudModels.Update), &Meter)
				fmt.Println(Meter)
				break
			case Meter := <-e.topics.Deleted:
				emitToListeners(e.listeners, TopicNames(crudModels.Delete), &Meter)
				fmt.Println(Meter)
				break
			}
		}
	}()
}

func emitToListeners(listeners *Listeners, topic TopicNames, Meter *models.Meter) {
	listeners.mu.RLock()
	for _, listener := range listeners.list[topic] {
		go func(listener chan<- *models.Meter) {
			listener <- Meter
		}(listener)
	}
	listeners.mu.RUnlock()
}
