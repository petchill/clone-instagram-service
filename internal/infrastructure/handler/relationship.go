package handler

import (
	mRela "clone-instagram-service/internal/domain/model/relationship"
	eRela "clone-instagram-service/internal/domain/model/relationship/entity"
	eUser "clone-instagram-service/internal/domain/model/user/entity"
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
	g.POST("/unfollow", h.PostUnfollow)

}

func (h *relationshipHandler) PostFollow(c echo.Context) error {
	ctx := c.Request().Context()

	user, ok := c.Get("user").(eUser.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, "invalid user type")
	}

	followPayload := eRela.PostFollowRequestBody{}

	err := c.Bind(&followPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return nil
	}

	err = h.relationshipService.FollowUser(ctx, user.ID, followPayload.TargetUserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "ERROR", "error": err.Error()})
		return nil
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})

}

func (h *relationshipHandler) PostUnfollow(c echo.Context) error {
	ctx := c.Request().Context()

	user, ok := c.Get("user").(eUser.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, "invalid user type")
	}

	unfollowPayload := eRela.PostFollowRequestBody{}

	err := c.Bind(&unfollowPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return nil
	}

	err = h.relationshipService.UnFollowUser(ctx, user.ID, unfollowPayload.TargetUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "ERROR", "error": err.Error()})
		return nil
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}
