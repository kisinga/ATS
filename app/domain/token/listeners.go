package token

import (
	"context"
	"sync"

	"github.com/kisinga/ATS/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Listeners keeps track of all the listeners to every topic
type Listeners struct {
	// list maps each topic to an index in the map as an int
	// each map index further contains a map of all the subscribers to this topic
	// @todo research on optimization where we can split the map into two maps and have the data duploicated
	// map[int]primitive.ObjectID            map[primitive.ObjectID]chan<- *models.Token
	// this would allow us to split read and write operations but introduce complexity for update operations
	// type ListenerList map[int]map[primitive.ObjectID]chan<- *models.Token

	list map[TopicNames]map[primitive.ObjectID]chan<- *models.Token
	mu   sync.RWMutex
}

// NewListeners returns an instance of active listeners
func NewListeners() *Listeners {
	list := make(map[int]map[primitive.ObjectID]chan<- *models.Token)
	for i := 0; i < LastEffectName; i++ {
		list[i] = make(map[primitive.ObjectID]chan<- *models.Token)
	}
	return &Listeners{
		list: make(map[TopicNames]map[primitive.ObjectID]chan<- *models.Token),
		mu:   sync.RWMutex{},
	}
}

// AddListener appends the provided value to the list of listeners
func (l *Listeners) AddListener(ctx context.Context, channel chan<- *models.Token, effectName TopicNames) primitive.ObjectID {
	// create a unique id to store the channel
	id := primitive.NewObjectID()
	l.mu.Lock()
	l.list[effectName][id] = channel
	l.mu.Unlock()
	// dont forget to remove the channel from the list once the context is done
	go func() {
		<-ctx.Done()
		l.mu.Lock()
		// make sure the id contains a value
		if _, ok := l.list[effectName][id]; ok {
			delete(l.list[effectName], id)
			l.mu.Unlock()
		} else {
			// This is a serious error, log it
		}
	}()
	return id
}
