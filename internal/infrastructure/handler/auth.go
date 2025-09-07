package handler

import (
	mAuth "clone-instagram-service/internal/domain/model/auth"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	authRepo mAuth.AuthRepository
}

func NewAuthHandler(authRepo mAuth.AuthRepository) *authHandler {

	return &authHandler{authRepo: authRepo}
}

func (h *authHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/auth/accessToken", h.PostAccessCode)
	e.GET("/auth/user", h.GetUser)
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

	userInfo, err := h.authRepo.GetUserInfoFromToken(ctx, authToken)
	if err != nil {
		fmt.Println("authToken", authToken)
		return c.JSON(http.StatusUnauthorized, map[string]string{"status": "failed", "error": err.Error()})
	}

	// TODO: check if the token is valid
	return c.JSON(http.StatusOK, userInfo)
}

func (h *authHandler) PostAccessCode(c echo.Context) error {
	ctx := c.Request().Context()
	payload := mAuth.AccessCodePayload{}
	if err := c.Bind(&payload); err != nil {
		return err
	}
	tok, err := h.authRepo.ExchangeCodeForToken(ctx, payload.Code)
	if err != nil {
		fmt.Printf("auth error ", err)
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"status": "failed", "error": err.Error()})
	}

	resp := mAuth.AccessCodeResponse{
		AccessToken:  tok.AccessToken,
		ExpiresIn:    int64(time.Until(tok.Expiry).Seconds()),
		RefreshToken: tok.RefreshToken, // present if you requested offline access & user consented
		IDToken:      fmt.Sprint(tok.Extra("id_token")),
	}

	return c.JSON(http.StatusOK, resp)
}
