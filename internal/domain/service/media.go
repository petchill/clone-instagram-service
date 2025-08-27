package service

import (
	mMedia "clone-instagram-service/internal/domain/model/media"
	"context"
	"mime/multipart"
	"time"
)

type mediaService struct {
	mediaRepo mMedia.MediaRepository
}

func NewMediaService() *mediaService {
	return &mediaService{}
}

// function to store media inti the DB included meta-data and file
// 1. upload to file store
// 2. get link from filestore and insert data to DB
func (s *mediaService) CreateAndStoreMedia(ctx context.Context, userID string, media multipart.File, caption string) error {
	link, err := s.mediaRepo.UploadFileToFileStorage(ctx, media)
	if err != nil {
		return err
	}

	mediaMetaData := mMedia.MediaMetaData{
		OwnerUserID:     userID,
		FileStorageLink: link,
		Caption:         caption,
		CreatedAt:       time.Now(),
	}
	// insert media meta-data to DB

	err = s.mediaRepo.InsertFileMetaData(ctx, mediaMetaData)
	if err != nil {
		return err
	}
	return nil
}
