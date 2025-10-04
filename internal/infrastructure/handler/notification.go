package handler

import (
	mNoti "clone-instagram-service/internal/domain/model/notification"
	eUser "clone-instagram-service/internal/domain/model/user/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type notificationHandler struct {
	notificationService mNoti.NotificationService
}

func NewNotificationHandler(notificationService mNoti.NotificationService) *notificationHandler {
	return &notificationHandler{notificationService: notificationService}
}

func (h *notificationHandler) RegisterRoutes(e *echo.Group) {
	e.GET("/notifications", h.GetAllNotifications)
	e.POST("/notifications/mark-all-as-read", h.MarkAllAsRead)
}

// GetAllNotifications godoc
func (h *notificationHandler) GetAllNotifications(c echo.Context) error {
	ctx := c.Request().Context()

	user, ok := c.Get("user").(eUser.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, "invalid user type")
	}

	notifications, err := h.notificationService.GetAllNotificationsByUserID(ctx, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "ERROR", "error": err.Error()})
		return nil
	}

	return c.JSON(http.StatusOK, notifications)
}

// MarkAllAsRead godoc
func (h *notificationHandler) MarkAllAsRead(c echo.Context) error {
	ctx := c.Request().Context()

	user, ok := c.Get("user").(eUser.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, "invalid user type")
	}

	err := h.notificationService.MarkAllNotificationsAsReadByUserID(ctx, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"status": "ERROR", "error": err.Error()})
		return nil
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}
