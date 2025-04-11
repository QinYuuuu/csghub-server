// Code generated by mockery v2.53.0. DO NOT EDIT.

package component

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	database "opencsg.com/csghub-server/builder/store/database"

	types "opencsg.com/csghub-server/common/types"
)

// MockRuntimeArchitectureComponent is an autogenerated mock type for the RuntimeArchitectureComponent type
type MockRuntimeArchitectureComponent struct {
	mock.Mock
}

type MockRuntimeArchitectureComponent_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRuntimeArchitectureComponent) EXPECT() *MockRuntimeArchitectureComponent_Expecter {
	return &MockRuntimeArchitectureComponent_Expecter{mock: &_m.Mock}
}

// AddResourceTag provides a mock function with given fields: ctx, rstags, modelname, repoId
func (_m *MockRuntimeArchitectureComponent) AddResourceTag(ctx context.Context, rstags []*database.Tag, modelname string, repoId int64) error {
	ret := _m.Called(ctx, rstags, modelname, repoId)

	if len(ret) == 0 {
		panic("no return value specified for AddResourceTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*database.Tag, string, int64) error); ok {
		r0 = rf(ctx, rstags, modelname, repoId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRuntimeArchitectureComponent_AddResourceTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddResourceTag'
type MockRuntimeArchitectureComponent_AddResourceTag_Call struct {
	*mock.Call
}

// AddResourceTag is a helper method to define mock.On call
//   - ctx context.Context
//   - rstags []*database.Tag
//   - modelname string
//   - repoId int64
func (_e *MockRuntimeArchitectureComponent_Expecter) AddResourceTag(ctx interface{}, rstags interface{}, modelname interface{}, repoId interface{}) *MockRuntimeArchitectureComponent_AddResourceTag_Call {
	return &MockRuntimeArchitectureComponent_AddResourceTag_Call{Call: _e.mock.On("AddResourceTag", ctx, rstags, modelname, repoId)}
}

func (_c *MockRuntimeArchitectureComponent_AddResourceTag_Call) Run(run func(ctx context.Context, rstags []*database.Tag, modelname string, repoId int64)) *MockRuntimeArchitectureComponent_AddResourceTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*database.Tag), args[2].(string), args[3].(int64))
	})
	return _c
}

func (_c *MockRuntimeArchitectureComponent_AddResourceTag_Call) Return(_a0 error) *MockRuntimeArchitectureComponent_AddResourceTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRuntimeArchitectureComponent_AddResourceTag_Call) RunAndReturn(run func(context.Context, []*database.Tag, string, int64) error) *MockRuntimeArchitectureComponent_AddResourceTag_Call {
	_c.Call.Return(run)
	return _c
}

// AddRuntimeFrameworkTag provides a mock function with given fields: ctx, rftags, repoId, rfId
func (_m *MockRuntimeArchitectureComponent) AddRuntimeFrameworkTag(ctx context.Context, rftags []*database.Tag, repoId int64, rfId int64) error {
	ret := _m.Called(ctx, rftags, repoId, rfId)

	if len(ret) == 0 {
		panic("no return value specified for AddRuntimeFrameworkTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*database.Tag, int64, int64) error); ok {
		r0 = rf(ctx, rftags, repoId, rfId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRuntimeArchitectureComponent_AddRuntimeFrameworkTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddRuntimeFrameworkTag'
type MockRuntimeArchitectureComponent_AddRuntimeFrameworkTag_Call struct {
	*mock.Call
}

// AddRuntimeFrameworkTag is a helper method to define mock.On call
//   - ctx context.Context
//   - rftags []*database.Tag
//   - repoId int64
//   - rfId int64
func (_e *MockRuntimeArchitectureComponent_Expecter) AddRuntimeFrameworkTag(ctx interface{}, rftags interface{}, repoId interface{}, rfId interface{}) *MockRuntimeArchitectureComponent_AddRuntimeFrameworkTag_Call {
	return &MockRuntimeArchitectureComponent_AddRuntimeFrameworkTag_Call{Call: _e.mock.On("AddRuntimeFrameworkTag", ctx, rftags, repoId, rfId)}
}

func (_c *MockRuntimeArchitectureComponent_AddRuntimeFrameworkTag_Call) Run(run func(ctx context.Context, rftags []*database.Tag, repoId int64, rfId int64)) *MockRuntimeArchitectureComponent_AddRuntimeFrameworkTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*database.Tag), args[2].(int64), args[3].(int64))
	})
	return _c
}

