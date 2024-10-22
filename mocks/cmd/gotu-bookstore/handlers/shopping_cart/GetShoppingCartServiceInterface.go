// Code generated by mockery v2.42.0. DO NOT EDIT.

package shopping_cart

import (
	contractsshopping_cart "gotu-bookstore/cmd/gotu-bookstore/contracts/shopping_cart"

	mock "github.com/stretchr/testify/mock"
)

// GetShoppingCartServiceInterface is an autogenerated mock type for the GetShoppingCartServiceInterface type
type GetShoppingCartServiceInterface struct {
	mock.Mock
}

type GetShoppingCartServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *GetShoppingCartServiceInterface) EXPECT() *GetShoppingCartServiceInterface_Expecter {
	return &GetShoppingCartServiceInterface_Expecter{mock: &_m.Mock}
}

// ProcessingGetShoppingCart provides a mock function with given fields:
func (_m *GetShoppingCartServiceInterface) ProcessingGetShoppingCart() (*contractsshopping_cart.ShoppingCartResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ProcessingGetShoppingCart")
	}

	var r0 *contractsshopping_cart.ShoppingCartResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*contractsshopping_cart.ShoppingCartResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *contractsshopping_cart.ShoppingCartResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contractsshopping_cart.ShoppingCartResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetShoppingCartServiceInterface_ProcessingGetShoppingCart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessingGetShoppingCart'
type GetShoppingCartServiceInterface_ProcessingGetShoppingCart_Call struct {
	*mock.Call
}

// ProcessingGetShoppingCart is a helper method to define mock.On call
func (_e *GetShoppingCartServiceInterface_Expecter) ProcessingGetShoppingCart() *GetShoppingCartServiceInterface_ProcessingGetShoppingCart_Call {
	return &GetShoppingCartServiceInterface_ProcessingGetShoppingCart_Call{Call: _e.mock.On("ProcessingGetShoppingCart")}
}

func (_c *GetShoppingCartServiceInterface_ProcessingGetShoppingCart_Call) Run(run func()) *GetShoppingCartServiceInterface_ProcessingGetShoppingCart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GetShoppingCartServiceInterface_ProcessingGetShoppingCart_Call) Return(_a0 *contractsshopping_cart.ShoppingCartResponse, _a1 error) *GetShoppingCartServiceInterface_ProcessingGetShoppingCart_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GetShoppingCartServiceInterface_ProcessingGetShoppingCart_Call) RunAndReturn(run func() (*contractsshopping_cart.ShoppingCartResponse, error)) *GetShoppingCartServiceInterface_ProcessingGetShoppingCart_Call {
	_c.Call.Return(run)
	return _c
}

// NewGetShoppingCartServiceInterface creates a new instance of GetShoppingCartServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetShoppingCartServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetShoppingCartServiceInterface {
	mock := &GetShoppingCartServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
