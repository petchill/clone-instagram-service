package notification

import (
	eRela "clone-instagram-service/internal/domain/model/relationship/entity"
	"context"
)

type NotificationSubscriber interface {
	SubscribeFollowing(callback func(ctx context.Context, message eRela.FollowingTopicMessage) error) error
	SubscribeFollowingWithUserID(ctx context.Context, userID int, callback func(ctx context.Context, message eRela.FollowingTopicMessage) error) error
}
