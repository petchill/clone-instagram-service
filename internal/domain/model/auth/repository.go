package auth

import (
	"context"

	"golang.org/x/oauth2"
)

type AuthRepository interface {
	GetUserInfoFromToken(ctx context.Context, accessToken string) (UserInfo, error)
	ExchangeCodeForToken(ctx context.Context, code string) (*oauth2.Token, error)
}
