package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/go-web/database/function"
	"github.com/go-web/pkg/constants"
	model "github.com/go-web/pkg/model/user"
)

const userTableName = `"user"`

var userTableColumns = []string{"member_id", "name", "short_name", "email", "mobile", "status"}
var userSessionTableColumns = []string{"session_id", "user_id", "session_token", "device_id", "device_type", "expiry_time"}

var UserFilterMap = map[string]string{
	"memberId":  "member_id",
	"name":      "name",
	"shortName": "short_name",
	"email":     "email",
	"mobile":    "mobile",
	"status":    "status",
}

type UserRepository interface {
	CreateUser(user model.User) error
	GetUserByMemberId(memberId string) (*model.User, error)
	GetUsers(page, perPage int) ([]*model.User, error)
	AuthenticateUserByEmail(email, password string) (*model.User, error)
	AuthenticateUserByMobile(mobile, password string) (*model.User, error)
	SearchUsers(filter model.UserFilter) ([]*model.User, error)
	GetSessionByUserIdAndDeviceId(userId string, deviceId string) ([]*model.UserSession, error)
	CreateUserDeviceSession(userSession *model.UserSession) error
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

func (u *UserRepoImpl) GetUsers(page, perPage int) ([]*model.User, error) {
	limit := perPage
	offset := (page - 1) * perPage

	condition := fmt.Sprintf("offset %d limit %d", offset, limit)

	rows, err := u.DB.SelectAll(userTableName, condition, userTableColumns)
	if err != nil || rows == nil {
		return nil, constants.ErrorReadingFromDB
	}

	return getUserListFromRows(rows), nil
}

func getUserListFromRows(rows *sql.Rows) []*model.User {
	result := make([]*model.User, 0)
	for rows.Next() {
		var userResult model.User
		rows.Scan(&userResult.MemberId, &userResult.Name, &userResult.ShortName, &userResult.Email, &userResult.Mobile, &userResult.Status)

		result = append(result, &userResult)
	}
	return result
}

func (u *UserRepoImpl) AuthenticateUserByMobile(mobile, password string) (*model.User, error) {
	row, err := u.DB.Select(userTableName, fmt.Sprintf(" where mobile = '%s' and password = '%s'", mobile, password), userTableColumns)
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

func (u *UserRepoImpl) AuthenticateUserByEmail(email, password string) (*model.User, error) {
	row, err := u.DB.Select(userTableName, fmt.Sprintf(" where email = '%s' and password = '%s'", email, password), userTableColumns)
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

func (u *UserRepoImpl) SearchUsers(filter model.UserFilter) ([]*model.User, error) {

	var rows *sql.Rows
	var err error
	if !filter.ExcludeUserByGroupId {
		rows, err = u.DB.SelectPaginateAndFilter(userTableName, filter.Filter, userTableColumns, UserFilterMap)
	} else {
		query := fmt.Sprintf(`select %s from %s u WHERE u.member_id not in 
			(select gm.member_id from group_member gm where gm.group_id = '%s')`,
			strings.Join(userTableColumns, ", "), userTableName, filter.GroupId)
		rows, err = u.DB.SelectPaginateAndFilterByQuery(query, filter.Filter, UserFilterMap)
	}

	if err != nil {
		log.Println(err)
		return nil, constants.ErrorReadingFromDB
	}

	return getUserListFromRows(rows), nil
}

func (u *UserRepoImpl) GetSessionByUserIdAndDeviceId(userId string, deviceId string) ([]*model.UserSession, error) {
	query := fmt.Sprintf(`select %s from user_session us WHERE us.userId = %s and us.deviceId = %s `,
		strings.Join(userSessionTableColumns, ","), userId, deviceId)
	rows, err := u.DB.SelectRaw(query)

	if err != nil {
		log.Println(err)
		return nil, constants.ErrorReadingFromDB
	}
	result := make([]*model.UserSession, 0)

	for rows.Next() {
		var userSessionResult model.UserSession
		rows.Scan(&userSessionResult.SessionId, &userSessionResult.UserId, &userSessionResult.SessionToken,
			&userSessionResult.DeviceId, &userSessionResult.DeviceType)

		result = append(result, &userSessionResult)
	}
	return result, nil
}

func (u *UserRepoImpl) CreateUserDeviceSession(userSession *model.UserSession) error {
	values := []interface{}{
		userSession.SessionId,
		userSession.UserId,
		userSession.SessionToken,
		userSession.DeviceId,
		userSession.DeviceType,
		userSession.ExpiryTime,
	}

	err := u.DB.Insert("user_session", columns, values)

	return err
}
