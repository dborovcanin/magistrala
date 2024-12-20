// Copyright (c) Abstract Machines

// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/absmach/supermq/internal/grpc/channels/v1"
)

// ChannelsServiceClient is an autogenerated mock type for the ChannelsServiceClient type
type ChannelsServiceClient struct {
	mock.Mock
}

type ChannelsServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *ChannelsServiceClient) EXPECT() *ChannelsServiceClient_Expecter {
	return &ChannelsServiceClient_Expecter{mock: &_m.Mock}
}

// Authorize provides a mock function with given fields: ctx, in, opts
func (_m *ChannelsServiceClient) Authorize(ctx context.Context, in *v1.AuthzReq, opts ...grpc.CallOption) (*v1.AuthzRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Authorize")
	}

	var r0 *v1.AuthzRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.AuthzReq, ...grpc.CallOption) (*v1.AuthzRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.AuthzReq, ...grpc.CallOption) *v1.AuthzRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.AuthzRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.AuthzReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChannelsServiceClient_Authorize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Authorize'
type ChannelsServiceClient_Authorize_Call struct {
	*mock.Call
}

// Authorize is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.AuthzReq
//   - opts ...grpc.CallOption
func (_e *ChannelsServiceClient_Expecter) Authorize(ctx interface{}, in interface{}, opts ...interface{}) *ChannelsServiceClient_Authorize_Call {
	return &ChannelsServiceClient_Authorize_Call{Call: _e.mock.On("Authorize",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ChannelsServiceClient_Authorize_Call) Run(run func(ctx context.Context, in *v1.AuthzReq, opts ...grpc.CallOption)) *ChannelsServiceClient_Authorize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.AuthzReq), variadicArgs...)
	})
	return _c
}

func (_c *ChannelsServiceClient_Authorize_Call) Return(_a0 *v1.AuthzRes, _a1 error) *ChannelsServiceClient_Authorize_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChannelsServiceClient_Authorize_Call) RunAndReturn(run func(context.Context, *v1.AuthzReq, ...grpc.CallOption) (*v1.AuthzRes, error)) *ChannelsServiceClient_Authorize_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveThingConnections provides a mock function with given fields: ctx, in, opts
func (_m *ChannelsServiceClient) RemoveThingConnections(ctx context.Context, in *v1.RemoveThingConnectionsReq, opts ...grpc.CallOption) (*v1.RemoveThingConnectionsRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RemoveThingConnections")
	}

	var r0 *v1.RemoveThingConnectionsRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.RemoveThingConnectionsReq, ...grpc.CallOption) (*v1.RemoveThingConnectionsRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.RemoveThingConnectionsReq, ...grpc.CallOption) *v1.RemoveThingConnectionsRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.RemoveThingConnectionsRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.RemoveThingConnectionsReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChannelsServiceClient_RemoveThingConnections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveThingConnections'
type ChannelsServiceClient_RemoveThingConnections_Call struct {
	*mock.Call
}

// RemoveThingConnections is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.RemoveThingConnectionsReq
//   - opts ...grpc.CallOption
func (_e *ChannelsServiceClient_Expecter) RemoveThingConnections(ctx interface{}, in interface{}, opts ...interface{}) *ChannelsServiceClient_RemoveThingConnections_Call {
	return &ChannelsServiceClient_RemoveThingConnections_Call{Call: _e.mock.On("RemoveThingConnections",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ChannelsServiceClient_RemoveThingConnections_Call) Run(run func(ctx context.Context, in *v1.RemoveThingConnectionsReq, opts ...grpc.CallOption)) *ChannelsServiceClient_RemoveThingConnections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.RemoveThingConnectionsReq), variadicArgs...)
	})
	return _c
}

func (_c *ChannelsServiceClient_RemoveThingConnections_Call) Return(_a0 *v1.RemoveThingConnectionsRes, _a1 error) *ChannelsServiceClient_RemoveThingConnections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChannelsServiceClient_RemoveThingConnections_Call) RunAndReturn(run func(context.Context, *v1.RemoveThingConnectionsReq, ...grpc.CallOption) (*v1.RemoveThingConnectionsRes, error)) *ChannelsServiceClient_RemoveThingConnections_Call {
	_c.Call.Return(run)
	return _c
}

// UnsetParentGroupFromChannels provides a mock function with given fields: ctx, in, opts
func (_m *ChannelsServiceClient) UnsetParentGroupFromChannels(ctx context.Context, in *v1.UnsetParentGroupFromChannelsReq, opts ...grpc.CallOption) (*v1.UnsetParentGroupFromChannelsRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UnsetParentGroupFromChannels")
	}

	var r0 *v1.UnsetParentGroupFromChannelsRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.UnsetParentGroupFromChannelsReq, ...grpc.CallOption) (*v1.UnsetParentGroupFromChannelsRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.UnsetParentGroupFromChannelsReq, ...grpc.CallOption) *v1.UnsetParentGroupFromChannelsRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.UnsetParentGroupFromChannelsRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.UnsetParentGroupFromChannelsReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChannelsServiceClient_UnsetParentGroupFromChannels_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnsetParentGroupFromChannels'
type ChannelsServiceClient_UnsetParentGroupFromChannels_Call struct {
	*mock.Call
}

// UnsetParentGroupFromChannels is a helper method to define mock.On call
//   - ctx context.Context
//   - in *v1.UnsetParentGroupFromChannelsReq
//   - opts ...grpc.CallOption
func (_e *ChannelsServiceClient_Expecter) UnsetParentGroupFromChannels(ctx interface{}, in interface{}, opts ...interface{}) *ChannelsServiceClient_UnsetParentGroupFromChannels_Call {
	return &ChannelsServiceClient_UnsetParentGroupFromChannels_Call{Call: _e.mock.On("UnsetParentGroupFromChannels",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ChannelsServiceClient_UnsetParentGroupFromChannels_Call) Run(run func(ctx context.Context, in *v1.UnsetParentGroupFromChannelsReq, opts ...grpc.CallOption)) *ChannelsServiceClient_UnsetParentGroupFromChannels_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*v1.UnsetParentGroupFromChannelsReq), variadicArgs...)
	})
	return _c
}

func (_c *ChannelsServiceClient_UnsetParentGroupFromChannels_Call) Return(_a0 *v1.UnsetParentGroupFromChannelsRes, _a1 error) *ChannelsServiceClient_UnsetParentGroupFromChannels_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChannelsServiceClient_UnsetParentGroupFromChannels_Call) RunAndReturn(run func(context.Context, *v1.UnsetParentGroupFromChannelsReq, ...grpc.CallOption) (*v1.UnsetParentGroupFromChannelsRes, error)) *ChannelsServiceClient_UnsetParentGroupFromChannels_Call {
	_c.Call.Return(run)
	return _c
}

// NewChannelsServiceClient creates a new instance of ChannelsServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChannelsServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChannelsServiceClient {
	mock := &ChannelsServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
