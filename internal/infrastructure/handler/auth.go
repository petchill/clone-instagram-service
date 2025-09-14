package handler

import (
	mAuth "clone-instagram-service/internal/domain/model/auth"
	mUser "clone-instagram-service/internal/domain/model/user"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	authRepo    mAuth.AuthRepository
	userService mUser.UserService
}

func NewAuthHandler(authRepo mAuth.AuthRepository, userService mUser.UserService) *authHandler {

	return &authHandler{authRepo: authRepo, userService: userService}
}

func (h *authHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/auth/accessToken", h.PostAccessCode)
	e.GET("/auth/user", h.GetUser)
}

func (h *authHandler) GetUser(c echo.Context) error {
	fmt.Println("enter")
	ctx := c.Request().Context()
	headers := c.Request().Header
	authToken := headers.Get("Authorization")
	fmt.Println("authToken", authToken)
	if authToken == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "error": "authorization header is missing"})
	}
	token := strings.Split(authToken, " ")
	if len(token) != 2 || token[0] != "Bearer" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "error": "invalid authorization header"})
	}
	authToken = token[1]

	userInfo, err := h.authRepo.GetUserInfoFromToken(ctx, authToken)
	if err != nil {
		fmt.Println("authToken", authToken)
		return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "error": err.Error()})
	}

	// TODO: check if the token is valid
	return c.JSON(http.StatusOK, userInfo)
}

func (h *authHandler) PostAccessCode(c echo.Context) error {
	ctx := c.Request().Context()
	payload := mAuth.AccessCodePayload{}
	if err := c.Bind(&payload); err != nil {
		return err
	}

	resp, err := h.userService.LoginWithGoogleAccessCode(ctx, payload.Code)
	if err != nil {

		return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}