func (_c *MockRuntimeArchitectureComponent_AddRuntimeFrameworkTag_Call) Return(_a0 error) *MockRuntimeArchitectureComponent_AddRuntimeFrameworkTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRuntimeArchitectureComponent_AddRuntimeFrameworkTag_Call) RunAndReturn(run func(context.Context, []*database.Tag, int64, int64) error) *MockRuntimeArchitectureComponent_AddRuntimeFrameworkTag_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteArchitectures provides a mock function with given fields: ctx, id, architectures
func (_m *MockRuntimeArchitectureComponent) DeleteArchitectures(ctx context.Context, id int64, architectures []string) ([]string, error) {
	ret := _m.Called(ctx, id, architectures)

	if len(ret) == 0 {
		panic("no return value specified for DeleteArchitectures")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []string) ([]string, error)); ok {
		return rf(ctx, id, architectures)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, []string) []string); ok {
		r0 = rf(ctx, id, architectures)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, []string) error); ok {
		r1 = rf(ctx, id, architectures)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRuntimeArchitectureComponent_DeleteArchitectures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteArchitectures'
type MockRuntimeArchitectureComponent_DeleteArchitectures_Call struct {
	*mock.Call
}

// DeleteArchitectures is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
//   - architectures []string
func (_e *MockRuntimeArchitectureComponent_Expecter) DeleteArchitectures(ctx interface{}, id interface{}, architectures interface{}) *MockRuntimeArchitectureComponent_DeleteArchitectures_Call {
	return &MockRuntimeArchitectureComponent_DeleteArchitectures_Call{Call: _e.mock.On("DeleteArchitectures", ctx, id, architectures)}
}

func (_c *MockRuntimeArchitectureComponent_DeleteArchitectures_Call) Run(run func(ctx context.Context, id int64, architectures []string)) *MockRuntimeArchitectureComponent_DeleteArchitectures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].([]string))
	})
	return _c
}

func (_c *MockRuntimeArchitectureComponent_DeleteArchitectures_Call) Return(_a0 []string, _a1 error) *MockRuntimeArchitectureComponent_DeleteArchitectures_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRuntimeArchitectureComponent_DeleteArchitectures_Call) RunAndReturn(run func(context.Context, int64, []string) ([]string, error)) *MockRuntimeArchitectureComponent_DeleteArchitectures_Call {
	_c.Call.Return(run)
	return _c
}

// InitRuntimeFrameworkAndArchitectures provides a mock function with no fields
func (_m *MockRuntimeArchitectureComponent) InitRuntimeFrameworkAndArchitectures() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for InitRuntimeFrameworkAndArchitectures")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRuntimeArchitectureComponent_InitRuntimeFrameworkAndArchitectures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitRuntimeFrameworkAndArchitectures'
type MockRuntimeArchitectureComponent_InitRuntimeFrameworkAndArchitectures_Call struct {
	*mock.Call
}

// InitRuntimeFrameworkAndArchitectures is a helper method to define mock.On call
func (_e *MockRuntimeArchitectureComponent_Expecter) InitRuntimeFrameworkAndArchitectures() *MockRuntimeArchitectureComponent_InitRuntimeFrameworkAndArchitectures_Call {
	return &MockRuntimeArchitectureComponent_InitRuntimeFrameworkAndArchitectures_Call{Call: _e.mock.On("InitRuntimeFrameworkAndArchitectures")}
}

func (_c *MockRuntimeArchitectureComponent_InitRuntimeFrameworkAndArchitectures_Call) Run(run func()) *MockRuntimeArchitectureComponent_InitRuntimeFrameworkAndArchitectures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRuntimeArchitectureComponent_InitRuntimeFrameworkAndArchitectures_Call) Return(_a0 error) *MockRuntimeArchitectureComponent_InitRuntimeFrameworkAndArchitectures_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRuntimeArchitectureComponent_InitRuntimeFrameworkAndArchitectures_Call) RunAndReturn(run func() error) *MockRuntimeArchitectureComponent_InitRuntimeFrameworkAndArchitectures_Call {
	_c.Call.Return(run)
	return _c
}

// ListByRuntimeFrameworkID provides a mock function with given fields: ctx, id
func (_m *MockRuntimeArchitectureComponent) ListByRuntimeFrameworkID(ctx context.Context, id int64) ([]database.RuntimeArchitecture, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for ListByRuntimeFrameworkID")
	}

	var r0 []database.RuntimeArchitecture
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]database.RuntimeArchitecture, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []database.RuntimeArchitecture); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.RuntimeArchitecture)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRuntimeArchitectureComponent_ListByRuntimeFrameworkID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByRuntimeFrameworkID'
type MockRuntimeArchitectureComponent_ListByRuntimeFrameworkID_Call struct {
	*mock.Call
}

