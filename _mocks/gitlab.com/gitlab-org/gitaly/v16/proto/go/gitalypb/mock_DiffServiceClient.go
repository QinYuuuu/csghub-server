// Code generated by mockery v2.53.0. DO NOT EDIT.

package gitalypb

import (
	context "context"

	gitalypb "gitlab.com/gitlab-org/gitaly/v16/proto/go/gitalypb"
	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockDiffServiceClient is an autogenerated mock type for the DiffServiceClient type
type MockDiffServiceClient struct {
	mock.Mock
}

type MockDiffServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDiffServiceClient) EXPECT() *MockDiffServiceClient_Expecter {
	return &MockDiffServiceClient_Expecter{mock: &_m.Mock}
}

// CommitDelta provides a mock function with given fields: ctx, in, opts
func (_m *MockDiffServiceClient) CommitDelta(ctx context.Context, in *gitalypb.CommitDeltaRequest, opts ...grpc.CallOption) (gitalypb.DiffService_CommitDeltaClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CommitDelta")
	}

	var r0 gitalypb.DiffService_CommitDeltaClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.CommitDeltaRequest, ...grpc.CallOption) (gitalypb.DiffService_CommitDeltaClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.CommitDeltaRequest, ...grpc.CallOption) gitalypb.DiffService_CommitDeltaClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gitalypb.DiffService_CommitDeltaClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gitalypb.CommitDeltaRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDiffServiceClient_CommitDelta_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CommitDelta'
type MockDiffServiceClient_CommitDelta_Call struct {
	*mock.Call
}

// CommitDelta is a helper method to define mock.On call
//   - ctx context.Context
//   - in *gitalypb.CommitDeltaRequest
//   - opts ...grpc.CallOption
func (_e *MockDiffServiceClient_Expecter) CommitDelta(ctx interface{}, in interface{}, opts ...interface{}) *MockDiffServiceClient_CommitDelta_Call {
	return &MockDiffServiceClient_CommitDelta_Call{Call: _e.mock.On("CommitDelta",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockDiffServiceClient_CommitDelta_Call) Run(run func(ctx context.Context, in *gitalypb.CommitDeltaRequest, opts ...grpc.CallOption)) *MockDiffServiceClient_CommitDelta_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*gitalypb.CommitDeltaRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockDiffServiceClient_CommitDelta_Call) Return(_a0 gitalypb.DiffService_CommitDeltaClient, _a1 error) *MockDiffServiceClient_CommitDelta_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDiffServiceClient_CommitDelta_Call) RunAndReturn(run func(context.Context, *gitalypb.CommitDeltaRequest, ...grpc.CallOption) (gitalypb.DiffService_CommitDeltaClient, error)) *MockDiffServiceClient_CommitDelta_Call {
	_c.Call.Return(run)
	return _c
}

// CommitDiff provides a mock function with given fields: ctx, in, opts
func (_m *MockDiffServiceClient) CommitDiff(ctx context.Context, in *gitalypb.CommitDiffRequest, opts ...grpc.CallOption) (gitalypb.DiffService_CommitDiffClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CommitDiff")
	}

	var r0 gitalypb.DiffService_CommitDiffClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.CommitDiffRequest, ...grpc.CallOption) (gitalypb.DiffService_CommitDiffClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.CommitDiffRequest, ...grpc.CallOption) gitalypb.DiffService_CommitDiffClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gitalypb.DiffService_CommitDiffClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gitalypb.CommitDiffRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDiffServiceClient_CommitDiff_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CommitDiff'
type MockDiffServiceClient_CommitDiff_Call struct {
	*mock.Call
}

// CommitDiff is a helper method to define mock.On call
//   - ctx context.Context
//   - in *gitalypb.CommitDiffRequest
//   - opts ...grpc.CallOption
func (_e *MockDiffServiceClient_Expecter) CommitDiff(ctx interface{}, in interface{}, opts ...interface{}) *MockDiffServiceClient_CommitDiff_Call {
	return &MockDiffServiceClient_CommitDiff_Call{Call: _e.mock.On("CommitDiff",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockDiffServiceClient_CommitDiff_Call) Run(run func(ctx context.Context, in *gitalypb.CommitDiffRequest, opts ...grpc.CallOption)) *MockDiffServiceClient_CommitDiff_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*gitalypb.CommitDiffRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockDiffServiceClient_CommitDiff_Call) Return(_a0 gitalypb.DiffService_CommitDiffClient, _a1 error) *MockDiffServiceClient_CommitDiff_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDiffServiceClient_CommitDiff_Call) RunAndReturn(run func(context.Context, *gitalypb.CommitDiffRequest, ...grpc.CallOption) (gitalypb.DiffService_CommitDiffClient, error)) *MockDiffServiceClient_CommitDiff_Call {
	_c.Call.Return(run)
	return _c
}

