package handler

import (
	mUser "clone-instagram-service/internal/domain/model/user"
	eUser "clone-instagram-service/internal/domain/model/user/entity"
	"fmt"
	"net/http"
	"strconv"

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
	e.GET("/user/:id", h.GetUserByID)
	e.GET("/user/search", h.GetUsersByPartialNameOrEmail)
}

func (h *userHandler) GetUserProfile(c echo.Context) error {
	ctx := c.Request().Context()

	user, ok := c.Get("user").(eUser.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, "invalid user type")
	}

	fmt.Println("user", user)

	profile, err := h.userService.GetUserProfileByGoogleSubID(ctx, user.GoogleSubID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "ERROR", "error": err.Error()})
		return nil
	}

	return c.JSON(http.StatusOK, profile)
}

func (h *userHandler) GetUsersByPartialNameOrEmail(c echo.Context) error {
	ctx := c.Request().Context()
	searchText := c.QueryParam("search-text")
	fmt.Println("c.Param", searchText)

	resultUsers, err := h.userService.SearchUsersByNameOrEmail(ctx, searchText)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "ERROR", "error": err.Error()})
		return nil
	}

	return c.JSON(http.StatusOK, resultUsers)
}

func (h *userHandler) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "ERROR", "error": "id must be number"})
	}

	profile, err := h.userService.GetUserProfileByUserID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "ERROR", "error": err.Error()})
		return nil
	}

	return c.JSON(http.StatusOK, profile)

}
