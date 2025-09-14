package service

import (
	mAuth "clone-instagram-service/internal/domain/model/auth"
	mUser "clone-instagram-service/internal/domain/model/user"
	"context"
	"fmt"
)

type userService struct {
	userRepo mUser.UserRepository
	authRepo mAuth.AuthRepository
}

func NewUserService(userRepo mUser.UserRepository, authRepo mAuth.AuthRepository) *userService {
	return &userService{
		userRepo: userRepo,
		authRepo: authRepo,
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
	fmt.Println("user", user)
	fmt.Println("exists", exists)
	if !exists {
		newUser := mUser.User{
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
