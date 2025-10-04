package repository

import (
	eNoti "clone-instagram-service/internal/domain/model/notification/entity"
	"context"

	"gorm.io/gorm"
)

type notificationRepository struct {
	gormDB *gorm.DB
}

func NewNotificationRepository(gormDB *gorm.DB) *notificationRepository {
	return &notificationRepository{
		gormDB: gormDB,
	}
}

func (r *notificationRepository) InsertNotification(ctx context.Context, notification eNoti.Notification) error {
	err := r.gormDB.Table("notification").Create(&notification).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *notificationRepository) GetAllNotificationsByUserID(ctx context.Context, userID int) ([]eNoti.Notification, error) {

	notifications := []eNoti.Notification{}
	err := r.gormDB.
		Table("notification").
		Where("owner_user_id = ?", userID).
		Order("created_at DESC").
		Find(&notifications).Error
	if err != nil {
		return notifications, err
	}
	return notifications, nil
}

func (r *notificationRepository) GetNotficationByID(ctx context.Context, notificationID int) (eNoti.Notification, bool, error) {
	notification := eNoti.Notification{}
	err := r.gormDB.First(&notification, "id = ?", notificationID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return eNoti.Notification{}, false, nil
		}
		return notification, false, err
	}
	return notification, true, nil
}
func (r *notificationRepository) MarkAllNotificationsAsReadByUserID(ctx context.Context, userID int) error {
	err := r.gormDB.Table("notification").
		Where("owner_user_id = ?", userID).
		Update("is_read", true).Error
	if err != nil {
		return err
	}
	return nil
}