// ListByRuntimeFrameworkID is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockRuntimeArchitectureComponent_Expecter) ListByRuntimeFrameworkID(ctx interface{}, id interface{}) *MockRuntimeArchitectureComponent_ListByRuntimeFrameworkID_Call {
	return &MockRuntimeArchitectureComponent_ListByRuntimeFrameworkID_Call{Call: _e.mock.On("ListByRuntimeFrameworkID", ctx, id)}
}

func (_c *MockRuntimeArchitectureComponent_ListByRuntimeFrameworkID_Call) Run(run func(ctx context.Context, id int64)) *MockRuntimeArchitectureComponent_ListByRuntimeFrameworkID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockRuntimeArchitectureComponent_ListByRuntimeFrameworkID_Call) Return(_a0 []database.RuntimeArchitecture, _a1 error) *MockRuntimeArchitectureComponent_ListByRuntimeFrameworkID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRuntimeArchitectureComponent_ListByRuntimeFrameworkID_Call) RunAndReturn(run func(context.Context, int64) ([]database.RuntimeArchitecture, error)) *MockRuntimeArchitectureComponent_ListByRuntimeFrameworkID_Call {
	_c.Call.Return(run)
	return _c
}

// ScanAllModels provides a mock function with given fields: ctx, scanType
func (_m *MockRuntimeArchitectureComponent) ScanAllModels(ctx context.Context, scanType int) error {
	ret := _m.Called(ctx, scanType)

	if len(ret) == 0 {
		panic("no return value specified for ScanAllModels")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, scanType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRuntimeArchitectureComponent_ScanAllModels_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ScanAllModels'
type MockRuntimeArchitectureComponent_ScanAllModels_Call struct {
	*mock.Call
}

// ScanAllModels is a helper method to define mock.On call
//   - ctx context.Context
//   - scanType int
func (_e *MockRuntimeArchitectureComponent_Expecter) ScanAllModels(ctx interface{}, scanType interface{}) *MockRuntimeArchitectureComponent_ScanAllModels_Call {
	return &MockRuntimeArchitectureComponent_ScanAllModels_Call{Call: _e.mock.On("ScanAllModels", ctx, scanType)}
}

func (_c *MockRuntimeArchitectureComponent_ScanAllModels_Call) Run(run func(ctx context.Context, scanType int)) *MockRuntimeArchitectureComponent_ScanAllModels_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockRuntimeArchitectureComponent_ScanAllModels_Call) Return(_a0 error) *MockRuntimeArchitectureComponent_ScanAllModels_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRuntimeArchitectureComponent_ScanAllModels_Call) RunAndReturn(run func(context.Context, int) error) *MockRuntimeArchitectureComponent_ScanAllModels_Call {
	_c.Call.Return(run)
	return _c
}

// SetArchitectures provides a mock function with given fields: ctx, id, architectures
func (_m *MockRuntimeArchitectureComponent) SetArchitectures(ctx context.Context, id int64, architectures []string) ([]string, error) {
	ret := _m.Called(ctx, id, architectures)

	if len(ret) == 0 {
		panic("no return value specified for SetArchitectures")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []string) ([]string, error)); ok {
		return rf(ctx, id, architectures)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, []string) []string); ok {
		r0 = rf(ctx, id, architectures)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, []string) error); ok {
		r1 = rf(ctx, id, architectures)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRuntimeArchitectureComponent_SetArchitectures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetArchitectures'
type MockRuntimeArchitectureComponent_SetArchitectures_Call struct {
	*mock.Call
}

// SetArchitectures is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
//   - architectures []string
func (_e *MockRuntimeArchitectureComponent_Expecter) SetArchitectures(ctx interface{}, id interface{}, architectures interface{}) *MockRuntimeArchitectureComponent_SetArchitectures_Call {
	return &MockRuntimeArchitectureComponent_SetArchitectures_Call{Call: _e.mock.On("SetArchitectures", ctx, id, architectures)}
}

func (_c *MockRuntimeArchitectureComponent_SetArchitectures_Call) Run(run func(ctx context.Context, id int64, architectures []string)) *MockRuntimeArchitectureComponent_SetArchitectures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].([]string))
	})
	return _c
}

