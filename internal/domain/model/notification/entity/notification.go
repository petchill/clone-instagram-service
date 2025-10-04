package entity

import "time"

type Notification struct {
	ID          int       `gorm:"omitempty"`
	Type        string    `gorm:"type"`
	Message     string    `gorm:"message"`
	OwnerUserID int       `gorm:"owner_user_id"`
	IsRead      bool      `gorm:"is_read"`
	CreatedAt   time.Time `gorm:"created_at"`
}
