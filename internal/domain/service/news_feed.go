package service

import (
	mMedia "clone-instagram-service/internal/domain/model/media"
	mRela "clone-instagram-service/internal/domain/model/relationship"
	"context"
)

type newsFeedService struct {
	mediaRepo        mMedia.MediaRepository
	relationshipRepo mRela.RelationshipRepository
}

func NewNewsFeedService(mediaRepo mMedia.MediaRepository, relationshipRepo mRela.RelationshipRepository) *newsFeedService {
	return &newsFeedService{
		mediaRepo:        mediaRepo,
		relationshipRepo: relationshipRepo,
	}
}

func (s *newsFeedService) GetNewsFeedByUserID(ctx context.Context, userID int) ([]mMedia.MediaMetaData, error) {
	followingIDs, err := s.relationshipRepo.GetAllFollowingIDsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	feedMedias := make([]mMedia.MediaMetaData, 0)
	for _, followingID := range followingIDs {
		medias, err := s.mediaRepo.GetMediasByOwnerUserID(ctx, followingID)
		if err != nil {
			return nil, err
		}
		feedMedias = append(feedMedias, medias...)
	}

	return feedMedias, nil
}
