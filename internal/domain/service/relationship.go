package service

import (
	mRelationship "clone-instagram-service/internal/domain/model/relationship"
	"context"
	"fmt"
)

type relationshipService struct {
	relationshipRepo mRelationship.RelationshipRepository
}

func NewRelationshipService(relationshipRepo mRelationship.RelationshipRepository) *relationshipService {
	return &relationshipService{
		relationshipRepo: relationshipRepo,
	}
}

func (s *relationshipService) FollowUser(ctx context.Context, userID string, targetUserID string) error {
	// TODO: should validate target_id is exist in user db

	followingIds, err := s.relationshipRepo.GetAllFollowingIDsByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if stringIsContained(followingIds, targetUserID) {
		return fmt.Errorf("this following is already exists")
	}

	err = s.relationshipRepo.InsertFollowing(ctx, mRelationship.Following{
		UserId:       userID,
		TargetUserID: targetUserID,
	})

	if err != nil {
		return err
	}

	return nil
}

func stringIsContained(list []string, target string) bool {
	for _, str := range list {
		if str == target {
			return true
		}
	}
	return false
}
