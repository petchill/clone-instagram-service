package repository

import (
	mRelationship "clone-instagram-service/internal/domain/model/relationship"
	"context"
	"log"

	"gorm.io/gorm"
)

type relationshipRepository struct {
	gormDB *gorm.DB
}

func NewRelationshipRepository(db *gorm.DB) *relationshipRepository {
	return &relationshipRepository{db}
}

func (r *relationshipRepository) GetAllFollowingIDsByUserID(ctx context.Context, userID string) ([]string, error) {
	followingUserIDs := []string{}
	err := r.gormDB.
		Table("following").
		Where("user_id = ?", userID).
		Pluck("target_user_id", &followingUserIDs).Error
	if err != nil {
		log.Printf("Error while getting following from database. Here's why: %v\n", err)
		return followingUserIDs, err
	}
	return followingUserIDs, nil
}

func (r *relationshipRepository) GetAllFollowerIDsByUserID(ctx context.Context, userID string) ([]string, error) {
	followerUserIDs := []string{}
	err := r.gormDB.
		Table("following").
		Where("target_user_id = ?", userID).
		Pluck("user_id", &followerUserIDs).Error
	if err != nil {
		log.Printf("Error while getting following from database. Here's why: %v\n", err)
		return followerUserIDs, err
	}
	return followerUserIDs, nil
}

func (r *relationshipRepository) InsertFollowing(ctx context.Context, following mRelationship.Following) error {
	err := r.gormDB.Table("following").Create(&following).Error
	if err != nil {
		log.Printf("Error while inserting following into database. Here's why: %v\n", err)
		return err
	}
	return nil
}

func (r *relationshipRepository) DeleteFollowingByUserIDAndTargetID(ctx context.Context, userID, targetID string) error {
	err := r.gormDB.
		Table("following").
		Where("user_id = ? AND target_user_id = ?", userID, targetID).
		Delete(&mRelationship.Following{}).Error
	if err != nil {
		log.Printf("Error while deleting following from database. Here's why: %v\n", err)
		return err
	}
	return nil
}
