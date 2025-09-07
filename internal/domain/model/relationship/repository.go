package relationship

import "context"

type RelationshipRepository interface {
	InsertFollowing(ctx context.Context, following Following) error
	DeleteFollowingByUserIDAndTargetID(ctx context.Context, userID, targetID string) error
}
