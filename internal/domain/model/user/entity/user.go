package user

import "time"

type User struct {
	ID          int       `json:"id,omitempty" gorm:"omitempty"`
	GoogleSubID string    `json:"sub" gorm:"google_sub_id"`
	Name        string    `json:"name"`
	GivenName   string    `json:"given_name"`
	FamilyName  string    `json:"family_name"`
	Picture     string    `json:"picture"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
}
