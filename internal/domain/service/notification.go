package service

import (
	mNoti "clone-instagram-service/internal/domain/model/notification"
	eNoti "clone-instagram-service/internal/domain/model/notification/entity"
	eRela "clone-instagram-service/internal/domain/model/relationship/entity"
	mUser "clone-instagram-service/internal/domain/model/user"
	"context"
	"fmt"
)

type notificationService struct {
	notificationRepo mNoti.NotificationRepository
	userRepo         mUser.UserRepository
}

func NewNotificationService(notificationRepo mNoti.NotificationRepository, userRepo mUser.UserRepository) *notificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
		userRepo:         userRepo,
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

func (s *notificationService) SubscribeFollowing(ctx context.Context, followingMessage eRela.FollowingTopicMessage) error {
	// get following user name by user_id
	followingUser, exist, err := s.userRepo.GetUserByID(ctx, followingMessage.UserID)
	if err != nil {
		return err
	}
	if !exist {
		errMsgh := "user not found with id: " + string(followingMessage.UserID)
		return fmt.Errorf(errMsgh)
	}

	// verify if follower user exists
	_, exist, err = s.userRepo.GetUserByID(ctx, followingMessage.TargetUserID)
	if err != nil {
		return err
	}
	if !exist {
		errMsgh := "target user not found with id: " + string(followingMessage.TargetUserID)
		return fmt.Errorf(errMsgh)
	}
	// insert database
	notiMessage := fmt.Sprintf("User %s followed you", followingUser.Name)
	notification := eNoti.Notification{
		Type:        "following",
		Message:     notiMessage,
		OwnerUserID: followingMessage.TargetUserID,
		IsRead:      false,
	}
	err = s.notificationRepo.InsertNotification(ctx, notification)
	if err != nil {
		return err
	}

	// TODO: send notification to user via websocket (future work)
	return nil
}
