package media

import (
	"context"
	"mime/multipart"
)

type MediaService interface {
	CreateAndStoreMedia(ctx context.Context, userID int, fileName string, file multipart.File, caption string) error
}
