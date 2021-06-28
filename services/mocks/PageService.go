// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entities "github.com/manicar2093/web-monitor/entities"
	mock "github.com/stretchr/testify/mock"
)

// PageService is an autogenerated mock type for the PageService type
type PageService struct {
	mock.Mock
}

// AddPage provides a mock function with given fields: page
func (_m *PageService) AddPage(page entities.Page) (entities.Page, error) {
	ret := _m.Called(page)

	var r0 entities.Page
	if rf, ok := ret.Get(0).(func(entities.Page) entities.Page); ok {
		r0 = rf(page)
	} else {
		r0 = ret.Get(0).(entities.Page)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.Page) error); ok {
		r1 = rf(page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PageExists provides a mock function with given fields: url
func (_m *PageService) PageExists(url string) (bool, error) {
	ret := _m.Called(url)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}