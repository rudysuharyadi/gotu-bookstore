// Code generated by mockery v2.42.0. DO NOT EDIT.

package auth

import (
	contractsauth "gotu-bookstore/cmd/gotu-bookstore/contracts/auth"

	mock "github.com/stretchr/testify/mock"
)

// UserValidatorInterface is an autogenerated mock type for the UserValidatorInterface type
type UserValidatorInterface struct {
	mock.Mock
}

type UserValidatorInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *UserValidatorInterface) EXPECT() *UserValidatorInterface_Expecter {
	return &UserValidatorInterface_Expecter{mock: &_m.Mock}
}

// Validate provides a mock function with given fields: request
func (_m *UserValidatorInterface) Validate(request contractsauth.RegisterRequest) []error {
	ret := _m.Called(request)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 []error
	if rf, ok := ret.Get(0).(func(contractsauth.RegisterRequest) []error); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]error)
		}
	}

	return r0
}

// UserValidatorInterface_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type UserValidatorInterface_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - request contractsauth.RegisterRequest
func (_e *UserValidatorInterface_Expecter) Validate(request interface{}) *UserValidatorInterface_Validate_Call {
	return &UserValidatorInterface_Validate_Call{Call: _e.mock.On("Validate", request)}
}

func (_c *UserValidatorInterface_Validate_Call) Run(run func(request contractsauth.RegisterRequest)) *UserValidatorInterface_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(contractsauth.RegisterRequest))
	})
	return _c
}

func (_c *UserValidatorInterface_Validate_Call) Return(_a0 []error) *UserValidatorInterface_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserValidatorInterface_Validate_Call) RunAndReturn(run func(contractsauth.RegisterRequest) []error) *UserValidatorInterface_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserValidatorInterface creates a new instance of UserValidatorInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserValidatorInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserValidatorInterface {
	mock := &UserValidatorInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
