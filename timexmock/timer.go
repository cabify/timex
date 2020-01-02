// Code generated by mockery v1.0.0. DO NOT EDIT.

package timexmock

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Timer is an autogenerated mock type for the Timer type
type Timer struct {
	mock.Mock
}

// C provides a mock function with given fields:
func (_m *Timer) C() <-chan time.Time {
	ret := _m.Called()

	var r0 <-chan time.Time
	if rf, ok := ret.Get(0).(func() <-chan time.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan time.Time)
		}
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Timer) Stop() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
