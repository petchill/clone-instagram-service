package entity

import "time"

type Notification struct {
	ID          int
	Type        string
	Message     string
	OwnerUserID int
	IsRead      bool
	CreatedAt   time.Time
}
