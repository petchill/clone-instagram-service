package media

import (
	"context"
	"mime/multipart"
)

type MediaService interface {
	CreateAndStoreMedia(ctx context.Context, userID string, fileName string, file multipart.File, caption string) error
}
