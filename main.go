package main

import (
	"clone-instagram-service/internal/domain/model"
	_service "clone-instagram-service/internal/domain/service"
	_infra "clone-instagram-service/internal/infrastructure"
	_handler "clone-instagram-service/internal/infrastructure/handler"
	_repo "clone-instagram-service/internal/infrastructure/repository"
	"clone-instagram-service/internal/util"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	awsConfig := model.AWSConfig{
		Region:              os.Getenv("AWS_REGION"),
		AccessKeyID:         os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey:     os.Getenv("AWS_SECRET_ACCESS_KEY"),
		BucketName:          os.Getenv("AWS_BUCKET_NAME"),
		PublicBucketBaseURL: os.Getenv("AWS_PUBLIC_BUCKET_BASE_URL"),
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsConfig.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsConfig.AccessKeyID, awsConfig.SecretAccessKey, "")),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Initialize S3 client
	s3Client := s3.NewFromConfig(cfg)

	mediaRepo := _repo.NewMediaRepository(awsConfig, s3Client)
	mediaService := _service.NewMediaService(mediaRepo)

	e := util.InitEchoApp()

	mediaHandler := _handler.NewMediaHandler(mediaService)
	healthCheckHandler := _infra.NewHealthCheckHandler()

	mediaHandler.RegisterRoutes(e)

	e.GET("/health", healthCheckHandler.HealthCheck)

	e.Logger.Fatal(e.Start(":5000"))
}
