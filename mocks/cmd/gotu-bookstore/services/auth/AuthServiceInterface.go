// Code generated by mockery v2.42.0. DO NOT EDIT.

package auth

import (
	config "gotu-bookstore/pkg/auth/config"
	dto "gotu-bookstore/pkg/auth/dto"

	mock "github.com/stretchr/testify/mock"
)

// AuthServiceInterface is an autogenerated mock type for the AuthServiceInterface type
type AuthServiceInterface struct {
	mock.Mock
}

type AuthServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthServiceInterface) EXPECT() *AuthServiceInterface_Expecter {
	return &AuthServiceInterface_Expecter{mock: &_m.Mock}
}

// GenerateTokenWithSessionDTO provides a mock function with given fields: _a0, sessionDTO
func (_m *AuthServiceInterface) GenerateTokenWithSessionDTO(_a0 config.BaseConfig, sessionDTO dto.SessionDTO) (string, error) {
	ret := _m.Called(_a0, sessionDTO)

	if len(ret) == 0 {
		panic("no return value specified for GenerateTokenWithSessionDTO")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(config.BaseConfig, dto.SessionDTO) (string, error)); ok {
		return rf(_a0, sessionDTO)
	}
	if rf, ok := ret.Get(0).(func(config.BaseConfig, dto.SessionDTO) string); ok {
		r0 = rf(_a0, sessionDTO)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(config.BaseConfig, dto.SessionDTO) error); ok {
		r1 = rf(_a0, sessionDTO)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthServiceInterface_GenerateTokenWithSessionDTO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateTokenWithSessionDTO'
type AuthServiceInterface_GenerateTokenWithSessionDTO_Call struct {
	*mock.Call
}

// GenerateTokenWithSessionDTO is a helper method to define mock.On call
//   - _a0 config.BaseConfig
//   - sessionDTO dto.SessionDTO
func (_e *AuthServiceInterface_Expecter) GenerateTokenWithSessionDTO(_a0 interface{}, sessionDTO interface{}) *AuthServiceInterface_GenerateTokenWithSessionDTO_Call {
	return &AuthServiceInterface_GenerateTokenWithSessionDTO_Call{Call: _e.mock.On("GenerateTokenWithSessionDTO", _a0, sessionDTO)}
}

func (_c *AuthServiceInterface_GenerateTokenWithSessionDTO_Call) Run(run func(_a0 config.BaseConfig, sessionDTO dto.SessionDTO)) *AuthServiceInterface_GenerateTokenWithSessionDTO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(config.BaseConfig), args[1].(dto.SessionDTO))
	})
	return _c
}

func (_c *AuthServiceInterface_GenerateTokenWithSessionDTO_Call) Return(_a0 string, _a1 error) *AuthServiceInterface_GenerateTokenWithSessionDTO_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthServiceInterface_GenerateTokenWithSessionDTO_Call) RunAndReturn(run func(config.BaseConfig, dto.SessionDTO) (string, error)) *AuthServiceInterface_GenerateTokenWithSessionDTO_Call {
	_c.Call.Return(run)
	return _c
}

// HashPassword provides a mock function with given fields: password
func (_m *AuthServiceInterface) HashPassword(password string) (string, error) {
	ret := _m.Called(password)

	if len(ret) == 0 {
		panic("no return value specified for HashPassword")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(password)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthServiceInterface_HashPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashPassword'
type AuthServiceInterface_HashPassword_Call struct {
	*mock.Call
}

// HashPassword is a helper method to define mock.On call
//   - password string
func (_e *AuthServiceInterface_Expecter) HashPassword(password interface{}) *AuthServiceInterface_HashPassword_Call {
	return &AuthServiceInterface_HashPassword_Call{Call: _e.mock.On("HashPassword", password)}
}

func (_c *AuthServiceInterface_HashPassword_Call) Run(run func(password string)) *AuthServiceInterface_HashPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AuthServiceInterface_HashPassword_Call) Return(_a0 string, _a1 error) *AuthServiceInterface_HashPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthServiceInterface_HashPassword_Call) RunAndReturn(run func(string) (string, error)) *AuthServiceInterface_HashPassword_Call {
	_c.Call.Return(run)
	return _c
}

