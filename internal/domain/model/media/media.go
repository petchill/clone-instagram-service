package media

import "time"

type MediaMetaData struct {
	ID              int `gorm:"omitempty"`
	OwnerUserID     string
	Caption         string
	FileStorageLink string
	CreatedAt       time.Time
}
