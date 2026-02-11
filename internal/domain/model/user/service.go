package user

import (
	mAgg "clone-instagram-service/internal/domain/model/aggregate"
	mAuth "clone-instagram-service/internal/domain/model/auth"
	eUser "clone-instagram-service/internal/domain/model/user/entity"
	"context"
)

type UserService interface {
	LoginWithGoogleAccessCode(ctx context.Context, googleAccessCode string) (mAuth.AccessCodeResponse, error)
	GetUserProfileByGoogleSubID(ctx context.Context, googleSubID string) (mAgg.UserProfile, error)
	SearchUsersByNameOrEmail(ctx context.Context, searchText string) ([]eUser.User, error)
}
