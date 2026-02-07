package handler

import (
	mNewsFeed "clone-instagram-service/internal/domain/model/news_feed"
	eUser "clone-instagram-service/internal/domain/model/user/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type newsFeedHandler struct {
	newsFeedService mNewsFeed.NewsFeedService
}

func NewNewsFeedHandler(newsFeedService mNewsFeed.NewsFeedService) *newsFeedHandler {
	return &newsFeedHandler{newsFeedService: newsFeedService}
}

func (h *newsFeedHandler) RegisterRoutes(e *echo.Group) {
	e.GET("/news", h.GetNewsFeed)
}

func (h *newsFeedHandler) GetNewsFeed(c echo.Context) error {
	ctx := c.Request().Context()

	user, ok := c.Get("user").(eUser.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, "invalid user type")
	}

	feed, err := h.newsFeedService.GetNewsFeedByUserID(ctx, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "ERROR", "error": err.Error()})
		return nil
	}

	return c.JSON(http.StatusOK, feed)
}
