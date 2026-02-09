package service

import (
	mMedia "clone-instagram-service/internal/domain/model/media"
	mNewsFeed "clone-instagram-service/internal/domain/model/news_feed"
	aNewsFeed "clone-instagram-service/internal/domain/model/news_feed/aggregate"
	eNewsFeed "clone-instagram-service/internal/domain/model/news_feed/entity"
	mRela "clone-instagram-service/internal/domain/model/relationship"
	"context"
)

type newsFeedService struct {
	newsFeedRepo     mNewsFeed.NewsFeedRepository
	mediaRepo        mMedia.MediaRepository
	relationshipRepo mRela.RelationshipRepository
}

func NewNewsFeedService(newsFeedRepo mNewsFeed.NewsFeedRepository, mediaRepo mMedia.MediaRepository, relationshipRepo mRela.RelationshipRepository) *newsFeedService {
	return &newsFeedService{
		newsFeedRepo:     newsFeedRepo,
		mediaRepo:        mediaRepo,
		relationshipRepo: relationshipRepo,
	}
}

func (s *newsFeedService) GetNewsFeedByUserID(ctx context.Context, userID int) ([]aNewsFeed.NewsFeed, error) {
	feedFilter := eNewsFeed.NewsFeedFilter{
		Offset: 0,
		Limit:  10,
		UserID: userID,
	}

	feeds, err := s.newsFeedRepo.GetNewsFeedByFilter(ctx, feedFilter)
	if err != nil {
		return []aNewsFeed.NewsFeed{}, err
	}

	return feeds, err
}
