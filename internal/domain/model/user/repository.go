package user

import "context"

type UserRepository interface {
	GetUserByGoogleID(ctx context.Context, googleID string) (User, bool, error)
	InsertUser(ctx context.Context, user User) error
}
