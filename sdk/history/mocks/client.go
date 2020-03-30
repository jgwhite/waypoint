// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	component "github.com/mitchellh/devflow/sdk/component"

	history "github.com/mitchellh/devflow/sdk/history"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Deployments provides a mock function with given fields: _a0, _a1
func (_m *Client) Deployments(_a0 context.Context, _a1 *history.Lookup) ([]component.Deployment, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []component.Deployment
	if rf, ok := ret.Get(0).(func(context.Context, *history.Lookup) []component.Deployment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]component.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *history.Lookup) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
