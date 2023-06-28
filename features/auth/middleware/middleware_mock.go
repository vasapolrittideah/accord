// Code generated by mockery v2.30.1. DO NOT EDIT.

package middleware

import (
	fiber "github.com/gofiber/fiber/v2"
	config "github.com/vasapolrittideah/accord/internal/config"

	mock "github.com/stretchr/testify/mock"
)

// MockAuthMiddleware is an autogenerated mock type for the AuthMiddleware type
type MockAuthMiddleware struct {
	mock.Mock
}

type MockAuthMiddleware_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAuthMiddleware) EXPECT() *MockAuthMiddleware_Expecter {
	return &MockAuthMiddleware_Expecter{mock: &_m.Mock}
}

// AuthenticateWithJWT provides a mock function with given fields: conf, tokenType
func (_m *MockAuthMiddleware) AuthenticateWithJWT(conf *config.Config, tokenType TokenType) func(*fiber.Ctx) error {
	ret := _m.Called(conf, tokenType)

	var r0 func(*fiber.Ctx) error
	if rf, ok := ret.Get(0).(func(*config.Config, TokenType) func(*fiber.Ctx) error); ok {
		r0 = rf(conf, tokenType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(*fiber.Ctx) error)
		}
	}

	return r0
}

// MockAuthMiddleware_AuthenticateWithJWT_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthenticateWithJWT'
type MockAuthMiddleware_AuthenticateWithJWT_Call struct {
	*mock.Call
}

// AuthenticateWithJWT is a helper method to define mock.On call
//   - conf *config.Config
//   - tokenType TokenType
func (_e *MockAuthMiddleware_Expecter) AuthenticateWithJWT(conf interface{}, tokenType interface{}) *MockAuthMiddleware_AuthenticateWithJWT_Call {
	return &MockAuthMiddleware_AuthenticateWithJWT_Call{Call: _e.mock.On("AuthenticateWithJWT", conf, tokenType)}
}

func (_c *MockAuthMiddleware_AuthenticateWithJWT_Call) Run(run func(conf *config.Config, tokenType TokenType)) *MockAuthMiddleware_AuthenticateWithJWT_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*config.Config), args[1].(TokenType))
	})
	return _c
}

func (_c *MockAuthMiddleware_AuthenticateWithJWT_Call) Return(_a0 func(*fiber.Ctx) error) *MockAuthMiddleware_AuthenticateWithJWT_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAuthMiddleware_AuthenticateWithJWT_Call) RunAndReturn(run func(*config.Config, TokenType) func(*fiber.Ctx) error) *MockAuthMiddleware_AuthenticateWithJWT_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAuthMiddleware creates a new instance of MockAuthMiddleware. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAuthMiddleware(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAuthMiddleware {
	mock := &MockAuthMiddleware{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
