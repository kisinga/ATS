package listeners

import (
	"context"
	"sync"

	"github.com/kisinga/ATS/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//APIKeyTopics defines a subscribable event type
type APIKeyTopics int

// KeyCreate is a topic emitted when a token is created
const KeyCreate APIKeyTopics = iota

// APIKeyListeners is the list of subscribers to a certain effecttype(topic)
type APIKeyListeners struct {
	Create map[primitive.ObjectID]chan<- *models.APIKey
	Mu     sync.Mutex
}

// New creates an instance of listeners
func NewAPIKeyListeners() *APIKeyListeners {
	return &APIKeyListeners{
		Create: make(map[primitive.ObjectID]chan<- *models.APIKey),
		Mu:     sync.Mutex{},
	}
}

// AddListener adds a listener to the list given the effect type
func (t *APIKeyListeners) AddListener(ctx context.Context, listener chan<- *models.APIKey, topic APIKeyTopics) {
	id := primitive.NewObjectID()
	switch topic {
	case KeyCreate:
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
