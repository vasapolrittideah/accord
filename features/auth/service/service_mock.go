// Code generated by mockery v2.30.1. DO NOT EDIT.

package service

import (
	jwt "github.com/golang-jwt/jwt/v4"
	mock "github.com/stretchr/testify/mock"

	models "github.com/vasapolrittideah/accord/models"

	time "time"

	uuid "github.com/google/uuid"
)

// MockAuthService is an autogenerated mock type for the AuthService type
type MockAuthService struct {
	mock.Mock
}

type MockAuthService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAuthService) EXPECT() *MockAuthService_Expecter {
	return &MockAuthService_Expecter{mock: &_m.Mock}
}

// GenerateToken provides a mock function with given fields: ttl, privateKey, userId
func (_m *MockAuthService) GenerateToken(ttl time.Duration, privateKey string, userId uuid.UUID) (string, error) {
	ret := _m.Called(ttl, privateKey, userId)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Duration, string, uuid.UUID) (string, error)); ok {
		return rf(ttl, privateKey, userId)
	}
	if rf, ok := ret.Get(0).(func(time.Duration, string, uuid.UUID) string); ok {
		r0 = rf(ttl, privateKey, userId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(time.Duration, string, uuid.UUID) error); ok {
		r1 = rf(ttl, privateKey, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthService_GenerateToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateToken'
type MockAuthService_GenerateToken_Call struct {
	*mock.Call
}

// GenerateToken is a helper method to define mock.On call
//   - ttl time.Duration
//   - privateKey string
//   - userId uuid.UUID
func (_e *MockAuthService_Expecter) GenerateToken(ttl interface{}, privateKey interface{}, userId interface{}) *MockAuthService_GenerateToken_Call {
	return &MockAuthService_GenerateToken_Call{Call: _e.mock.On("GenerateToken", ttl, privateKey, userId)}
}

func (_c *MockAuthService_GenerateToken_Call) Run(run func(ttl time.Duration, privateKey string, userId uuid.UUID)) *MockAuthService_GenerateToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration), args[1].(string), args[2].(uuid.UUID))
	})
	return _c
}

func (_c *MockAuthService_GenerateToken_Call) Return(_a0 string, _a1 error) *MockAuthService_GenerateToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthService_GenerateToken_Call) RunAndReturn(run func(time.Duration, string, uuid.UUID) (string, error)) *MockAuthService_GenerateToken_Call {
	_c.Call.Return(run)
	return _c
}

// HashRefreshToken provides a mock function with given fields: refreshToken
func (_m *MockAuthService) HashRefreshToken(refreshToken string) (string, error) {
	ret := _m.Called(refreshToken)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(refreshToken)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(refreshToken)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthService_HashRefreshToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashRefreshToken'
type MockAuthService_HashRefreshToken_Call struct {
	*mock.Call
}

// HashRefreshToken is a helper method to define mock.On call
//   - refreshToken string
func (_e *MockAuthService_Expecter) HashRefreshToken(refreshToken interface{}) *MockAuthService_HashRefreshToken_Call {
	return &MockAuthService_HashRefreshToken_Call{Call: _e.mock.On("HashRefreshToken", refreshToken)}
}

func (_c *MockAuthService_HashRefreshToken_Call) Run(run func(refreshToken string)) *MockAuthService_HashRefreshToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAuthService_HashRefreshToken_Call) Return(_a0 string, _a1 error) *MockAuthService_HashRefreshToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthService_HashRefreshToken_Call) RunAndReturn(run func(string) (string, error)) *MockAuthService_HashRefreshToken_Call {
	_c.Call.Return(run)
	return _c
}

// RefreshToken provides a mock function with given fields: userId, userRefreshToken
func (_m *MockAuthService) RefreshToken(userId uuid.UUID, userRefreshToken string) (*Tokens, error) {
	ret := _m.Called(userId, userRefreshToken)

	var r0 *Tokens
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, string) (*Tokens, error)); ok {
		return rf(userId, userRefreshToken)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, string) *Tokens); ok {
		r0 = rf(userId, userRefreshToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Tokens)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, string) error); ok {
		r1 = rf(userId, userRefreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthService_RefreshToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefreshToken'
type MockAuthService_RefreshToken_Call struct {
	*mock.Call
}

// RefreshToken is a helper method to define mock.On call
//   - userId uuid.UUID
//   - userRefreshToken string
func (_e *MockAuthService_Expecter) RefreshToken(userId interface{}, userRefreshToken interface{}) *MockAuthService_RefreshToken_Call {
	return &MockAuthService_RefreshToken_Call{Call: _e.mock.On("RefreshToken", userId, userRefreshToken)}
}

func (_c *MockAuthService_RefreshToken_Call) Run(run func(userId uuid.UUID, userRefreshToken string)) *MockAuthService_RefreshToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(string))
	})
	return _c
}

func (_c *MockAuthService_RefreshToken_Call) Return(_a0 *Tokens, _a1 error) *MockAuthService_RefreshToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthService_RefreshToken_Call) RunAndReturn(run func(uuid.UUID, string) (*Tokens, error)) *MockAuthService_RefreshToken_Call {
	_c.Call.Return(run)
	return _c
}

// SignIn provides a mock function with given fields: payload
func (_m *MockAuthService) SignIn(payload SignInRequest) (*Tokens, error) {
	ret := _m.Called(payload)

	var r0 *Tokens
	var r1 error
	if rf, ok := ret.Get(0).(func(SignInRequest) (*Tokens, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(SignInRequest) *Tokens); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Tokens)
		}
	}

	if rf, ok := ret.Get(1).(func(SignInRequest) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthService_SignIn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignIn'
type MockAuthService_SignIn_Call struct {
	*mock.Call
}

// SignIn is a helper method to define mock.On call
//   - payload SignInRequest
func (_e *MockAuthService_Expecter) SignIn(payload interface{}) *MockAuthService_SignIn_Call {
	return &MockAuthService_SignIn_Call{Call: _e.mock.On("SignIn", payload)}
}

func (_c *MockAuthService_SignIn_Call) Run(run func(payload SignInRequest)) *MockAuthService_SignIn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(SignInRequest))
	})
	return _c
}

