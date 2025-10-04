package media

import "time"

type MediaMetaData struct {
	ID              int       `gorm:"omitempty" json:"id"`
	OwnerUserID     int       `json:"owner_user_id"`
	Caption         string    `json:"caption"`
	FileStorageLink string    `json:"file_storage_link"`
	CreatedAt       time.Time `json:"created_at"`
}
