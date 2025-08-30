package repository

import (
	"clone-instagram-service/internal/domain/model"
	"context"
	"errors"
	"log"
	"mime/multipart"
	"time"

	mMedia "clone-instagram-service/internal/domain/model/media"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
)

type mediaRepository struct {
	awsConfig model.AWSConfig
	s3Client  *s3.Client
}

func NewMediaRepository(awsConfig model.AWSConfig, s3Client *s3.Client) *mediaRepository {
	return &mediaRepository{
		awsConfig: awsConfig,
		s3Client:  s3Client,
	}
}

func (r *mediaRepository) UploadFileToFileStorage(ctx context.Context, objectKey string, file multipart.File) (string, error) {
	_, err := r.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(r.awsConfig.BucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode() == "EntityTooLarge" {
			log.Printf("Error while uploading object to %s. The object is too large.\n"+
				"To upload objects larger than 5GB, use the S3 console (160GB max)\n"+
				"or the multipart upload API (5TB max).", r.awsConfig.BucketName)
		} else {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				r.awsConfig.BucketName, objectKey, err)
		}
		return "", err
	} else {
		err = s3.NewObjectExistsWaiter(r.s3Client).Wait(
			ctx, &s3.HeadObjectInput{Bucket: aws.String(r.awsConfig.BucketName), Key: aws.String(objectKey)}, time.Minute)
		if err != nil {
			log.Printf("Failed attempt to wait for object %s to exist.\n", objectKey)
			return "", err
		}
	}

	publicObjectLink := r.awsConfig.PublicBucketBaseURL + "/" + objectKey

	return publicObjectLink, nil

}

func (r *mediaRepository) InsertFileMetaData(ctx context.Context, mediaMetaData mMedia.MediaMetaData) error {
	return nil
}
