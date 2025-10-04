package relationship

import (
	eRela "clone-instagram-service/internal/domain/model/relationship/entity"
	"context"
)

type RelationshipRepository interface {
	InsertFollowing(ctx context.Context, following eRela.Following) error
	DeleteFollowingByUserIDAndTargetID(ctx context.Context, userID, targetID int) error
	GetAllFollowerIDsByUserID(ctx context.Context, userID int) ([]int, error)
	GetAllFollowingIDsByUserID(ctx context.Context, userID int) ([]int, error)
	PublishFollowingTopic(ctx context.Context, message eRela.FollowingTopicMessage) error
}
