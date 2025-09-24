package middleware

import (
	mAuth "clone-instagram-service/internal/domain/model/auth"
	mUser "clone-instagram-service/internal/domain/model/user"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type authMiddleware struct {
	authRepo mAuth.AuthRepository
	userRepo mUser.UserRepository
}

func NewAuthMiddleWare(authRepo mAuth.AuthRepository, userRepo mUser.UserRepository) *authMiddleware {
	return &authMiddleware{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (mw *authMiddleware) AuthWithUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		headers := c.Request().Header
		authToken := headers.Get("Authorization")
		if authToken == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "error": "authorization header is missing"})
		}
		token := strings.Split(authToken, " ")
		if len(token) != 2 || token[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "error": "invalid authorization header"})
		}
		authToken = token[1]
		authUser, err := mw.authRepo.GetUserInfoFromToken(ctx, authToken)
		if err != nil {
			fmt.Println("authToken", authToken)
			return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "error": err.Error()})
		}

		user, exists, err := mw.userRepo.GetUserByGoogleID(ctx, authUser.Sub)
		if err != nil || !exists {
			return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "error": "user not found"})
		}

		c.Set("user", user)

		// Call the next middleware or handler in the chain
		err = next(c)

		// Perform actions after the request is handled
		fmt.Println("After request:", c.Request().URL.Path)

		return err
	}
}