func (_c *MockAuthService_SignIn_Call) Return(_a0 *Tokens, _a1 error) *MockAuthService_SignIn_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthService_SignIn_Call) RunAndReturn(run func(SignInRequest) (*Tokens, error)) *MockAuthService_SignIn_Call {
	_c.Call.Return(run)
	return _c
}

// SignOut provides a mock function with given fields: userId
func (_m *MockAuthService) SignOut(userId uuid.UUID) (*models.User, error) {
	ret := _m.Called(userId)

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*models.User, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.User); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthService_SignOut_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignOut'
type MockAuthService_SignOut_Call struct {
	*mock.Call
}

// SignOut is a helper method to define mock.On call
//   - userId uuid.UUID
func (_e *MockAuthService_Expecter) SignOut(userId interface{}) *MockAuthService_SignOut_Call {
	return &MockAuthService_SignOut_Call{Call: _e.mock.On("SignOut", userId)}
}

func (_c *MockAuthService_SignOut_Call) Run(run func(userId uuid.UUID)) *MockAuthService_SignOut_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *MockAuthService_SignOut_Call) Return(_a0 *models.User, _a1 error) *MockAuthService_SignOut_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthService_SignOut_Call) RunAndReturn(run func(uuid.UUID) (*models.User, error)) *MockAuthService_SignOut_Call {
	_c.Call.Return(run)
	return _c
}

// SignUp provides a mock function with given fields: payload
func (_m *MockAuthService) SignUp(payload SignUpRequest) (*models.User, error) {
	ret := _m.Called(payload)

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(SignUpRequest) (*models.User, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(SignUpRequest) *models.User); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(SignUpRequest) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthService_SignUp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignUp'
type MockAuthService_SignUp_Call struct {
	*mock.Call
}

// SignUp is a helper method to define mock.On call
//   - payload SignUpRequest
func (_e *MockAuthService_Expecter) SignUp(payload interface{}) *MockAuthService_SignUp_Call {
	return &MockAuthService_SignUp_Call{Call: _e.mock.On("SignUp", payload)}
}

func (_c *MockAuthService_SignUp_Call) Run(run func(payload SignUpRequest)) *MockAuthService_SignUp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(SignUpRequest))
	})
	return _c
}

func (_c *MockAuthService_SignUp_Call) Return(_a0 *models.User, _a1 error) *MockAuthService_SignUp_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthService_SignUp_Call) RunAndReturn(run func(SignUpRequest) (*models.User, error)) *MockAuthService_SignUp_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateToken provides a mock function with given fields: token, publicKey
func (_m *MockAuthService) ValidateToken(token string, publicKey string) (*jwt.Token, error) {
	ret := _m.Called(token, publicKey)

	var r0 *jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*jwt.Token, error)); ok {
		return rf(token, publicKey)
	}
	if rf, ok := ret.Get(0).(func(string, string) *jwt.Token); ok {
		r0 = rf(token, publicKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(token, publicKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthService_ValidateToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateToken'
type MockAuthService_ValidateToken_Call struct {
	*mock.Call
}

// ValidateToken is a helper method to define mock.On call
//   - token string
//   - publicKey string
func (_e *MockAuthService_Expecter) ValidateToken(token interface{}, publicKey interface{}) *MockAuthService_ValidateToken_Call {
	return &MockAuthService_ValidateToken_Call{Call: _e.mock.On("ValidateToken", token, publicKey)}
}

func (_c *MockAuthService_ValidateToken_Call) Run(run func(token string, publicKey string)) *MockAuthService_ValidateToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockAuthService_ValidateToken_Call) Return(_a0 *jwt.Token, _a1 error) *MockAuthService_ValidateToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthService_ValidateToken_Call) RunAndReturn(run func(string, string) (*jwt.Token, error)) *MockAuthService_ValidateToken_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyRefreshToken provides a mock function with given fields: encoded, refreshToken
func (_m *MockAuthService) VerifyRefreshToken(encoded string, refreshToken string) (bool, error) {
	ret := _m.Called(encoded, refreshToken)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(encoded, refreshToken)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(encoded, refreshToken)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(encoded, refreshToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthService_VerifyRefreshToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyRefreshToken'
type MockAuthService_VerifyRefreshToken_Call struct {
	*mock.Call
}

// VerifyRefreshToken is a helper method to define mock.On call
//   - encoded string
//   - refreshToken string
func (_e *MockAuthService_Expecter) VerifyRefreshToken(encoded interface{}, refreshToken interface{}) *MockAuthService_VerifyRefreshToken_Call {
	return &MockAuthService_VerifyRefreshToken_Call{Call: _e.mock.On("VerifyRefreshToken", encoded, refreshToken)}
}

func (_c *MockAuthService_VerifyRefreshToken_Call) Run(run func(encoded string, refreshToken string)) *MockAuthService_VerifyRefreshToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockAuthService_VerifyRefreshToken_Call) Return(_a0 bool, _a1 error) *MockAuthService_VerifyRefreshToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthService_VerifyRefreshToken_Call) RunAndReturn(run func(string, string) (bool, error)) *MockAuthService_VerifyRefreshToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAuthService creates a new instance of MockAuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAuthService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAuthService {
	mock := &MockAuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