func (_c *MockRuntimeArchitectureComponent_SetArchitectures_Call) Return(_a0 []string, _a1 error) *MockRuntimeArchitectureComponent_SetArchitectures_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRuntimeArchitectureComponent_SetArchitectures_Call) RunAndReturn(run func(context.Context, int64, []string) ([]string, error)) *MockRuntimeArchitectureComponent_SetArchitectures_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateModelMetadata provides a mock function with given fields: ctx, repo
func (_m *MockRuntimeArchitectureComponent) UpdateModelMetadata(ctx context.Context, repo *database.Repository) (*types.ModelInfo, error) {
	ret := _m.Called(ctx, repo)

	if len(ret) == 0 {
		panic("no return value specified for UpdateModelMetadata")
	}

	var r0 *types.ModelInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *database.Repository) (*types.ModelInfo, error)); ok {
		return rf(ctx, repo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *database.Repository) *types.ModelInfo); ok {
		r0 = rf(ctx, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ModelInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *database.Repository) error); ok {
		r1 = rf(ctx, repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRuntimeArchitectureComponent_UpdateModelMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateModelMetadata'
type MockRuntimeArchitectureComponent_UpdateModelMetadata_Call struct {
	*mock.Call
}

// UpdateModelMetadata is a helper method to define mock.On call
//   - ctx context.Context
//   - repo *database.Repository
func (_e *MockRuntimeArchitectureComponent_Expecter) UpdateModelMetadata(ctx interface{}, repo interface{}) *MockRuntimeArchitectureComponent_UpdateModelMetadata_Call {
	return &MockRuntimeArchitectureComponent_UpdateModelMetadata_Call{Call: _e.mock.On("UpdateModelMetadata", ctx, repo)}
}

func (_c *MockRuntimeArchitectureComponent_UpdateModelMetadata_Call) Run(run func(ctx context.Context, repo *database.Repository)) *MockRuntimeArchitectureComponent_UpdateModelMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*database.Repository))
	})
	return _c
}

func (_c *MockRuntimeArchitectureComponent_UpdateModelMetadata_Call) Return(_a0 *types.ModelInfo, _a1 error) *MockRuntimeArchitectureComponent_UpdateModelMetadata_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRuntimeArchitectureComponent_UpdateModelMetadata_Call) RunAndReturn(run func(context.Context, *database.Repository) (*types.ModelInfo, error)) *MockRuntimeArchitectureComponent_UpdateModelMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRuntimeFrameworkTag provides a mock function with given fields: ctx, modelInfo, repo
func (_m *MockRuntimeArchitectureComponent) UpdateRuntimeFrameworkTag(ctx context.Context, modelInfo *types.ModelInfo, repo *database.Repository) error {
	ret := _m.Called(ctx, modelInfo, repo)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRuntimeFrameworkTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.ModelInfo, *database.Repository) error); ok {
		r0 = rf(ctx, modelInfo, repo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRuntimeArchitectureComponent_UpdateRuntimeFrameworkTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRuntimeFrameworkTag'
type MockRuntimeArchitectureComponent_UpdateRuntimeFrameworkTag_Call struct {
	*mock.Call
}

// UpdateRuntimeFrameworkTag is a helper method to define mock.On call
//   - ctx context.Context
//   - modelInfo *types.ModelInfo
//   - repo *database.Repository
func (_e *MockRuntimeArchitectureComponent_Expecter) UpdateRuntimeFrameworkTag(ctx interface{}, modelInfo interface{}, repo interface{}) *MockRuntimeArchitectureComponent_UpdateRuntimeFrameworkTag_Call {
	return &MockRuntimeArchitectureComponent_UpdateRuntimeFrameworkTag_Call{Call: _e.mock.On("UpdateRuntimeFrameworkTag", ctx, modelInfo, repo)}
}

func (_c *MockRuntimeArchitectureComponent_UpdateRuntimeFrameworkTag_Call) Run(run func(ctx context.Context, modelInfo *types.ModelInfo, repo *database.Repository)) *MockRuntimeArchitectureComponent_UpdateRuntimeFrameworkTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.ModelInfo), args[2].(*database.Repository))
	})
	return _c
}

func (_c *MockRuntimeArchitectureComponent_UpdateRuntimeFrameworkTag_Call) Return(_a0 error) *MockRuntimeArchitectureComponent_UpdateRuntimeFrameworkTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRuntimeArchitectureComponent_UpdateRuntimeFrameworkTag_Call) RunAndReturn(run func(context.Context, *types.ModelInfo, *database.Repository) error) *MockRuntimeArchitectureComponent_UpdateRuntimeFrameworkTag_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRuntimeArchitectureComponent creates a new instance of MockRuntimeArchitectureComponent. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRuntimeArchitectureComponent(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRuntimeArchitectureComponent {
	mock := &MockRuntimeArchitectureComponent{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
