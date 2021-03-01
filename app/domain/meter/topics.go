package meter

import (
	"github.com/kisinga/ATS/app/domain/crudModels"
	"github.com/kisinga/ATS/app/models"
)

// TopicNames is the enumm representation of the Subscribabble topics
type TopicNames crudModels.RepoOperation

const (
	CustomTopicName = iota + 3
)

// LastEffectName is an alias to the largest int value of the custom effects
// This is used when iterating over the effects for mapping listeners
const LastEffectName = CustomTopicName

// Topics keeps the list of all possible operations that might produce effects
type Topics struct {
	*CrudChannels
}

// Emit is a helper function that ensures emissions are run in a goroutine to avoid blocking
func (t Topics) Emit(channel chan models.Meter, Meter models.Meter) {
	go func() {
		channel <- Meter
	}()
}

// NewTopics creates an instance of MeterActions
func NewTopics(channels *CrudChannels) *Topics {
	return &Topics{
		channels,
	}
}
