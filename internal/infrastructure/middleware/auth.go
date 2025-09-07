package middleware

import (
	mAuth "clone-instagram-service/internal/domain/model/auth"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type authMiddleware struct {
	authRepo mAuth.AuthRepository
}

func NewAuthMiddleWare(authRepo mAuth.AuthRepository) *authMiddleware {
	return &authMiddleware{
		authRepo: authRepo,
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
		userInfo, err := mw.authRepo.GetUserInfoFromToken(ctx, authToken)
		if err != nil {
			fmt.Println("authToken", authToken)
			return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "error": err.Error()})
		}

		c.Set("user", userInfo)

		// Call the next middleware or handler in the chain
		err = next(c)

		// Perform actions after the request is handled
		fmt.Println("After request:", c.Request().URL.Path)

		return err
	}
}
