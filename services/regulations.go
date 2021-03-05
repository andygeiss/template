package services

import (
	"github.com/andygeiss/template/resources"
	"github.com/andygeiss/utilities/logging"
)

// RegulationsEngine ...
type RegulationsEngine struct {
	logger logging.Logger
}

// ValidateMember ...
func (a *RegulationsEngine) ValidateMember(member *resources.Member) (valid bool) {
	a.logger.Print("core.RegulationsEngine ValidateMember called")
	return true
}

// NewRegulationsEngine ...
func NewRegulationsEngine(logger logging.Logger) *RegulationsEngine {
	return &RegulationsEngine{
		logger: logger,
	}
}
