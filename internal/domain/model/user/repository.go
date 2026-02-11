package user

import (
	eUser "clone-instagram-service/internal/domain/model/user/entity"
	"context"
)

type UserRepository interface {
	GetUserByGoogleID(ctx context.Context, googleID string) (eUser.User, bool, error)
	InsertUser(ctx context.Context, user eUser.User) error
	GetFollowingUsersByUserID(ctx context.Context, userID int) ([]eUser.User, error)
	GetFollowerUsersByUserID(ctx context.Context, userID int) ([]eUser.User, error)
	GetUserByID(ctx context.Context, userID int) (eUser.User, bool, error)
	GetUserByNameOrEmail(ctx context.Context, searchText string) ([]eUser.User, error)
}
