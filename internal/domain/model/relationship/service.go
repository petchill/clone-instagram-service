package relationship

import "context"

type RelationshipService interface {
	FollowUser(ctx context.Context, userID string, targetUserID string) error
	// UnFollowUser(ctx context.Context, userID string, targetUserID string) error
}
