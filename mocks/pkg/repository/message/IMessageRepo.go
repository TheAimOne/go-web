// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	message "github.com/go-web/pkg/model/message"
	mock "github.com/stretchr/testify/mock"
)

// IMessageRepo is an autogenerated mock type for the IMessageRepo type
type IMessageRepo struct {
	mock.Mock
}

// CreateMessage provides a mock function with given fields: m
func (_m *IMessageRepo) CreateMessage(m *message.Message) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(*message.Message) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveMessageForEvent provides a mock function with given fields: eventId, offset
func (_m *IMessageRepo) RetrieveMessageForEvent(eventId string, offset int) ([]*message.Message, error) {
	ret := _m.Called(eventId, offset)

	var r0 []*message.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) ([]*message.Message, error)); ok {
		return rf(eventId, offset)
	}
	if rf, ok := ret.Get(0).(func(string, int) []*message.Message); ok {
		r0 = rf(eventId, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*message.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(eventId, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIMessageRepo creates a new instance of IMessageRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMessageRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMessageRepo {
	mock := &IMessageRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
