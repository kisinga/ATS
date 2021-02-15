package actions

import (
	"github.com/kisinga/ATS/app/models"
)

// APIKeyActions is the list of all possible operations that might have effects attached to them
type APIKeyActions struct {
	create chan *models.APIKey
}

// NewAPIKey creates an instance of TokenActions
func NewAPIKeyActions() *APIKeyActions {
	return &APIKeyActions{
		create: make(chan *models.APIKey),
	}
}

// GetCreate returns the local variable create whenever absolutely necessary to use it in a external package
// like for instance effects package
func (a APIKeyActions) GetCreate() chan *models.APIKey {
	return a.create
}

// EmitCreate should be called whenever a create action occurs
// This is done within package scope so that we do not run the rist of ever having a blocking call to emit a value to the channel
func (a APIKeyActions) EmitCreate(t *models.APIKey) {
	go func() {
		a.create <- t
	}()
}
