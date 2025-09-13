package relationship

import "context"

type RelationshipRepository interface {
	InsertFollowing(ctx context.Context, following Following) error
	DeleteFollowingByUserIDAndTargetID(ctx context.Context, userID, targetID string) error
	GetAllFollowerIDsByUserID(ctx context.Context, userID string) ([]string, error)
	GetAllFollowingIDsByUserID(ctx context.Context, userID string) ([]string, error)
	PublishFollowingTopic(ctx context.Context, message FollowingTopicMessage) error
}
