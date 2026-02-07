package news_feed

import (
	mMedia "clone-instagram-service/internal/domain/model/media"
	"context"
)

type NewsFeedService interface {
	GetNewsFeedByUserID(ctx context.Context, userID int) ([]mMedia.MediaMetaData, error)
}
