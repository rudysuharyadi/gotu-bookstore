// Code generated by mockery v2.42.0. DO NOT EDIT.

package transactions

import (
	contractstransactions "gotu-bookstore/cmd/gotu-bookstore/contracts/transactions"

	mock "github.com/stretchr/testify/mock"
)

// GetTransactionsServiceInterface is an autogenerated mock type for the GetTransactionsServiceInterface type
type GetTransactionsServiceInterface struct {
	mock.Mock
}

type GetTransactionsServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *GetTransactionsServiceInterface) EXPECT() *GetTransactionsServiceInterface_Expecter {
	return &GetTransactionsServiceInterface_Expecter{mock: &_m.Mock}
}

// ProcessingGetTransactions provides a mock function with given fields: request
func (_m *GetTransactionsServiceInterface) ProcessingGetTransactions(request contractstransactions.GetTransactionsRequest) (*contractstransactions.GetTransactionsResponse, map[string]interface{}, error) {
	ret := _m.Called(request)

	if len(ret) == 0 {
		panic("no return value specified for ProcessingGetTransactions")
	}

	var r0 *contractstransactions.GetTransactionsResponse
	var r1 map[string]interface{}
	var r2 error
	if rf, ok := ret.Get(0).(func(contractstransactions.GetTransactionsRequest) (*contractstransactions.GetTransactionsResponse, map[string]interface{}, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(contractstransactions.GetTransactionsRequest) *contractstransactions.GetTransactionsResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contractstransactions.GetTransactionsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(contractstransactions.GetTransactionsRequest) map[string]interface{}); ok {
		r1 = rf(request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(2).(func(contractstransactions.GetTransactionsRequest) error); ok {
		r2 = rf(request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTransactionsServiceInterface_ProcessingGetTransactions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessingGetTransactions'
type GetTransactionsServiceInterface_ProcessingGetTransactions_Call struct {
	*mock.Call
}

// ProcessingGetTransactions is a helper method to define mock.On call
//   - request contractstransactions.GetTransactionsRequest
func (_e *GetTransactionsServiceInterface_Expecter) ProcessingGetTransactions(request interface{}) *GetTransactionsServiceInterface_ProcessingGetTransactions_Call {
	return &GetTransactionsServiceInterface_ProcessingGetTransactions_Call{Call: _e.mock.On("ProcessingGetTransactions", request)}
}

func (_c *GetTransactionsServiceInterface_ProcessingGetTransactions_Call) Run(run func(request contractstransactions.GetTransactionsRequest)) *GetTransactionsServiceInterface_ProcessingGetTransactions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(contractstransactions.GetTransactionsRequest))
	})
	return _c
}

func (_c *GetTransactionsServiceInterface_ProcessingGetTransactions_Call) Return(_a0 *contractstransactions.GetTransactionsResponse, _a1 map[string]interface{}, _a2 error) *GetTransactionsServiceInterface_ProcessingGetTransactions_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GetTransactionsServiceInterface_ProcessingGetTransactions_Call) RunAndReturn(run func(contractstransactions.GetTransactionsRequest) (*contractstransactions.GetTransactionsResponse, map[string]interface{}, error)) *GetTransactionsServiceInterface_ProcessingGetTransactions_Call {
	_c.Call.Return(run)
	return _c
}

// NewGetTransactionsServiceInterface creates a new instance of GetTransactionsServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetTransactionsServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetTransactionsServiceInterface {
	mock := &GetTransactionsServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
