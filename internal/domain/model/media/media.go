package media

import "time"

type MediaMetaData struct {
	ID              int `gorm:"omitempty"`
	OwnerUserID     int
	Caption         string
	FileStorageLink string
	CreatedAt       time.Time
}
