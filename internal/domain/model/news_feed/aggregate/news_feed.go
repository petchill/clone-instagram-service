package aggregate

import (
	mMedia "clone-instagram-service/internal/domain/model/media"
	eUser "clone-instagram-service/internal/domain/model/user/entity"
)

type NewsFeed struct {
	Media mMedia.MediaMetaData `json:"media" gorm:"embedded;embeddedPrefix:media_"`
	User  eUser.User           `json:"user" gorm:"embedded;embeddedPrefix:user_"`
}
