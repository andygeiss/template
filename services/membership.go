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

// Receive ...
func (a *MembershipManager) Receive(message interface{}) {
	switch message.(type) {
	case messages.VerifyApplication:
		a.logger.Print("core.MembershipManager VerifyApplication received")
		id := "foo@bar.com"
		member := a.memberAccess.GetMemberByID(id)
		if valid := a.regulationsEngine.ValidateMember(member); valid {
			a.Send(messages.TradesmanOrContractorApproved{})
			a.logger.Print("core.MembershipManager TradesmanOrContractorApproved sent")
		} else {
			a.Send(messages.Error{})
			a.logger.Print("core.MembershipManager Error sent")
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
