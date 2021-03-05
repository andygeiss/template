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

// Name ...
func (a *Tradesman) Name() string {
	return "clients.Tradesman"
}

// Receive ...
func (a *Tradesman) Receive(message interface{}) {
	switch m := message.(type) {
	case messages.ApplyForMembership:
		a.Send(messages.VerifyApplication{})
	case messages.Error:
		a.logger.Printf("%s Error %v", a.Name(), m)
	case messages.TradesmanOrContractorApproved:
		a.logger.Printf("%s Result %v", a.Name(), m)
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
