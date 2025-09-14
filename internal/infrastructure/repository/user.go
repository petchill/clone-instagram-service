package repository

import (
	mUser "clone-instagram-service/internal/domain/model/user"
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

func (r *userRepository) GetUserByGoogleID(ctx context.Context, googleID string) (mUser.User, bool, error) {
	user := mUser.User{}
	err := r.gormDB.Table("user").First(&user, "google_sub_id = ?", googleID).Error
	if err != nil {
		fmt.Println("error finding user by google id:", err)
		if err == gorm.ErrRecordNotFound {
			return mUser.User{}, false, nil
		}
		log.Println("error finding user by google id:", err)
		return mUser.User{}, false, err
	}
	return user, true, nil
}

func (r *userRepository) InsertUser(ctx context.Context, user mUser.User) error {
	err := r.gormDB.Table("user").Create(&user).Error
	if err != nil {
		log.Println("error inserting user:", err)
		return err
	}
	return nil
}
