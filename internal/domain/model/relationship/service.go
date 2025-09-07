package relationship

import "context"

type RelationshipService interface {
	FollowUser(ctx context.Context, userID string, targetUserID string) error
	UnFollowUSer(ctx context.Context, userID string, targetUserID string) error
}
