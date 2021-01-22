package app

import (
	"context"

	"github.com/kisinga/ATS/app/models"
	"github.com/kisinga/ATS/app/registry"
)

// TransmitToken sends the token downstram to the CIU
func TransmitToken(token models.Token) error {
	return nil
}
func listenForNewTokens(domain *registry.Domain) {
	listener := make(chan *models.Token, 0)
	domain.Token.ListenForNew(context.Background(), listener)
	go handleTokenEmission(listener)
}

func handleTokenEmission(channel <-chan *models.Token) {
	for {
		select {
		case token := <-channel:
			TransmitToken(*token)
		}
	}
}
