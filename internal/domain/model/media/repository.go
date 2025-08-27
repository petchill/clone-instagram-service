package media

import (
	"context"
	"mime/multipart"
)

type MediaRepository interface {
	UploadFileToFileStorage(ctx context.Context, file multipart.File) (string, error)
	InsertFileMetaData(ctx context.Context, mediaMetaData MediaMetaData) error
}
