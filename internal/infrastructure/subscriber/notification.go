package subscriber

import (
	"clone-instagram-service/internal/domain/model"
	eRela "clone-instagram-service/internal/domain/model/relationship/entity"
	"context"
	"encoding/json"
	"log"
	"time"

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

func (sub *notificationSubscriber) SubscribeFollowing(callback func(ctx context.Context, message eRela.FollowingTopicMessage) error) error {
	topic := "following"
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        sub.kafkaConfig.Brokers,
		Topic:          topic,
		StartOffset:    kafka.FirstOffset, // Start at newest if no committed offset
		CommitInterval: time.Second,       // auto-commit interval
	})
	defer r.Close()

	go func() {
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
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
