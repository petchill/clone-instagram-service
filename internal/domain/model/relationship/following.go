package relationship

import "time"

type Following struct {
	ID           int       `json:"id,omitempty" gorm:"omitempty"`
	UserId       int       `json:"user_id"`
	TargetUserID int       `json:"target_user_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type PostFollowRequestBody struct {
	TargetID int `json:"target_id"`
}

type FollowingTopicMessage struct {
	UserID    int       `json:"user_id"`
	TargetID  int       `json:"target_id"`
	CreatedAt time.Time `json:"created_at"`
}