// DiffStats provides a mock function with given fields: ctx, in, opts
func (_m *MockDiffServiceClient) DiffStats(ctx context.Context, in *gitalypb.DiffStatsRequest, opts ...grpc.CallOption) (gitalypb.DiffService_DiffStatsClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DiffStats")
	}

	var r0 gitalypb.DiffService_DiffStatsClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.DiffStatsRequest, ...grpc.CallOption) (gitalypb.DiffService_DiffStatsClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.DiffStatsRequest, ...grpc.CallOption) gitalypb.DiffService_DiffStatsClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gitalypb.DiffService_DiffStatsClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gitalypb.DiffStatsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDiffServiceClient_DiffStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DiffStats'
type MockDiffServiceClient_DiffStats_Call struct {
	*mock.Call
}

// DiffStats is a helper method to define mock.On call
//   - ctx context.Context
//   - in *gitalypb.DiffStatsRequest
//   - opts ...grpc.CallOption
func (_e *MockDiffServiceClient_Expecter) DiffStats(ctx interface{}, in interface{}, opts ...interface{}) *MockDiffServiceClient_DiffStats_Call {
	return &MockDiffServiceClient_DiffStats_Call{Call: _e.mock.On("DiffStats",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockDiffServiceClient_DiffStats_Call) Run(run func(ctx context.Context, in *gitalypb.DiffStatsRequest, opts ...grpc.CallOption)) *MockDiffServiceClient_DiffStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*gitalypb.DiffStatsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockDiffServiceClient_DiffStats_Call) Return(_a0 gitalypb.DiffService_DiffStatsClient, _a1 error) *MockDiffServiceClient_DiffStats_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDiffServiceClient_DiffStats_Call) RunAndReturn(run func(context.Context, *gitalypb.DiffStatsRequest, ...grpc.CallOption) (gitalypb.DiffService_DiffStatsClient, error)) *MockDiffServiceClient_DiffStats_Call {
	_c.Call.Return(run)
	return _c
}

// FindChangedPaths provides a mock function with given fields: ctx, in, opts
func (_m *MockDiffServiceClient) FindChangedPaths(ctx context.Context, in *gitalypb.FindChangedPathsRequest, opts ...grpc.CallOption) (gitalypb.DiffService_FindChangedPathsClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for FindChangedPaths")
	}

	var r0 gitalypb.DiffService_FindChangedPathsClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.FindChangedPathsRequest, ...grpc.CallOption) (gitalypb.DiffService_FindChangedPathsClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.FindChangedPathsRequest, ...grpc.CallOption) gitalypb.DiffService_FindChangedPathsClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gitalypb.DiffService_FindChangedPathsClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gitalypb.FindChangedPathsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDiffServiceClient_FindChangedPaths_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindChangedPaths'
type MockDiffServiceClient_FindChangedPaths_Call struct {
	*mock.Call
}

// FindChangedPaths is a helper method to define mock.On call
//   - ctx context.Context
//   - in *gitalypb.FindChangedPathsRequest
//   - opts ...grpc.CallOption
func (_e *MockDiffServiceClient_Expecter) FindChangedPaths(ctx interface{}, in interface{}, opts ...interface{}) *MockDiffServiceClient_FindChangedPaths_Call {
	return &MockDiffServiceClient_FindChangedPaths_Call{Call: _e.mock.On("FindChangedPaths",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockDiffServiceClient_FindChangedPaths_Call) Run(run func(ctx context.Context, in *gitalypb.FindChangedPathsRequest, opts ...grpc.CallOption)) *MockDiffServiceClient_FindChangedPaths_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*gitalypb.FindChangedPathsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockDiffServiceClient_FindChangedPaths_Call) Return(_a0 gitalypb.DiffService_FindChangedPathsClient, _a1 error) *MockDiffServiceClient_FindChangedPaths_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDiffServiceClient_FindChangedPaths_Call) RunAndReturn(run func(context.Context, *gitalypb.FindChangedPathsRequest, ...grpc.CallOption) (gitalypb.DiffService_FindChangedPathsClient, error)) *MockDiffServiceClient_FindChangedPaths_Call {
	_c.Call.Return(run)
	return _c
}

// GetPatchID provides a mock function with given fields: ctx, in, opts
func (_m *MockDiffServiceClient) GetPatchID(ctx context.Context, in *gitalypb.GetPatchIDRequest, opts ...grpc.CallOption) (*gitalypb.GetPatchIDResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetPatchID")
	}

	var r0 *gitalypb.GetPatchIDResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.GetPatchIDRequest, ...grpc.CallOption) (*gitalypb.GetPatchIDResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.GetPatchIDRequest, ...grpc.CallOption) *gitalypb.GetPatchIDResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitalypb.GetPatchIDResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gitalypb.GetPatchIDRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDiffServiceClient_GetPatchID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPatchID'
