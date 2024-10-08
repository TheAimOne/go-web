// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	model "github.com/go-web/pkg/model/user"
	mock "github.com/stretchr/testify/mock"

	pkgmodel "github.com/go-web/pkg/model"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *UserService) CreateUser(user *model.UserBase) (*model.User, error) {
	ret := _m.Called(user)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.UserBase) (*model.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*model.UserBase) *model.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.UserBase) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByMemberId provides a mock function with given fields: memberId
func (_m *UserService) GetUserByMemberId(memberId string) (*model.User, error) {
	ret := _m.Called(memberId)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(memberId)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(memberId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(memberId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: _a0
func (_m *UserService) GetUsers(_a0 model.GetUsersRequest) (*model.GetUsersResponse, error) {
	ret := _m.Called(_a0)

	var r0 *model.GetUsersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(model.GetUsersRequest) (*model.GetUsersResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(model.GetUsersRequest) *model.GetUsersResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GetUsersResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(model.GetUsersRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchUsers provides a mock function with given fields: filter
func (_m *UserService) SearchUsers(filter model.UserFilter) (*pkgmodel.PaginationResponse[*model.User], error) {
	ret := _m.Called(filter)

	var r0 *pkgmodel.PaginationResponse[*model.User]
	var r1 error
	if rf, ok := ret.Get(0).(func(model.UserFilter) (*pkgmodel.PaginationResponse[*model.User], error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(model.UserFilter) *pkgmodel.PaginationResponse[*model.User]); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pkgmodel.PaginationResponse[*model.User])
		}
	}

	if rf, ok := ret.Get(1).(func(model.UserFilter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
