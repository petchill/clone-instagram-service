package news_feed

import (
	aNewsFeed "clone-instagram-service/internal/domain/model/news_feed/aggregate"

	"context"
)

type NewsFeedService interface {
	GetNewsFeedByUserID(ctx context.Context, userID int) ([]aNewsFeed.NewsFeed, error)
}
