package media

import (
	"context"
	"mime/multipart"
)

type MediaService interface {
	CreateAndStoreMedia(ctx context.Context, userID string, media multipart.File, caption string) error
}
