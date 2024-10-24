// Code generated by mockery v2.42.0. DO NOT EDIT.

package books

import (
	contractsbooks "gotu-bookstore/cmd/gotu-bookstore/contracts/books"

	mock "github.com/stretchr/testify/mock"
)

// GetBookDetailsServiceInterface is an autogenerated mock type for the GetBookDetailsServiceInterface type
type GetBookDetailsServiceInterface struct {
	mock.Mock
}

type GetBookDetailsServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *GetBookDetailsServiceInterface) EXPECT() *GetBookDetailsServiceInterface_Expecter {
	return &GetBookDetailsServiceInterface_Expecter{mock: &_m.Mock}
}

// ProcessingGetBookDetails provides a mock function with given fields: request
func (_m *GetBookDetailsServiceInterface) ProcessingGetBookDetails(request contractsbooks.GetBookDetailsRequest) (*contractsbooks.GetBookDetailsResponse, error) {
	ret := _m.Called(request)

	if len(ret) == 0 {
		panic("no return value specified for ProcessingGetBookDetails")
	}

	var r0 *contractsbooks.GetBookDetailsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(contractsbooks.GetBookDetailsRequest) (*contractsbooks.GetBookDetailsResponse, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(contractsbooks.GetBookDetailsRequest) *contractsbooks.GetBookDetailsResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contractsbooks.GetBookDetailsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(contractsbooks.GetBookDetailsRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBookDetailsServiceInterface_ProcessingGetBookDetails_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessingGetBookDetails'
type GetBookDetailsServiceInterface_ProcessingGetBookDetails_Call struct {
	*mock.Call
}

// ProcessingGetBookDetails is a helper method to define mock.On call
//   - request contractsbooks.GetBookDetailsRequest
func (_e *GetBookDetailsServiceInterface_Expecter) ProcessingGetBookDetails(request interface{}) *GetBookDetailsServiceInterface_ProcessingGetBookDetails_Call {
	return &GetBookDetailsServiceInterface_ProcessingGetBookDetails_Call{Call: _e.mock.On("ProcessingGetBookDetails", request)}
}

func (_c *GetBookDetailsServiceInterface_ProcessingGetBookDetails_Call) Run(run func(request contractsbooks.GetBookDetailsRequest)) *GetBookDetailsServiceInterface_ProcessingGetBookDetails_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(contractsbooks.GetBookDetailsRequest))
	})
	return _c
}

func (_c *GetBookDetailsServiceInterface_ProcessingGetBookDetails_Call) Return(_a0 *contractsbooks.GetBookDetailsResponse, _a1 error) *GetBookDetailsServiceInterface_ProcessingGetBookDetails_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GetBookDetailsServiceInterface_ProcessingGetBookDetails_Call) RunAndReturn(run func(contractsbooks.GetBookDetailsRequest) (*contractsbooks.GetBookDetailsResponse, error)) *GetBookDetailsServiceInterface_ProcessingGetBookDetails_Call {
	_c.Call.Return(run)
	return _c
}

// NewGetBookDetailsServiceInterface creates a new instance of GetBookDetailsServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetBookDetailsServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetBookDetailsServiceInterface {
	mock := &GetBookDetailsServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
