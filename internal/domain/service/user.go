package service

import (
	mAgg "clone-instagram-service/internal/domain/model/aggregate"
	mAuth "clone-instagram-service/internal/domain/model/auth"
	mMedia "clone-instagram-service/internal/domain/model/media"
	mUser "clone-instagram-service/internal/domain/model/user"
	eUser "clone-instagram-service/internal/domain/model/user/entity"
	"context"
	"fmt"
)

type userService struct {
	userRepo  mUser.UserRepository
	authRepo  mAuth.AuthRepository
	mediaRepo mMedia.MediaRepository
}

func NewUserService(userRepo mUser.UserRepository, authRepo mAuth.AuthRepository, mediaRepo mMedia.MediaRepository) *userService {
	return &userService{
		userRepo:  userRepo,
		authRepo:  authRepo,
		mediaRepo: mediaRepo,
	}
}

func (s *userService) LoginWithGoogleAccessCode(ctx context.Context, googleAccessCode string) (mAuth.AccessCodeResponse, error) {
	// exchange access code for id token and access token
	accessCodeResponse, err := s.authRepo.ExchangeCodeForToken(ctx, googleAccessCode)
	if err != nil {
		return mAuth.AccessCodeResponse{}, err
	}

	// get user info from id token
	userInfo, err := s.authRepo.GetUserInfoFromToken(ctx, accessCodeResponse.AccessToken)
	if err != nil {
		return mAuth.AccessCodeResponse{}, err
	}

	// check if user exists in db, if not create new user
	user, exists, err := s.userRepo.GetUserByGoogleID(ctx, userInfo.Sub)
	if err != nil {
		return mAuth.AccessCodeResponse{}, err
	}
	if !exists {
		newUser := eUser.User{
			GoogleSubID: userInfo.Sub,
			Name:        userInfo.Name,
			GivenName:   userInfo.GivenName,
			FamilyName:  userInfo.FamilyName,
			Picture:     userInfo.Picture,
			Email:       userInfo.Email,
		}
		err = s.userRepo.InsertUser(ctx, newUser)
		if err != nil {
			return mAuth.AccessCodeResponse{}, err
		}
		user = newUser
	}

	// return access token and user info
	return mAuth.AccessCodeResponse{
		AccessToken: accessCodeResponse.AccessToken,
		ExpiresIn:   3600, // 1 hour
		UserInfo:    user,
	}, nil
}

func (s *userService) GetUserProfileByGoogleSubID(ctx context.Context, googleSubID string) (mAgg.UserProfile, error) {
	resp := mAgg.UserProfile{}
	// get user from DB by googleSubID
	user, exist, err := s.userRepo.GetUserByGoogleID(ctx, googleSubID)
	if err != nil {
		return resp, err
	}
	if !exist {
		err = fmt.Errorf("user not found")
		return resp, err
	}

	// get followers - all followers should have user info
	followingUsers, err := s.userRepo.GetFollowingUsersByUserID(ctx, user.ID)
	if err != nil {
		return resp, err
	}
	// get followings - all followers should have user info
	followerUsers, err := s.userRepo.GetFollowerUsersByUserID(ctx, user.ID)
	if err != nil {
		return resp, err
	}
	// get posts
	posts, err := s.mediaRepo.GetMediasByOwnerUserID(ctx, user.ID)
	if err != nil {
		return resp, err
	}

	resp = mAgg.UserProfile{
		User:       user,
		Followers:  followerUsers,
		Followings: followingUsers,
		Posts:      posts,
	}

	return resp, nil
}
