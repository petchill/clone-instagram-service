package media

import (
	"context"
	"mime/multipart"
)

type MediaRepository interface {
	UploadFileToFileStorage(ctx context.Context, objectKey string, file multipart.File) (string, error)
	InsertFileMetaData(ctx context.Context, mediaMetaData MediaMetaData) error
}
