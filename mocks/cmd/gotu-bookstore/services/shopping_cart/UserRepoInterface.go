// Code generated by mockery v2.42.0. DO NOT EDIT.

package shopping_cart

import (
	models "gotu-bookstore/cmd/gotu-bookstore/models"

	mock "github.com/stretchr/testify/mock"
)

// UserRepoInterface is an autogenerated mock type for the UserRepoInterface type
type UserRepoInterface struct {
	mock.Mock
}

type UserRepoInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *UserRepoInterface) EXPECT() *UserRepoInterface_Expecter {
	return &UserRepoInterface_Expecter{mock: &_m.Mock}
}

// GetById provides a mock function with given fields: id
func (_m *UserRepoInterface) GetById(id string) (*models.Users, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *models.Users
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Users, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Users); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Users)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepoInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type UserRepoInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - id string
func (_e *UserRepoInterface_Expecter) GetById(id interface{}) *UserRepoInterface_GetById_Call {
	return &UserRepoInterface_GetById_Call{Call: _e.mock.On("GetById", id)}
}

func (_c *UserRepoInterface_GetById_Call) Run(run func(id string)) *UserRepoInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UserRepoInterface_GetById_Call) Return(_a0 *models.Users, _a1 error) *UserRepoInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepoInterface_GetById_Call) RunAndReturn(run func(string) (*models.Users, error)) *UserRepoInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserRepoInterface creates a new instance of UserRepoInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepoInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepoInterface {
	mock := &UserRepoInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
