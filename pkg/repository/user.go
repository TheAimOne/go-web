package repository

import (
	"fmt"
	"log"

	"github.com/go-web/database/function"
	"github.com/go-web/pkg/constants"
	model "github.com/go-web/pkg/model/user"
)

const userTableName = `"user"`

var userTableColumns = []string{"member_id", "name", "short_name", "email", "mobile", "status"}

type UserRepository interface {
	CreateUser(user model.User) error
	GetUserByMemberId(memberId string) (*model.User, error)
}

func NewMemberRepository(DB function.DBFunction) UserRepository {
	return &UserRepoImpl{
		DB: DB,
	}
}

type UserRepoImpl struct {
	DB function.DBFunction
}

func (u *UserRepoImpl) CreateUser(user model.User) error {

	values := []interface{}{
		&user.MemberId,
		&user.Name,
		&user.ShortName,
		&user.Email,
		&user.Mobile,
		&user.Status,
	}

	err := u.DB.Insert(userTableName, userTableColumns, values)
	return err
}

func (u *UserRepoImpl) GetUserByMemberId(memberId string) (*model.User, error) {
	row, err := u.DB.Select(userTableName, fmt.Sprintf(" where member_id = '%s'", memberId), userTableColumns)
	if err != nil {
		log.Println(err)
		return nil, constants.ErrorReadingFromDB
	}
	if row == nil {
		return nil, constants.ErrorReadingFromDB
	}

	var userResult model.User
	row.Scan(&userResult.MemberId, &userResult.Name, &userResult.ShortName, &userResult.Email, &userResult.Mobile, &userResult.Status)
	return &userResult, nil
}
