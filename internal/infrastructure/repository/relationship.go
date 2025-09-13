package repository

import (
	"clone-instagram-service/internal/domain/model"
	mRelationship "clone-instagram-service/internal/domain/model/relationship"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type relationshipRepository struct {
	gormDB           *gorm.DB
	followingEventWr *kafka.Writer
}

func NewRelationshipRepository(db *gorm.DB, kafkaConfig model.KafkaConfig) *relationshipRepository {

	followingEventWr := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      kafkaConfig.Brokers,
		Topic:        "following",
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: int(kafka.RequireAll), // good default
		Async:        false,
	})

	return &relationshipRepository{db, followingEventWr}
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

func (r *relationshipRepository) PublishFollowingTopic(ctx context.Context, message mRelationship.FollowingTopicMessage) error {

	messageJson, err := json.Marshal(message)
	if err != nil {
		log.Println("Error from marshal message in topic following")
		return err
	}

	err = r.followingEventWr.WriteMessages(ctx,
		kafka.Message{
			Value: messageJson,
		},
	)
	if err != nil {
		log.Println("Error from publish message in topic following")
		return err
	}

	return nil
}
