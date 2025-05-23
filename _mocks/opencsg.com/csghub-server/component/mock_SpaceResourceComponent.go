// Code generated by mockery v2.53.0. DO NOT EDIT.

package component

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	types "opencsg.com/csghub-server/common/types"
)

// MockSpaceResourceComponent is an autogenerated mock type for the SpaceResourceComponent type
type MockSpaceResourceComponent struct {
	mock.Mock
}

type MockSpaceResourceComponent_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSpaceResourceComponent) EXPECT() *MockSpaceResourceComponent_Expecter {
	return &MockSpaceResourceComponent_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, req
func (_m *MockSpaceResourceComponent) Create(ctx context.Context, req *types.CreateSpaceResourceReq) (*types.SpaceResource, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *types.SpaceResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.CreateSpaceResourceReq) (*types.SpaceResource, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.CreateSpaceResourceReq) *types.SpaceResource); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.SpaceResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.CreateSpaceResourceReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceResourceComponent_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockSpaceResourceComponent_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - req *types.CreateSpaceResourceReq
func (_e *MockSpaceResourceComponent_Expecter) Create(ctx interface{}, req interface{}) *MockSpaceResourceComponent_Create_Call {
	return &MockSpaceResourceComponent_Create_Call{Call: _e.mock.On("Create", ctx, req)}
}

func (_c *MockSpaceResourceComponent_Create_Call) Run(run func(ctx context.Context, req *types.CreateSpaceResourceReq)) *MockSpaceResourceComponent_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.CreateSpaceResourceReq))
	})
	return _c
}

func (_c *MockSpaceResourceComponent_Create_Call) Return(_a0 *types.SpaceResource, _a1 error) *MockSpaceResourceComponent_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSpaceResourceComponent_Create_Call) RunAndReturn(run func(context.Context, *types.CreateSpaceResourceReq) (*types.SpaceResource, error)) *MockSpaceResourceComponent_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockSpaceResourceComponent) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSpaceResourceComponent_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockSpaceResourceComponent_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockSpaceResourceComponent_Expecter) Delete(ctx interface{}, id interface{}) *MockSpaceResourceComponent_Delete_Call {
	return &MockSpaceResourceComponent_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *MockSpaceResourceComponent_Delete_Call) Run(run func(ctx context.Context, id int64)) *MockSpaceResourceComponent_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockSpaceResourceComponent_Delete_Call) Return(_a0 error) *MockSpaceResourceComponent_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpaceResourceComponent_Delete_Call) RunAndReturn(run func(context.Context, int64) error) *MockSpaceResourceComponent_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Index provides a mock function with given fields: ctx, clusterId, deployType, currentUser
func (_m *MockSpaceResourceComponent) Index(ctx context.Context, clusterId string, deployType int, currentUser string) ([]types.SpaceResource, error) {
	ret := _m.Called(ctx, clusterId, deployType, currentUser)

	if len(ret) == 0 {
		panic("no return value specified for Index")
	}

	var r0 []types.SpaceResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, string) ([]types.SpaceResource, error)); ok {
		return rf(ctx, clusterId, deployType, currentUser)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, string) []types.SpaceResource); ok {
		r0 = rf(ctx, clusterId, deployType, currentUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.SpaceResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, string) error); ok {
		r1 = rf(ctx, clusterId, deployType, currentUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceResourceComponent_Index_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Index'
type MockSpaceResourceComponent_Index_Call struct {
	*mock.Call
}

// Index is a helper method to define mock.On call
//   - ctx context.Context
//   - clusterId string
//   - deployType int
//   - currentUser string
func (_e *MockSpaceResourceComponent_Expecter) Index(ctx interface{}, clusterId interface{}, deployType interface{}, currentUser interface{}) *MockSpaceResourceComponent_Index_Call {
	return &MockSpaceResourceComponent_Index_Call{Call: _e.mock.On("Index", ctx, clusterId, deployType, currentUser)}
}

func (_c *MockSpaceResourceComponent_Index_Call) Run(run func(ctx context.Context, clusterId string, deployType int, currentUser string)) *MockSpaceResourceComponent_Index_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int), args[3].(string))
	})
	return _c
}

func (_c *MockSpaceResourceComponent_Index_Call) Return(_a0 []types.SpaceResource, _a1 error) *MockSpaceResourceComponent_Index_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSpaceResourceComponent_Index_Call) RunAndReturn(run func(context.Context, string, int, string) ([]types.SpaceResource, error)) *MockSpaceResourceComponent_Index_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, req
func (_m *MockSpaceResourceComponent) Update(ctx context.Context, req *types.UpdateSpaceResourceReq) (*types.SpaceResource, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *types.SpaceResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.UpdateSpaceResourceReq) (*types.SpaceResource, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.UpdateSpaceResourceReq) *types.SpaceResource); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.SpaceResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.UpdateSpaceResourceReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSpaceResourceComponent_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockSpaceResourceComponent_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - req *types.UpdateSpaceResourceReq
func (_e *MockSpaceResourceComponent_Expecter) Update(ctx interface{}, req interface{}) *MockSpaceResourceComponent_Update_Call {
	return &MockSpaceResourceComponent_Update_Call{Call: _e.mock.On("Update", ctx, req)}
}

func (_c *MockSpaceResourceComponent_Update_Call) Run(run func(ctx context.Context, req *types.UpdateSpaceResourceReq)) *MockSpaceResourceComponent_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.UpdateSpaceResourceReq))
	})
	return _c
}

func (_c *MockSpaceResourceComponent_Update_Call) Return(_a0 *types.SpaceResource, _a1 error) *MockSpaceResourceComponent_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSpaceResourceComponent_Update_Call) RunAndReturn(run func(context.Context, *types.UpdateSpaceResourceReq) (*types.SpaceResource, error)) *MockSpaceResourceComponent_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSpaceResourceComponent creates a new instance of MockSpaceResourceComponent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSpaceResourceComponent(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSpaceResourceComponent {
	mock := &MockSpaceResourceComponent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
