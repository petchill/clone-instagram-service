package service

import (
	mNoti "clone-instagram-service/internal/domain/model/notification"
	eNoti "clone-instagram-service/internal/domain/model/notification/entity"
	"context"
)

type notificationService struct {
	notificationRepo mNoti.NotificationRepository
}

func NewNotificationService(notificationRepo mNoti.NotificationRepository) *notificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
	}
}

// GetAllNotificationsByUserID(ctx context.Context, userID int) ([]eNoti.Notification, error)
// 	MarkAllNotificationsAsReadByUserID(ctx context.Context, userID int) error

func (s *notificationService) GetAllNotificationsByUserID(ctx context.Context, userID int) ([]eNoti.NotificationResponse, error) {
	notifications, err := s.notificationRepo.GetAllNotificationsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	// Convert to NotificationResponse
	var notificationResponses []eNoti.NotificationResponse
	for _, notification := range notifications {
		notificationResponses = append(notificationResponses, convertNotificationToResponse(notification))
	}
	//
	return notificationResponses, nil
}

func convertNotificationToResponse(notification eNoti.Notification) eNoti.NotificationResponse {
	return eNoti.NotificationResponse{
		ID:        notification.ID,
		Type:      notification.Type,
		Message:   notification.Message,
		IsRead:    notification.IsRead,
		CreatedAt: notification.CreatedAt,
	}
}

func (s *notificationService) MarkAllNotificationsAsReadByUserID(ctx context.Context, userID int) error {
	return s.notificationRepo.MarkAllNotificationsAsReadByUserID(ctx, userID)
}
