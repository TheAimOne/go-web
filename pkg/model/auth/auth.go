package auth

import (
	"time"

	userModel "github.com/go-web/pkg/model/user"
)

const (
	AUTH_TYPE_MOBILE = "mobile"
	AUTH_TYPE_EMAIL  = "email"
)

type AuthRequest struct {
	UserId       string `json:"userId"`
	Password     string `json:"password"`
	Type         string `json:"type"`
	DeviceId     string `json:"deviceId"`
	DeviceType   string `json:"deviceType"`
	RefreshToken string `json:"refreshToken"`
}

type AuthResponse struct {
	Data *AuthData `json:"data"`
}

type AuthData struct {
	Token              string          `json:"token"`
	TokenExpiry        time.Time       `json:"tokenExpiry"`
	RefreshToken       string          `json:"refreshToken"`
	DeviceId           string          `json:"deviceId"`
	DeviceType         string          `json:"deviceType"`
	RefreshTokenExpiry time.Time       `json:"refreshTokenExpiry"`
	User               *userModel.User `json:"user"`
}

type Auth struct {
	UserId    string
	CreatedAt int64
}

type Session struct {
	DeviceId     string
	DeviceType   string
	SessionId    string
	SessionToken string
}
