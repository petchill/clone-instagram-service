package service

import (
	mRelationship "clone-instagram-service/internal/domain/model/relationship"
	eRela "clone-instagram-service/internal/domain/model/relationship/entity"
	"context"
	"fmt"
	"log"
	"time"
)

type relationshipService struct {
	relationshipRepo mRelationship.RelationshipRepository
}

func NewRelationshipService(relationshipRepo mRelationship.RelationshipRepository) *relationshipService {
	return &relationshipService{
		relationshipRepo: relationshipRepo,
	}
}

func (s *relationshipService) FollowUser(ctx context.Context, userID int, targetUserID int) error {
	// TODO: should validate target_id is exist in user db

	followingIds, err := s.relationshipRepo.GetAllFollowingIDsByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if IntIsContained(followingIds, targetUserID) {
		err = fmt.Errorf("this following is already exists")
		log.Printf("Error: %s", err)
		return err
	}

	err = s.relationshipRepo.InsertFollowing(ctx, eRela.Following{
		UserId:       userID,
		TargetUserID: targetUserID,
	})

	if err != nil {
		return err
	}

	topicMessage := eRela.FollowingTopicMessage{
		UserID:       userID,
		TargetUserID: targetUserID,
		CreatedAt:    time.Now(),
	}

	err = s.relationshipRepo.PublishFollowingTopic(ctx, topicMessage)

	if err != nil {
		return err
	}

	return nil
}

func IntIsContained(list []int, target int) bool {
	for _, i := range list {
		if i == target {
			return true
		}
	}
	return false
}

func stringIsContained(list []string, target string) bool {
	for _, str := range list {
		if str == target {
			return true
		}
	}
	return false
}
