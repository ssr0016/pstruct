package userroles

import "context"

type Service interface {
	Assign(ctx context.Context, userID, roleID int) error
}
