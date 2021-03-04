package resources

import "github.com/andygeiss/utilities/logging"

// Member ...
type Member struct{}

// MemberAccess ...
type MemberAccess struct {
	logger logging.Logger
}

// GetMemberByID ...
func (a *MemberAccess) GetMemberByID(id string) *Member {
	a.logger.Print("resource.MemberAccess GetMemberByID called")
	return &Member{}
}

// NewMemberAccess ...
func NewMemberAccess(logger logging.Logger) *MemberAccess {
	return &MemberAccess{
		logger: logger,
	}
}
