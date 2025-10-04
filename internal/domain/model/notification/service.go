package notification

import (
	eNoti "clone-instagram-service/internal/domain/model/notification/entity"
	"context"
)

type NotificationService interface {
	GetAllNotificationsByUserID(ctx context.Context, userID int) ([]eNoti.NotificationResponse, error)
	MarkAllNotificationsAsReadByUserID(ctx context.Context, userID int) error
}
