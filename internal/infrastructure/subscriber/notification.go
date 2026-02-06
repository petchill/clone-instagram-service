package subscriber

import (
	"clone-instagram-service/internal/domain/model"
	eRela "clone-instagram-service/internal/domain/model/relationship/entity"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type notificationSubscriber struct {
	kafkaConfig model.KafkaConfig
}

func NewNotificationSubscriber(kafkaConfig model.KafkaConfig) *notificationSubscriber {
	return &notificationSubscriber{
		kafkaConfig: kafkaConfig,
	}
}

func (sub *notificationSubscriber) subscribeFollowingWithUserID(ctx context.Context, userID int, callback func(ctx context.Context, message eRela.FollowingTopicMessage) error) error {
	fmt.Println(" start sub")
	topic := fmt.Sprintf("following-user-%d", userID)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        sub.kafkaConfig.Brokers,
		Topic:          topic,
		GroupID:        fmt.Sprintf("noti-%d", userID),
		StartOffset:    kafka.LastOffset, // Start at newest if no committed offset
		CommitInterval: 0,                // auto-commit interval
	})
	defer r.Close()
	for {
		select {
		case <-ctx.Done():
			// Context canceled — stop the function
			fmt.Println("stopped:", ctx.Err())
			return nil
		default:
			fmt.Println("he")
			m, err := r.ReadMessage(ctx)
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}

			log.Printf("[partition=%d offset=%d] key=%s value=%s", m.Partition, m.Offset, string(m.Key), string(m.Value))
			// Process the message
			var message eRela.FollowingTopicMessage
			err = json.Unmarshal(m.Value, &message) // Uncomment this line when the struct is filled
			if err != nil {
				log.Printf("Error unmarshaling message: %v", err)
				continue
			}
			err = callback(context.Background(), message)
			if err != nil {
				log.Printf("Error processing message: %v", err)
				continue
			}
			// Commit the message
			if err := r.CommitMessages(context.Background(), m); err != nil {
				continue
			}

		}
	}

}

func (sub *notificationSubscriber) SubscribeFollowingWithUserID(ctx context.Context, userID int, callback func(ctx context.Context, message eRela.FollowingTopicMessage) error) error {
	return sub.subscribeFollowingWithUserID(ctx, userID, callback)
}

func (sub *notificationSubscriber) SubscribeFollowing(callback func(ctx context.Context, message eRela.FollowingTopicMessage) error) error {
	go func() {
		topic := "following"
		fmt.Println("Starting Kafka consumer for topic:", topic)
		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers:        sub.kafkaConfig.Brokers,
			Topic:          topic,
			GroupID:        "notification-group",
			StartOffset:    kafka.LastOffset, // Start at newest if no committed offset
			CommitInterval: 0,                // auto-commit interval
		})
		defer func() {
			fmt.Println("reader closed")
			r.Close()
		}()
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {

				log.Printf("Error reading message: %v", err)
				continue
			}

			log.Printf("[partition=%d offset=%d] key=%s value=%s", m.Partition, m.Offset, string(m.Key), string(m.Value))
			// Process the message
			var message eRela.FollowingTopicMessage
			err = json.Unmarshal(m.Value, &message) // Uncomment this line when the struct is filled
			if err != nil {
				log.Printf("Error unmarshaling message: %v", err)
				continue
			}
			err = callback(context.Background(), message)
			if err != nil {
				log.Printf("Error processing message: %v", err)
				continue
			}
			// Commit the message
			if err := r.CommitMessages(context.Background(), m); err != nil {
				continue
			}
		}
	}()

	return nil
}
