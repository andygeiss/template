package clients

import (
	"github.com/andygeiss/template/messages"
	"github.com/andygeiss/utilities/logging"
	"github.com/andygeiss/utilities/messaging"
)

// Tradesman ...
type Tradesman struct {
	bus    messaging.Bus
	logger logging.Logger
}

// Receive ...
func (a *Tradesman) Receive(message interface{}) {
	switch message.(type) {
	case messages.ApplyForMembership:
		a.logger.Print("clients.Tradesman ApplyForMembership received")
		a.logger.Print("clients.Tradesman VerifyApplication sent")
		a.Send(messages.VerifyApplication{})
	case messages.Error:
		a.logger.Print("clients.Tradesman Error received")
	case messages.TradesmanOrContractorApproved:
		a.logger.Print("clients.Tradesman TradesmanOrContractorApproved received")
		a.logger.Print("clients.Tradesman Result displayed")
	}
}

// Send ...
func (a *Tradesman) Send(message interface{}) {
	a.bus.Publish(message)
}

// NewTradesman ...
func NewTradesman(bus messaging.Bus, logger logging.Logger) messaging.Actor {
	return &Tradesman{
		bus:    bus,
		logger: logger,
	}
}
