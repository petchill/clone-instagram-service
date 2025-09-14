package user

import (
	mAuth "clone-instagram-service/internal/domain/model/auth"
	"context"
)

type UserService interface {
	LoginWithGoogleAccessCode(ctx context.Context, googleAccessCode string) (mAuth.AccessCodeResponse, error)
}
