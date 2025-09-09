package handler

import (
	mAuth "clone-instagram-service/internal/domain/model/auth"
	mRela "clone-instagram-service/internal/domain/model/relationship"
	"net/http"

	"github.com/labstack/echo/v4"
)

type relationshipHandler struct {
	relationshipService mRela.RelationshipService
}

func NewRelationshipHandler(relationshipService mRela.RelationshipService) *relationshipHandler {
	return &relationshipHandler{
		relationshipService: relationshipService,
	}
}

func (h *relationshipHandler) RegisterRoutes(e *echo.Group) {
	g := e.Group("/following")
	g.POST("/follow", h.PostFollow)

}

func (h *relationshipHandler) PostFollow(c echo.Context) error {
	ctx := c.Request().Context()

	user, ok := c.Get("user").(mAuth.UserInfo)
	if !ok {
		c.JSON(http.StatusUnauthorized, "invalid user type")
	}

	followPayload := mRela.PostFollowRequestBody{}

	err := c.Bind(&followPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return nil
	}

	err = h.relationshipService.FollowUser(ctx, user.Sub, followPayload.TargetID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return nil
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})

}
