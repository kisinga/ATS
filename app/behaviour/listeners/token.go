package listeners

import (
	"context"
	"sync"

	"github.com/kisinga/ATS/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TokenTopics defines a subscribable event type
type TokenTopics int

// TokenCreate is a topic emitted when a token is created
const TokenCreate TokenTopics = iota

// TokenListeners is the list of subscribers to a certain effecttype(topic)
type TokenListeners struct {
	Create map[primitive.ObjectID]chan<- *models.Token
	Mu     sync.Mutex
}

// NewTokenListeners creates an instance of listeners
func NewTokenListeners() *TokenListeners {
	return &TokenListeners{
		Create: make(map[primitive.ObjectID]chan<- *models.Token),
		Mu:     sync.Mutex{},
	}
}

// AddListener adds a listener to the list given the effect type
func (t *TokenListeners) AddListener(ctx context.Context, listener chan<- *models.Token, topic TokenTopics) {
	id := primitive.NewObjectID()
	switch topic {
	case TokenCreate:
		t.Mu.Lock()
		t.Create[id] = listener
		t.Mu.Unlock()
		go func() {
			<-ctx.Done()
			t.Mu.Lock()
			if _, ok := t.Create[id]; ok {
				delete(t.Create, id)
				t.Mu.Lock()
			}
		}()
		break
	}
}
