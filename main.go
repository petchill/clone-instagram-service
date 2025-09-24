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

	_middleware "clone-instagram-service/internal/infrastructure/middleware"

	mAuth "clone-instagram-service/internal/domain/model/auth"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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

	authConfig := mAuth.OAuthConfig{
		GoogleOAuthClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		GoogleOAuthClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		GoogleOAuthRedirectURL:  os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
	}

	oauthConfig := &oauth2.Config{
		ClientID:     authConfig.GoogleOAuthClientID,
		ClientSecret: authConfig.GoogleOAuthClientSecret,
		RedirectURL:  authConfig.GoogleOAuthRedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/bigquery",
			"https://www.googleapis.com/auth/blogger",
		},
		Endpoint: google.Endpoint,
	}

	kafkaConfig := model.KafkaConfig{
		Brokers: []string{os.Getenv("KAFKA_BROKER_ADDRESS")},
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsConfig.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsConfig.AccessKeyID, awsConfig.SecretAccessKey, "")),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// postgres
	dsn := os.Getenv("POSTGRES_CONNECTION")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to connect to database, %v", err)
	}

	// Initialize S3 client
	s3Client := s3.NewFromConfig(cfg)

	mediaRepo := _repo.NewMediaRepository(awsConfig, s3Client, db)
	mediaService := _service.NewMediaService(mediaRepo)

	relationshipRepo := _repo.NewRelationshipRepository(db, kafkaConfig)
	relationshipService := _service.NewRelationshipService(relationshipRepo)

	authRepo := _repo.NewAuthRepository(oauthConfig)

	userRepo := _repo.NewUserRepository(db)
	userService := _service.NewUserService(userRepo, authRepo)

	authMiddleWare := _middleware.NewAuthMiddleWare(authRepo, userRepo)

	e := util.InitEchoApp()

	mediaHandler := _handler.NewMediaHandler(mediaService)
	healthCheckHandler := _infra.NewHealthCheckHandler()
	authHandler := _handler.NewAuthHandler(authRepo, userService)
	relationshipHandler := _handler.NewRelationshipHandler(relationshipService)
	userHandler := _handler.NewUserHandler(userService)

	publicRoute := e.Group("/public")
	privateRoute := e.Group("/private")
	privateRoute.Use(authMiddleWare.AuthWithUser)
	mediaHandler.RegisterRoutes(privateRoute)
	relationshipHandler.RegisterRoutes(privateRoute)
	userHandler.RegisterRoutes(privateRoute)
	authHandler.RegisterRoutes(publicRoute)

	e.GET("/health", healthCheckHandler.HealthCheck)

	e.Logger.Fatal(e.Start(":5000"))
}
