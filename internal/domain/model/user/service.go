package user

import (
	mAgg "clone-instagram-service/internal/domain/model/aggregate"
	mAuth "clone-instagram-service/internal/domain/model/auth"
	"context"
)

type UserService interface {
	LoginWithGoogleAccessCode(ctx context.Context, googleAccessCode string) (mAuth.AccessCodeResponse, error)
	GetUserProfileByGoogleSubID(ctx context.Context, googleSubID string) (mAgg.UserProfile, error)
}
