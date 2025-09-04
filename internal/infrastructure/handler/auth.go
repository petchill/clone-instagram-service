package handler

import (
	"clone-instagram-service/internal/domain/model"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type authHandler struct {
	config model.OAuthConfig
	oauth2 *oauth2.Config
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
	return &authHandler{oauth2: conf}
}

func (h *authHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/auth/accessToken", h.PostAccessCode)
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
