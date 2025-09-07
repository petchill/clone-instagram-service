package repository

import (
	mAuth "clone-instagram-service/internal/domain/model/auth"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"
)

type authRepository struct {
	httpClient *http.Client
	oauth2     *oauth2.Config
}

func NewAuthRepository(oauth2 *oauth2.Config) *authRepository {

	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	return &authRepository{
		httpClient: &httpClient,
		oauth2:     oauth2,
	}
}

func (r *authRepository) ExchangeCodeForToken(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := r.oauth2.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (r *authRepository) GetUserInfoFromToken(ctx context.Context, accessToken string) (mAuth.UserInfo, error) {
	url, err := url.Parse("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		fmt.Println("error parsing url")
		return mAuth.UserInfo{}, err
	}

	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	request := http.Request{
		Method: http.MethodGet,
		URL:    url,
		Header: headers,
	}

	fmt.Println("request", request)

	resp, err := r.httpClient.Do(&request)
	if err != nil {
		fmt.Println("error making request")
		return mAuth.UserInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
		return mAuth.UserInfo{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return mAuth.UserInfo{}, err
	}

	userInfo := mAuth.UserInfo{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
		return mAuth.UserInfo{}, err
	}

	return userInfo, nil
}