// InvalidateToken provides a mock function with given fields: _a0, accessToken
func (_m *AuthServiceInterface) InvalidateToken(_a0 config.BaseConfig, accessToken string) error {
	ret := _m.Called(_a0, accessToken)

	if len(ret) == 0 {
		panic("no return value specified for InvalidateToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(config.BaseConfig, string) error); ok {
		r0 = rf(_a0, accessToken)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthServiceInterface_InvalidateToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InvalidateToken'
type AuthServiceInterface_InvalidateToken_Call struct {
	*mock.Call
}

// InvalidateToken is a helper method to define mock.On call
//   - _a0 config.BaseConfig
//   - accessToken string
func (_e *AuthServiceInterface_Expecter) InvalidateToken(_a0 interface{}, accessToken interface{}) *AuthServiceInterface_InvalidateToken_Call {
	return &AuthServiceInterface_InvalidateToken_Call{Call: _e.mock.On("InvalidateToken", _a0, accessToken)}
}

func (_c *AuthServiceInterface_InvalidateToken_Call) Run(run func(_a0 config.BaseConfig, accessToken string)) *AuthServiceInterface_InvalidateToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(config.BaseConfig), args[1].(string))
	})
	return _c
}

func (_c *AuthServiceInterface_InvalidateToken_Call) Return(_a0 error) *AuthServiceInterface_InvalidateToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthServiceInterface_InvalidateToken_Call) RunAndReturn(run func(config.BaseConfig, string) error) *AuthServiceInterface_InvalidateToken_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyPassword provides a mock function with given fields: hashedPassword, password
func (_m *AuthServiceInterface) VerifyPassword(hashedPassword string, password string) error {
	ret := _m.Called(hashedPassword, password)

	if len(ret) == 0 {
		panic("no return value specified for VerifyPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(hashedPassword, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AuthServiceInterface_VerifyPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyPassword'
type AuthServiceInterface_VerifyPassword_Call struct {
	*mock.Call
}

// VerifyPassword is a helper method to define mock.On call
//   - hashedPassword string
//   - password string
func (_e *AuthServiceInterface_Expecter) VerifyPassword(hashedPassword interface{}, password interface{}) *AuthServiceInterface_VerifyPassword_Call {
	return &AuthServiceInterface_VerifyPassword_Call{Call: _e.mock.On("VerifyPassword", hashedPassword, password)}
}

func (_c *AuthServiceInterface_VerifyPassword_Call) Run(run func(hashedPassword string, password string)) *AuthServiceInterface_VerifyPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *AuthServiceInterface_VerifyPassword_Call) Return(_a0 error) *AuthServiceInterface_VerifyPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthServiceInterface_VerifyPassword_Call) RunAndReturn(run func(string, string) error) *AuthServiceInterface_VerifyPassword_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyTokenToSessionDTO provides a mock function with given fields: _a0, accessToken, leeway
func (_m *AuthServiceInterface) VerifyTokenToSessionDTO(_a0 config.BaseConfig, accessToken string, leeway int64) (*dto.SessionDTO, error) {
	ret := _m.Called(_a0, accessToken, leeway)

	if len(ret) == 0 {
		panic("no return value specified for VerifyTokenToSessionDTO")
	}

	var r0 *dto.SessionDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(config.BaseConfig, string, int64) (*dto.SessionDTO, error)); ok {
		return rf(_a0, accessToken, leeway)
	}
	if rf, ok := ret.Get(0).(func(config.BaseConfig, string, int64) *dto.SessionDTO); ok {
		r0 = rf(_a0, accessToken, leeway)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.SessionDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(config.BaseConfig, string, int64) error); ok {
		r1 = rf(_a0, accessToken, leeway)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthServiceInterface_VerifyTokenToSessionDTO_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyTokenToSessionDTO'
type AuthServiceInterface_VerifyTokenToSessionDTO_Call struct {
	*mock.Call
}

// VerifyTokenToSessionDTO is a helper method to define mock.On call
//   - _a0 config.BaseConfig
//   - accessToken string
//   - leeway int64
func (_e *AuthServiceInterface_Expecter) VerifyTokenToSessionDTO(_a0 interface{}, accessToken interface{}, leeway interface{}) *AuthServiceInterface_VerifyTokenToSessionDTO_Call {
	return &AuthServiceInterface_VerifyTokenToSessionDTO_Call{Call: _e.mock.On("VerifyTokenToSessionDTO", _a0, accessToken, leeway)}
}

func (_c *AuthServiceInterface_VerifyTokenToSessionDTO_Call) Run(run func(_a0 config.BaseConfig, accessToken string, leeway int64)) *AuthServiceInterface_VerifyTokenToSessionDTO_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(config.BaseConfig), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *AuthServiceInterface_VerifyTokenToSessionDTO_Call) Return(_a0 *dto.SessionDTO, _a1 error) *AuthServiceInterface_VerifyTokenToSessionDTO_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AuthServiceInterface_VerifyTokenToSessionDTO_Call) RunAndReturn(run func(config.BaseConfig, string, int64) (*dto.SessionDTO, error)) *AuthServiceInterface_VerifyTokenToSessionDTO_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuthServiceInterface creates a new instance of AuthServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthServiceInterface {
	mock := &AuthServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
