package repository

import (
	eUser "clone-instagram-service/internal/domain/model/user/entity"
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type userRepository struct {
	gormDB *gorm.DB
}

func NewUserRepository(gormDB *gorm.DB) *userRepository {
	return &userRepository{gormDB: gormDB}
}

func (r *userRepository) GetUserByGoogleID(ctx context.Context, googleID string) (eUser.User, bool, error) {
	user := eUser.User{}
	err := r.gormDB.Table("users").First(&user, "google_sub_id = ?", googleID).Error
	if err != nil {
		fmt.Println("error finding user by google id:", err)
		if err == gorm.ErrRecordNotFound {
			return eUser.User{}, false, nil
		}
		log.Println("error finding user by google id:", err)
		return eUser.User{}, false, err
	}
	return user, true, nil
}

func (r *userRepository) InsertUser(ctx context.Context, user eUser.User) error {
	err := r.gormDB.Table("users").Create(&user).Error
	if err != nil {
		log.Println("error inserting user:", err)
		return err
	}
	return nil
}

func (r *userRepository) GetFollowingUsersByUserID(ctx context.Context, userID int) ([]eUser.User, error) {
	followingUsers := []eUser.User{}
	err := r.gormDB.
		Table("users").
		Joins(`JOIN followings ON users.id = followings.target_user_id`).
		Where("followings.user_id = ?", userID).
		Find(&followingUsers).Error
	if err != nil {
		log.Printf("Error while getting following users from database. Here's why: %v\n", err)
		return followingUsers, err
	}
	return followingUsers, nil
}

func (r *userRepository) GetFollowerUsersByUserID(ctx context.Context, userID int) ([]eUser.User, error) {
	followerUsers := []eUser.User{}
	err := r.gormDB.
		Table("users").
		Joins(`JOIN followings ON users.id = followings.user_id`).
		Where("followings.target_user_id = ?", userID).
		Find(&followerUsers).Error
	if err != nil {
		log.Printf("Error while getting follower users from database. Here's why: %v\n", err)
		return followerUsers, err
	}
	return followerUsers, nil
}
