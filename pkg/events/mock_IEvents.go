// Code generated by mockery v2.20.0. DO NOT EDIT.

package events

import mock "github.com/stretchr/testify/mock"

// MockIEvents is an autogenerated mock type for the IEvents type
type MockIEvents struct {
	mock.Mock
}

type MockIEvents_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIEvents) EXPECT() *MockIEvents_Expecter {
	return &MockIEvents_Expecter{mock: &_m.Mock}
}

// Publish provides a mock function with given fields: event, data
func (_m *MockIEvents) Publish(event string, data interface{}) {
	_m.Called(event, data)
}

// MockIEvents_Publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Publish'
type MockIEvents_Publish_Call struct {
	*mock.Call
}

// Publish is a helper method to define mock.On call
//   - event string
//   - data interface{}
func (_e *MockIEvents_Expecter) Publish(event interface{}, data interface{}) *MockIEvents_Publish_Call {
	return &MockIEvents_Publish_Call{Call: _e.mock.On("Publish", event, data)}
}

func (_c *MockIEvents_Publish_Call) Run(run func(event string, data interface{})) *MockIEvents_Publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *MockIEvents_Publish_Call) Return() *MockIEvents_Publish_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockIEvents_Publish_Call) RunAndReturn(run func(string, interface{})) *MockIEvents_Publish_Call {
	_c.Call.Return(run)
	return _c
}

// SubscribeUserRegistered provides a mock function with given fields:
func (_m *MockIEvents) SubscribeUserRegistered() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIEvents_SubscribeUserRegistered_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeUserRegistered'
type MockIEvents_SubscribeUserRegistered_Call struct {
	*mock.Call
}

// SubscribeUserRegistered is a helper method to define mock.On call
func (_e *MockIEvents_Expecter) SubscribeUserRegistered() *MockIEvents_SubscribeUserRegistered_Call {
	return &MockIEvents_SubscribeUserRegistered_Call{Call: _e.mock.On("SubscribeUserRegistered")}
}

func (_c *MockIEvents_SubscribeUserRegistered_Call) Run(run func()) *MockIEvents_SubscribeUserRegistered_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIEvents_SubscribeUserRegistered_Call) Return(_a0 error) *MockIEvents_SubscribeUserRegistered_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIEvents_SubscribeUserRegistered_Call) RunAndReturn(run func() error) *MockIEvents_SubscribeUserRegistered_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockIEvents interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockIEvents creates a new instance of MockIEvents. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockIEvents(t mockConstructorTestingTNewMockIEvents) *MockIEvents {
	mock := &MockIEvents{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}