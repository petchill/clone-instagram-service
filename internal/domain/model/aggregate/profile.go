package aggregate

import (
	mMedia "clone-instagram-service/internal/domain/model/media"
	eUser "clone-instagram-service/internal/domain/model/user/entity"
)

type UserProfile struct {
	User       eUser.User             `json:"user"`
	Followers  []eUser.User           `json:"followers"`
	Followings []eUser.User           `json:"followings"`
	Posts      []mMedia.MediaMetaData `json:"posts"`
}
