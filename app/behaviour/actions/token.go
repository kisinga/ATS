package actions

import (
	"github.com/kisinga/ATS/app/models"
)

// TokenActions is the list of all possible operations that might have effects attached to them
type TokenActions struct {
	create chan *models.Token
}

// New creates an instance of TokenActions
func New() *TokenActions {
	return &TokenActions{
		create: make(chan *models.Token),
	}
}

// GetCreate returns the local variable create whenever absolutely necessary to use it in a external package
// like for instance effects package
func (a TokenActions) GetCreate() chan *models.Token {
	return a.create
}

// EmitCreate should be called whenever a create action occurs
// This is done within package scope so that we do not run the rist of ever having a blocking call to emit a value to the channel
func (a TokenActions) EmitCreate(t *models.Token) {
	go func() {
		a.create <- t
	}()
}
