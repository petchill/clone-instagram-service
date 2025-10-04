package notification

import (
	eNoti "clone-instagram-service/internal/domain/model/notification/entity"
	"context"
)

type NotificationRepository interface {
	InsertNotification(ctx context.Context, notification eNoti.Notification) error
	GetAllNotificationsByUserID(ctx context.Context, userID int) ([]eNoti.Notification, error)
	GetNotficationByID(ctx context.Context, notificationID int) (eNoti.Notification, error)
	MarkAllNotificationsAsReadByUserID(ctx context.Context, userID int) error
}
