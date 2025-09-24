package handler

import (
	mAuth "clone-instagram-service/internal/domain/model/auth"
	mUser "clone-instagram-service/internal/domain/model/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService mUser.UserService
}

func NewUserHandler(userService mUser.UserService) *userHandler {

	return &userHandler{userService: userService}
}

func (h *userHandler) RegisterRoutes(e *echo.Group) {
	e.GET("/user/profile", h.GetUserProfile)
}

func (h *userHandler) GetUserProfile(c echo.Context) error {
	ctx := c.Request().Context()

	authUser, ok := c.Get("user").(mAuth.UserInfo)
	if !ok {
		c.JSON(http.StatusUnauthorized, "invalid user type")
	}

	profile, err := h.userService.GetUserProfileByGoogleSubID(ctx, authUser.Sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "ERROR", "error": err.Error()})
		return nil
	}

	return c.JSON(http.StatusOK, profile)

}
