package services

import (
	"github.com/andygeiss/template/messages"
	"github.com/andygeiss/template/resources"
	"github.com/andygeiss/utilities/logging"
	"github.com/andygeiss/utilities/messaging"
)

// MembershipManager ...
type MembershipManager struct {
	bus               messaging.Bus
	logger            logging.Logger
	memberAccess      *resources.MemberAccess
	regulationsEngine *RegulationsEngine
}

// Name ...
func (a *MembershipManager) Name() string {
	return "services.MembershipManager"
}

// Receive ...
func (a *MembershipManager) Receive(message interface{}) {
	switch message.(type) {
	case messages.VerifyApplication:
		id := "foo@bar.com"
		member := a.memberAccess.GetMemberByID(id)
		a.logger.Printf("%s Member %v", a.Name(), member)
		if valid := a.regulationsEngine.ValidateMember(member); valid {
			a.logger.Printf("%s Member is valid", a.Name())
			a.Send(messages.TradesmanOrContractorApproved{})
		} else {
			a.logger.Printf("%s Member is invalid!", a.Name())
			a.Send(messages.Error{})
		}
	}
}

// Send ...
func (a *MembershipManager) Send(message interface{}) {
	a.bus.Publish(message)
}

// NewMembershipManager ...
func NewMembershipManager(bus messaging.Bus, logger logging.Logger, regulationsEngine *RegulationsEngine, memberAccess *resources.MemberAccess) messaging.Actor {
	return &MembershipManager{
		bus:               bus,
		logger:            logger,
		memberAccess:      memberAccess,
		regulationsEngine: regulationsEngine,
	}
}
