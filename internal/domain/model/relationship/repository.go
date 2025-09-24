package relationship

import "context"

type RelationshipRepository interface {
	InsertFollowing(ctx context.Context, following Following) error
	DeleteFollowingByUserIDAndTargetID(ctx context.Context, userID, targetID int) error
	GetAllFollowerIDsByUserID(ctx context.Context, userID int) ([]int, error)
	GetAllFollowingIDsByUserID(ctx context.Context, userID int) ([]int, error)
	PublishFollowingTopic(ctx context.Context, message FollowingTopicMessage) error
}
