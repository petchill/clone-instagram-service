package media

import "time"

type MediaMetaData struct {
	OwnerUserID     string
	Caption         string
	FileStorageLink string
	CreatedAt       time.Time
}
