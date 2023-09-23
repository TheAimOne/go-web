package service

import (
	"errors"
	"log"

	"github.com/go-web/pkg/constants"
	model "github.com/go-web/pkg/model/user"
	"github.com/go-web/pkg/repository"
	uuid "github.com/satori/go.uuid"
)

type UserImpl struct {
	userRepository repository.UserRepository
}

func (u *UserImpl) CreateUser(user *model.UserBase) (*model.User, error) {
	if user == nil {
		return nil, errors.New("user object is empty")
	}
	if user.MemberId == uuid.Nil {
		user.MemberId = uuid.NewV4()
	}

	userModel := &model.User{
		UserBase: *user,
		Status:   model.ACTIVE,
	}

	err := u.userRepository.CreateUser(*userModel)
	if err != nil {
		log.Println(err)
		return nil, constants.ErrorCreatingUser
	}

	return userModel, nil
}

func (u *UserImpl) GetUserByMemberId(memberId string) (*model.User, error) {
	if memberId == "" {
		return nil, errors.New("userId is empty")
	}

	user, err := u.userRepository.GetUserByMemberId(memberId)
	if err != nil {
		return nil, constants.ErrorGettingUser
	}

	return user, nil
}