type MockDiffServiceClient_GetPatchID_Call struct {
	*mock.Call
}

// GetPatchID is a helper method to define mock.On call
//   - ctx context.Context
//   - in *gitalypb.GetPatchIDRequest
//   - opts ...grpc.CallOption
func (_e *MockDiffServiceClient_Expecter) GetPatchID(ctx interface{}, in interface{}, opts ...interface{}) *MockDiffServiceClient_GetPatchID_Call {
	return &MockDiffServiceClient_GetPatchID_Call{Call: _e.mock.On("GetPatchID",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockDiffServiceClient_GetPatchID_Call) Run(run func(ctx context.Context, in *gitalypb.GetPatchIDRequest, opts ...grpc.CallOption)) *MockDiffServiceClient_GetPatchID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*gitalypb.GetPatchIDRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockDiffServiceClient_GetPatchID_Call) Return(_a0 *gitalypb.GetPatchIDResponse, _a1 error) *MockDiffServiceClient_GetPatchID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDiffServiceClient_GetPatchID_Call) RunAndReturn(run func(context.Context, *gitalypb.GetPatchIDRequest, ...grpc.CallOption) (*gitalypb.GetPatchIDResponse, error)) *MockDiffServiceClient_GetPatchID_Call {
	_c.Call.Return(run)
	return _c
}

// RangeDiff provides a mock function with given fields: ctx, in, opts
func (_m *MockDiffServiceClient) RangeDiff(ctx context.Context, in *gitalypb.RangeDiffRequest, opts ...grpc.CallOption) (gitalypb.DiffService_RangeDiffClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RangeDiff")
	}

	var r0 gitalypb.DiffService_RangeDiffClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.RangeDiffRequest, ...grpc.CallOption) (gitalypb.DiffService_RangeDiffClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.RangeDiffRequest, ...grpc.CallOption) gitalypb.DiffService_RangeDiffClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gitalypb.DiffService_RangeDiffClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gitalypb.RangeDiffRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDiffServiceClient_RangeDiff_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RangeDiff'
type MockDiffServiceClient_RangeDiff_Call struct {
	*mock.Call
}

// RangeDiff is a helper method to define mock.On call
//   - ctx context.Context
//   - in *gitalypb.RangeDiffRequest
//   - opts ...grpc.CallOption
func (_e *MockDiffServiceClient_Expecter) RangeDiff(ctx interface{}, in interface{}, opts ...interface{}) *MockDiffServiceClient_RangeDiff_Call {
	return &MockDiffServiceClient_RangeDiff_Call{Call: _e.mock.On("RangeDiff",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockDiffServiceClient_RangeDiff_Call) Run(run func(ctx context.Context, in *gitalypb.RangeDiffRequest, opts ...grpc.CallOption)) *MockDiffServiceClient_RangeDiff_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*gitalypb.RangeDiffRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockDiffServiceClient_RangeDiff_Call) Return(_a0 gitalypb.DiffService_RangeDiffClient, _a1 error) *MockDiffServiceClient_RangeDiff_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDiffServiceClient_RangeDiff_Call) RunAndReturn(run func(context.Context, *gitalypb.RangeDiffRequest, ...grpc.CallOption) (gitalypb.DiffService_RangeDiffClient, error)) *MockDiffServiceClient_RangeDiff_Call {
	_c.Call.Return(run)
	return _c
}

// RawDiff provides a mock function with given fields: ctx, in, opts
func (_m *MockDiffServiceClient) RawDiff(ctx context.Context, in *gitalypb.RawDiffRequest, opts ...grpc.CallOption) (gitalypb.DiffService_RawDiffClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RawDiff")
	}

	var r0 gitalypb.DiffService_RawDiffClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.RawDiffRequest, ...grpc.CallOption) (gitalypb.DiffService_RawDiffClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.RawDiffRequest, ...grpc.CallOption) gitalypb.DiffService_RawDiffClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gitalypb.DiffService_RawDiffClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gitalypb.RawDiffRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDiffServiceClient_RawDiff_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RawDiff'
type MockDiffServiceClient_RawDiff_Call struct {
	*mock.Call
}

// RawDiff is a helper method to define mock.On call
//   - ctx context.Context
//   - in *gitalypb.RawDiffRequest
//   - opts ...grpc.CallOption
func (_e *MockDiffServiceClient_Expecter) RawDiff(ctx interface{}, in interface{}, opts ...interface{}) *MockDiffServiceClient_RawDiff_Call {
	return &MockDiffServiceClient_RawDiff_Call{Call: _e.mock.On("RawDiff",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockDiffServiceClient_RawDiff_Call) Run(run func(ctx context.Context, in *gitalypb.RawDiffRequest, opts ...grpc.CallOption)) *MockDiffServiceClient_RawDiff_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*gitalypb.RawDiffRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockDiffServiceClient_RawDiff_Call) Return(_a0 gitalypb.DiffService_RawDiffClient, _a1 error) *MockDiffServiceClient_RawDiff_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDiffServiceClient_RawDiff_Call) RunAndReturn(run func(context.Context, *gitalypb.RawDiffRequest, ...grpc.CallOption) (gitalypb.DiffService_RawDiffClient, error)) *MockDiffServiceClient_RawDiff_Call {
	_c.Call.Return(run)
	return _c
}

