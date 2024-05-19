package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-web/pkg/constants"
	authModel "github.com/go-web/pkg/model/auth"
	userModel "github.com/go-web/pkg/model/user"
	"github.com/go-web/pkg/repository"
	"github.com/go-web/pkg/util"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

const (
	ExpiryTimeInMinutes             = 10
	DefaultJwtExpiryTimeInSeconds   = "300"
	DefaultRefreshTokenExpiryInDays = "30"
	DefaultAccessTokenSigningKey    = "eJZ6GgukXTrBsaE"
	DefaultRefreshTokenSigningKey   = "eJZ6GgukXTrBsaE"
	JWT_SECRET                      = "JWT_SECRET"
	REFRESH_TOKEN_SECRET            = "REFRESH_TOKEN_SECRET"
	JWT_EXPIRY_TIME_IN_SECONDS      = "JWT_EXPIRY_TIME_IN_SECONDS"
	REFRESH_TOKEN_EXPIRY_IN_DAYS    = "REFRESH_TOKEN_EXPIRY_IN_DAYS"
)

type authImpl struct {
	userRepository repository.UserRepository
}

func (a *authImpl) Authenticate(authRequest *authModel.AuthRequest) (*authModel.AuthResponse, error) {
	fmt.Println("GEtting hit auth.....")
	userId := authRequest.UserId
	password := authRequest.Password
	authType := authRequest.Type
	usr, err := loginByType(a, authType, userId, password)

	if err != nil {
		return nil, constants.ErrorAuthenticating
	}
	token, accessTokenExpiry, err := a.CreateJwtAuthToken(userId)

	var refreshToken string
	var refreshTokenError error
	var expiryTime time.Time
	if authRequest.DeviceId == "" {
		session := &authModel.Session{
			DeviceId:   uuid.NewV4().String(), // change this
			DeviceType: authRequest.DeviceType,
			SessionId:  uuid.NewV4().String(),
		}
		refreshToken, expiryTime, refreshTokenError = a.CreateRefreshToken(session)
		authRequest.DeviceId = session.DeviceId
	}

	if err != nil || refreshTokenError != nil {
		return nil, constants.ErrorCreatingToken
	}

	resp := &authModel.AuthResponse{
		Data: &authModel.AuthData{
			Token:              token,
			TokenExpiry:        accessTokenExpiry,
			RefreshToken:       refreshToken,
			DeviceId:           authRequest.DeviceId,
			DeviceType:         authRequest.DeviceType,
			RefreshTokenExpiry: expiryTime,
			User:               usr,
		},
	}

	return resp, nil
}

func (a *authImpl) CreateJwtAuthToken(userId string) (string, time.Time, error) {
	if userId == "" {
		return "", time.Time{}, constants.ErrorCreatingToken
	}

	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte(util.GetEnvOrDefault(JWT_SECRET, DefaultAccessTokenSigningKey))
	expiryTimeInSeconds, _ := strconv.Atoi(util.GetEnvOrDefault(JWT_EXPIRY_TIME_IN_SECONDS, DefaultJwtExpiryTimeInSeconds))
	accessTokenExpiryInSeconds := time.Now().UTC().Add(time.Second * time.Duration(expiryTimeInSeconds))
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userId,
		"iat":  time.Now().UTC().Unix(),
		"exp":  accessTokenExpiryInSeconds.Unix(),
	})
	s, err := t.SignedString(key)

	if err != nil {
		return "", time.Time{}, constants.ErrorCreatingToken
	}

	return s, accessTokenExpiryInSeconds, nil
}

func (a *authImpl) CreateRefreshToken(session *authModel.Session) (string, time.Time, error) {
	if session.DeviceId == "" || session.DeviceType == "" || session.SessionId == "" {
		return "", time.Time{}, constants.ErrorCreatingToken
	}

	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte(util.GetEnvOrDefault(REFRESH_TOKEN_SECRET, DefaultRefreshTokenSigningKey))
	refreshTokenExpiryInDays, _ := strconv.Atoi(util.GetEnvOrDefault(REFRESH_TOKEN_EXPIRY_IN_DAYS, DefaultRefreshTokenExpiryInDays))
	refreshTokenExpiryInSeconds := refreshTokenExpiryInDays * 24 * 60 * 60
	refreshTokenExpiryTime := time.Now().UTC().Add(time.Second * time.Duration(refreshTokenExpiryInSeconds))

	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"deviceId":   session.DeviceId,
		"deviceType": session.DeviceType,
		"sessionId":  session.SessionId,
		"iat":        time.Now().UTC().Unix(),
		"exp":        refreshTokenExpiryTime.Unix(),
	})
	s, err := t.SignedString(key)

	if err != nil {
		return "", time.Time{}, constants.ErrorCreatingToken
	}

	return s, refreshTokenExpiryTime, nil
}

func loginByType(a *authImpl, authType string, userId string, password string) (*userModel.User, error) {
	switch authType {
	case authModel.AUTH_TYPE_EMAIL:
		return a.userRepository.AuthenticateUserByEmail(userId, password)
	case authModel.AUTH_TYPE_MOBILE:
		return a.userRepository.AuthenticateUserByMobile(userId, password)
	default:
		return nil, constants.ErrorAuthTypeNotProvided
	}
}

func (a *authImpl) ParseAccessToken(accessToken string) *jwt.Token {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(util.GetEnvOrDefault(JWT_SECRET, DefaultAccessTokenSigningKey)), nil
	})
	return parsedAccessToken
}

func (a *authImpl) ParseRefreshToken(refreshToken string) *jwt.Token {
	parsedAccessToken, _ := jwt.ParseWithClaims(refreshToken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(util.GetEnvOrDefault(JWT_SECRET, DefaultAccessTokenSigningKey)), nil
	})
	return parsedAccessToken
}

func (a *authImpl) CheckJwtAccessTokenValidity(token string) bool {
	token = strings.TrimPrefix(token, "Bearer ")

	parsedToken := a.ParseAccessToken(token)
	if parsedToken != nil {
		return parsedToken.Valid
	}
	return false
}

func (a *authImpl) CheckJwtRefreshTokenValidity(token string) bool {
	parsedToken := a.ParseRefreshToken(token)
	if parsedToken != nil {
		return parsedToken.Valid
	}
	return false
}

func (a *authImpl) GenerateToken(authRequest authModel.AuthRequest) (*authModel.AuthData, error) {
	if authRequest.RefreshToken == "" {
		return nil, constants.ErrorRefreshTokenNotProvided
	}

	parsedRefreshToken := a.ParseRefreshToken(authRequest.RefreshToken)
	if !parsedRefreshToken.Valid {
		return nil, constants.ErrorRefreshTokenExpired
	}

	if authRequest.DeviceId == "" || authRequest.UserId == "" {
		return nil, constants.ErrorInvalidDetails
	}

	accessToken, tokenExpiry, err := a.CreateJwtAuthToken(authRequest.UserId)
	if err != nil {
		return nil, constants.ErrorCreatingToken
	}
	return &authModel.AuthData{
		Token:        accessToken,
		TokenExpiry:  tokenExpiry,
		RefreshToken: authRequest.RefreshToken,
		DeviceId:     authRequest.DeviceId,
		DeviceType:   authRequest.DeviceType,
	}, nil
}
