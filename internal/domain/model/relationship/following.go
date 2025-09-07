package relationship

import "time"

type Following struct {
	ID          int       `json:"id,omitempty" gorm:"omitempty"`
	UserId      int64     `json:"user_id"`
	FollowingId int64     `json:"following_id"`
	CreatedAt   time.Time `json:"created_at"`
}
