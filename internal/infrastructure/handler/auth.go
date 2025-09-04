package handler

import (
	"clone-instagram-service/internal/domain/model"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type authHandler struct {
	oauth2     *oauth2.Config
	httpClient *http.Client
}

func NewAuthHandler(config model.OAuthConfig) *authHandler {
	conf := &oauth2.Config{
		ClientID:     config.GoogleOAuthClientID,
		ClientSecret: config.GoogleOAuthClientSecret,
		RedirectURL:  config.GoogleOAuthRedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/bigquery",
			"https://www.googleapis.com/auth/blogger",
		},
		Endpoint: google.Endpoint,
	}

	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	return &authHandler{oauth2: conf, httpClient: &httpClient}
}

func (h *authHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/auth/accessToken", h.PostAccessCode)
	e.GET("/auth/user", h.GetUser)
}

func (h *authHandler) getUserFromToken(ctx context.Context, accessToken string) (model.UserInfo, error) {
	url, err := url.Parse("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		fmt.Println("error parsing url")
		return model.UserInfo{}, err
	}

	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	request := http.Request{
		Method: http.MethodGet,
		URL:    url,
		Header: headers,
	}

	fmt.Println("request", request)

	resp, err := h.httpClient.Do(&request)
	if err != nil {
		fmt.Println("error making request")
		return model.UserInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
		return model.UserInfo{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return model.UserInfo{}, err
	}

	userInfo := model.UserInfo{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
		return model.UserInfo{}, err
	}

	return userInfo, nil
}

func (h *authHandler) GetUser(c echo.Context) error {
	fmt.Println("enter")
	ctx := c.Request().Context()
	headers := c.Request().Header
	authToken := headers.Get("Authorization")
	fmt.Println("authToken", authToken)
	if authToken == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "error": "authorization header is missing"})
	}
	token := strings.Split(authToken, " ")
	if len(token) != 2 || token[0] != "Bearer" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "error": "invalid authorization header"})
	}
	authToken = token[1]

	userInfo, err := h.getUserFromToken(ctx, authToken)
	if err != nil {
		fmt.Println("authToken", authToken)
		return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "error": err.Error()})
	}

	// TODO: check if the token is valid
	return c.JSON(http.StatusOK, userInfo)
}

func (h *authHandler) PostAccessCode(c echo.Context) error {
	ctx := c.Request().Context()
	payload := model.AccessCodePayload{}
	if err := c.Bind(&payload); err != nil {
		return err
	}
	tok, err := h.oauth2.Exchange(ctx, payload.Code)
	if err != nil {
		fmt.Printf("auth error ", err)
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"status": "failed", "error": err.Error()})
	}

	resp := model.AccessCodeResponse{
		AccessToken:  tok.AccessToken,
		ExpiresIn:    int64(time.Until(tok.Expiry).Seconds()),
		RefreshToken: tok.RefreshToken, // present if you requested offline access & user consented
		IDToken:      fmt.Sprint(tok.Extra("id_token")),
	}

	return c.JSON(http.StatusOK, resp)
}