// RawPatch provides a mock function with given fields: ctx, in, opts
func (_m *MockDiffServiceClient) RawPatch(ctx context.Context, in *gitalypb.RawPatchRequest, opts ...grpc.CallOption) (gitalypb.DiffService_RawPatchClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RawPatch")
	}

	var r0 gitalypb.DiffService_RawPatchClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.RawPatchRequest, ...grpc.CallOption) (gitalypb.DiffService_RawPatchClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.RawPatchRequest, ...grpc.CallOption) gitalypb.DiffService_RawPatchClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gitalypb.DiffService_RawPatchClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gitalypb.RawPatchRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDiffServiceClient_RawPatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RawPatch'
type MockDiffServiceClient_RawPatch_Call struct {
	*mock.Call
}

// RawPatch is a helper method to define mock.On call
//   - ctx context.Context
//   - in *gitalypb.RawPatchRequest
//   - opts ...grpc.CallOption
func (_e *MockDiffServiceClient_Expecter) RawPatch(ctx interface{}, in interface{}, opts ...interface{}) *MockDiffServiceClient_RawPatch_Call {
	return &MockDiffServiceClient_RawPatch_Call{Call: _e.mock.On("RawPatch",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockDiffServiceClient_RawPatch_Call) Run(run func(ctx context.Context, in *gitalypb.RawPatchRequest, opts ...grpc.CallOption)) *MockDiffServiceClient_RawPatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*gitalypb.RawPatchRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockDiffServiceClient_RawPatch_Call) Return(_a0 gitalypb.DiffService_RawPatchClient, _a1 error) *MockDiffServiceClient_RawPatch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDiffServiceClient_RawPatch_Call) RunAndReturn(run func(context.Context, *gitalypb.RawPatchRequest, ...grpc.CallOption) (gitalypb.DiffService_RawPatchClient, error)) *MockDiffServiceClient_RawPatch_Call {
	_c.Call.Return(run)
	return _c
}

// RawRangeDiff provides a mock function with given fields: ctx, in, opts
func (_m *MockDiffServiceClient) RawRangeDiff(ctx context.Context, in *gitalypb.RawRangeDiffRequest, opts ...grpc.CallOption) (gitalypb.DiffService_RawRangeDiffClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RawRangeDiff")
	}

	var r0 gitalypb.DiffService_RawRangeDiffClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.RawRangeDiffRequest, ...grpc.CallOption) (gitalypb.DiffService_RawRangeDiffClient, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gitalypb.RawRangeDiffRequest, ...grpc.CallOption) gitalypb.DiffService_RawRangeDiffClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gitalypb.DiffService_RawRangeDiffClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gitalypb.RawRangeDiffRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDiffServiceClient_RawRangeDiff_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RawRangeDiff'
type MockDiffServiceClient_RawRangeDiff_Call struct {
	*mock.Call
}

// RawRangeDiff is a helper method to define mock.On call
//   - ctx context.Context
//   - in *gitalypb.RawRangeDiffRequest
//   - opts ...grpc.CallOption
func (_e *MockDiffServiceClient_Expecter) RawRangeDiff(ctx interface{}, in interface{}, opts ...interface{}) *MockDiffServiceClient_RawRangeDiff_Call {
	return &MockDiffServiceClient_RawRangeDiff_Call{Call: _e.mock.On("RawRangeDiff",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockDiffServiceClient_RawRangeDiff_Call) Run(run func(ctx context.Context, in *gitalypb.RawRangeDiffRequest, opts ...grpc.CallOption)) *MockDiffServiceClient_RawRangeDiff_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*gitalypb.RawRangeDiffRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockDiffServiceClient_RawRangeDiff_Call) Return(_a0 gitalypb.DiffService_RawRangeDiffClient, _a1 error) *MockDiffServiceClient_RawRangeDiff_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDiffServiceClient_RawRangeDiff_Call) RunAndReturn(run func(context.Context, *gitalypb.RawRangeDiffRequest, ...grpc.CallOption) (gitalypb.DiffService_RawRangeDiffClient, error)) *MockDiffServiceClient_RawRangeDiff_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDiffServiceClient creates a new instance of MockDiffServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDiffServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDiffServiceClient {
	mock := &MockDiffServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
