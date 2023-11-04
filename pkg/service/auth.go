package service

import (
	"fmt"
	"time"

	"github.com/go-web/middleware"
	"github.com/go-web/pkg/constants"
	authModel "github.com/go-web/pkg/model/auth"
	"github.com/go-web/pkg/repository"
)

type authImpl struct {
	userRepository repository.UserRepository
}

func (a *authImpl) Authenticate(authRequest *authModel.AuthRequest) (*authModel.AuthResponse, error) {
	fmt.Println("GEtting hit auth.....")
	userId := authRequest.UserId
	password := authRequest.Password
	authType := authRequest.Type

	switch authType {
	case authModel.AUTH_TYPE_EMAIL:
		user, err := a.userRepository.AuthenticateUserByEmail(userId, password)
		if err != nil || user.Email == "" {
			return nil, constants.ErrorAuthenticating
		}
	case authModel.AUTH_TYPE_MOBILE:
		user, err := a.userRepository.AuthenticateUserByMobile(userId, password)
		if err != nil || user.Mobile == "" {
			return nil, constants.ErrorAuthenticating
		}
	default:
		return nil, constants.ErrorAuthTypeNotProvided
	}

	auth := middleware.Auth{
		UserId:    userId,
		CreatedAt: time.Now().Unix(),
	}

	token, err := middleware.CreateAuthToken(&auth)

	if err != nil {
		return nil, constants.ErrorCreatingToken
	}

	resp := &authModel.AuthResponse{
		Data: &authModel.AuthData{
			Token: token,
		},
	}

	return resp, nil
}
