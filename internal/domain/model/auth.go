package model

type OAuthConfig struct {
	GoogleOAuthClientID     string
	GoogleOAuthClientSecret string
	GoogleOAuthRedirectURL  string
}

type AccessCodePayload struct {
	Code string `json:"code"`
}

type AccessCodeResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	IDToken      string `json:"id_token,omitempty"`
	UserInfo     any    `json:"userinfo,omitempty"`
}
