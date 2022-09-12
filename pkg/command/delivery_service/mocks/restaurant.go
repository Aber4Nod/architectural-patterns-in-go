// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// restaurant is an autogenerated mock type for the restaurant type
type restaurant struct {
	mock.Mock
}

type restaurant_Expecter struct {
	mock *mock.Mock
}

func (_m *restaurant) EXPECT() *restaurant_Expecter {
	return &restaurant_Expecter{mock: &_m.Mock}
}

// CookOrder provides a mock function with given fields: dish
func (_m *restaurant) CookOrder(dish string) error {
	ret := _m.Called(dish)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(dish)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// restaurant_CookOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CookOrder'
type restaurant_CookOrder_Call struct {
	*mock.Call
}

// CookOrder is a helper method to define mock.On call
//  - dish string
func (_e *restaurant_Expecter) CookOrder(dish interface{}) *restaurant_CookOrder_Call {
	return &restaurant_CookOrder_Call{Call: _e.mock.On("CookOrder", dish)}
}

func (_c *restaurant_CookOrder_Call) Run(run func(dish string)) *restaurant_CookOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *restaurant_CookOrder_Call) Return(_a0 error) *restaurant_CookOrder_Call {
	_c.Call.Return(_a0)
	return _c
}

// GiveMenu provides a mock function with given fields:
func (_m *restaurant) GiveMenu() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// restaurant_GiveMenu_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GiveMenu'
type restaurant_GiveMenu_Call struct {
	*mock.Call
}

// GiveMenu is a helper method to define mock.On call
func (_e *restaurant_Expecter) GiveMenu() *restaurant_GiveMenu_Call {
	return &restaurant_GiveMenu_Call{Call: _e.mock.On("GiveMenu")}
}

func (_c *restaurant_GiveMenu_Call) Run(run func()) *restaurant_GiveMenu_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *restaurant_GiveMenu_Call) Return(_a0 error) *restaurant_GiveMenu_Call {
	_c.Call.Return(_a0)
	return _c
}

// Добавлять не обязательно
func NewRestaurant() *restaurant {
	return &restaurant{}
}
