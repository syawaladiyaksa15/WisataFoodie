// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	favourites "capstone/group3/features/favourites"

	mock "github.com/stretchr/testify/mock"
)

// FavouriteData is an autogenerated mock type for the Data type
type FavouriteData struct {
	mock.Mock
}

// AddFavDB provides a mock function with given fields: idResto, idFromToken
func (_m *FavouriteData) AddFavDB(idResto int, idFromToken int) (int, error) {
	ret := _m.Called(idResto, idFromToken)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(idResto, idFromToken)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(idResto, idFromToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllRestoData provides a mock function with given fields: limitint, offsetint, idFromToken
func (_m *FavouriteData) AllRestoData(limitint int, offsetint int, idFromToken int) ([]favourites.RestoCore, error) {
	ret := _m.Called(limitint, offsetint, idFromToken)

	var r0 []favourites.RestoCore
	if rf, ok := ret.Get(0).(func(int, int, int) []favourites.RestoCore); ok {
		r0 = rf(limitint, offsetint, idFromToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]favourites.RestoCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, int) error); ok {
		r1 = rf(limitint, offsetint, idFromToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFavDB provides a mock function with given fields: idResto, idFromToken
func (_m *FavouriteData) DeleteFavDB(idResto int, idFromToken int) (int, error) {
	ret := _m.Called(idResto, idFromToken)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(idResto, idFromToken)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(idResto, idFromToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RatingData provides a mock function with given fields: idResto
func (_m *FavouriteData) RatingData(idResto int) (float64, error) {
	ret := _m.Called(idResto)

	var r0 float64
	if rf, ok := ret.Get(0).(func(int) float64); ok {
		r0 = rf(idResto)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idResto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoImageData provides a mock function with given fields: idResto
func (_m *FavouriteData) RestoImageData(idResto int) (string, error) {
	ret := _m.Called(idResto)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(idResto)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idResto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFavouriteData interface {
	mock.TestingT
	Cleanup(func())
}

// NewFavouriteData creates a new instance of FavouriteData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFavouriteData(t mockConstructorTestingTNewFavouriteData) *FavouriteData {
	mock := &FavouriteData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
