package news_feed

import (
	aNewsFeed "clone-instagram-service/internal/domain/model/news_feed/aggregate"
	eNewsFeed "clone-instagram-service/internal/domain/model/news_feed/entity"
	"context"
)

type NewsFeedRepository interface {
	GetNewsFeedByFilter(ctx context.Context, filter eNewsFeed.NewsFeedFilter) ([]aNewsFeed.NewsFeed, error)
}
