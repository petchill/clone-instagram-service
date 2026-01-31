package auth

import (
	eUser "clone-instagram-service/internal/domain/model/user/entity"
	"context"

	"github.com/labstack/echo/v4"
)

type AuthMiddleWare interface {
	GetUserInfoByAccessToken(ctx context.Context, accessToken string) (eUser.User, error)
	AuthWithUser(next echo.HandlerFunc) echo.HandlerFunc
}
