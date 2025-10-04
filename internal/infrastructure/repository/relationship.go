package repository

import (
	"clone-instagram-service/internal/domain/model"
	eRela "clone-instagram-service/internal/domain/model/relationship/entity"
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

func (r *relationshipRepository) GetAllFollowingIDsByUserID(ctx context.Context, userID int) ([]int, error) {
	followingUserIDs := []int{}
	err := r.gormDB.
		Table("followings").
		Where("user_id = ?", userID).
		Pluck("target_user_id", &followingUserIDs).Error
	if err != nil {
		log.Printf("Error while getting following from database. Here's why: %v\n", err)
		return followingUserIDs, err
	}
	return followingUserIDs, nil
}

func (r *relationshipRepository) GetAllFollowerIDsByUserID(ctx context.Context, userID int) ([]int, error) {
	followerUserIDs := []int{}
	err := r.gormDB.
		Table("followings").
		Where("target_user_id = ?", userID).
		Pluck("user_id", &followerUserIDs).Error
	if err != nil {
		log.Printf("Error while getting following from database. Here's why: %v\n", err)
		return followerUserIDs, err
	}
	return followerUserIDs, nil
}

func (r *relationshipRepository) InsertFollowing(ctx context.Context, following eRela.Following) error {
	err := r.gormDB.Table("followings").Create(&following).Error
	if err != nil {
		log.Printf("Error while inserting following into database. Here's why: %v\n", err)
		return err
	}
	return nil
}

func (r *relationshipRepository) DeleteFollowingByUserIDAndTargetID(ctx context.Context, userID, targetID int) error {
	err := r.gormDB.
		Table("followings").
		Where("user_id = ? AND target_user_id = ?", userID, targetID).
		Delete(&eRela.Following{}).Error
	if err != nil {
		log.Printf("Error while deleting following from database. Here's why: %v\n", err)
		return err
	}
	return nil
}

func (r *relationshipRepository) PublishFollowingTopic(ctx context.Context, message eRela.FollowingTopicMessage) error {

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
