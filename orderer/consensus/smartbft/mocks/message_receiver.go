// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	smartbftprotos "github.com/hyperledger-labs/SmartBFT/smartbftprotos"
)

// MessageReceiver is an autogenerated mock type for the MessageReceiver type
type MessageReceiver struct {
	mock.Mock
}

// HandleMessage provides a mock function with given fields: sender, m
func (_m *MessageReceiver) HandleMessage(sender uint64, m *smartbftprotos.Message) {
	_m.Called(sender, m)
}

// HandleRequest provides a mock function with given fields: sender, req
func (_m *MessageReceiver) HandleRequest(sender uint64, req []byte) {
	_m.Called(sender, req)
}

type mockConstructorTestingTNewMessageReceiver interface {
	mock.TestingT
	Cleanup(func())
}

// NewMessageReceiver creates a new instance of MessageReceiver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMessageReceiver(t mockConstructorTestingTNewMessageReceiver) *MessageReceiver {
	mock := &MessageReceiver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
