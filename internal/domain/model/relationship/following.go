package relationship

import "time"

type Following struct {
	ID           int       `json:"id,omitempty" gorm:"omitempty"`
	UserId       string    `json:"user_id"`
	TargetUserID string    `json:"target_user_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type PostFollowRequestBody struct {
	TargetID string `json:"target_id"`
}

type FollowingTopicMessage struct {
	UserID    string    `json:"user_id"`
	TargetID  string    `json:"target_id"`
	CreatedAt time.Time `json:"created_at"`
}
