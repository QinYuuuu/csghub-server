// Code generated by mockery v2.49.1. DO NOT EDIT.

package database

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	database "opencsg.com/csghub-server/builder/store/database"
)

// MockRepoRelationsStore is an autogenerated mock type for the RepoRelationsStore type
type MockRepoRelationsStore struct {
	mock.Mock
}

type MockRepoRelationsStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepoRelationsStore) EXPECT() *MockRepoRelationsStore_Expecter {
	return &MockRepoRelationsStore_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, from, to
func (_m *MockRepoRelationsStore) Delete(ctx context.Context, from int64, to int64) error {
	ret := _m.Called(ctx, from, to)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, from, to)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepoRelationsStore_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockRepoRelationsStore_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - from int64
//   - to int64
func (_e *MockRepoRelationsStore_Expecter) Delete(ctx interface{}, from interface{}, to interface{}) *MockRepoRelationsStore_Delete_Call {
	return &MockRepoRelationsStore_Delete_Call{Call: _e.mock.On("Delete", ctx, from, to)}
}

func (_c *MockRepoRelationsStore_Delete_Call) Run(run func(ctx context.Context, from int64, to int64)) *MockRepoRelationsStore_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int64))
	})
	return _c
}

func (_c *MockRepoRelationsStore_Delete_Call) Return(_a0 error) *MockRepoRelationsStore_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepoRelationsStore_Delete_Call) RunAndReturn(run func(context.Context, int64, int64) error) *MockRepoRelationsStore_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// From provides a mock function with given fields: ctx, repoID
func (_m *MockRepoRelationsStore) From(ctx context.Context, repoID int64) ([]*database.RepoRelation, error) {
	ret := _m.Called(ctx, repoID)

	if len(ret) == 0 {
		panic("no return value specified for From")
	}

	var r0 []*database.RepoRelation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*database.RepoRelation, error)); ok {
		return rf(ctx, repoID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*database.RepoRelation); ok {
		r0 = rf(ctx, repoID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*database.RepoRelation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, repoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepoRelationsStore_From_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'From'
type MockRepoRelationsStore_From_Call struct {
	*mock.Call
}

// From is a helper method to define mock.On call
//   - ctx context.Context
//   - repoID int64
func (_e *MockRepoRelationsStore_Expecter) From(ctx interface{}, repoID interface{}) *MockRepoRelationsStore_From_Call {
	return &MockRepoRelationsStore_From_Call{Call: _e.mock.On("From", ctx, repoID)}
}

func (_c *MockRepoRelationsStore_From_Call) Run(run func(ctx context.Context, repoID int64)) *MockRepoRelationsStore_From_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockRepoRelationsStore_From_Call) Return(_a0 []*database.RepoRelation, _a1 error) *MockRepoRelationsStore_From_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepoRelationsStore_From_Call) RunAndReturn(run func(context.Context, int64) ([]*database.RepoRelation, error)) *MockRepoRelationsStore_From_Call {
	_c.Call.Return(run)
	return _c
}

// Override provides a mock function with given fields: ctx, from, to
func (_m *MockRepoRelationsStore) Override(ctx context.Context, from int64, to ...int64) error {
	_va := make([]interface{}, len(to))
	for _i := range to {
		_va[_i] = to[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, from)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Override")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...int64) error); ok {
		r0 = rf(ctx, from, to...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepoRelationsStore_Override_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Override'
type MockRepoRelationsStore_Override_Call struct {
	*mock.Call
}

// Override is a helper method to define mock.On call
//   - ctx context.Context
//   - from int64
//   - to ...int64
func (_e *MockRepoRelationsStore_Expecter) Override(ctx interface{}, from interface{}, to ...interface{}) *MockRepoRelationsStore_Override_Call {
	return &MockRepoRelationsStore_Override_Call{Call: _e.mock.On("Override",
		append([]interface{}{ctx, from}, to...)...)}
}

func (_c *MockRepoRelationsStore_Override_Call) Run(run func(ctx context.Context, from int64, to ...int64)) *MockRepoRelationsStore_Override_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int64, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(int64)
			}
		}
		run(args[0].(context.Context), args[1].(int64), variadicArgs...)
	})
	return _c
}

func (_c *MockRepoRelationsStore_Override_Call) Return(_a0 error) *MockRepoRelationsStore_Override_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepoRelationsStore_Override_Call) RunAndReturn(run func(context.Context, int64, ...int64) error) *MockRepoRelationsStore_Override_Call {
	_c.Call.Return(run)
	return _c
}

// To provides a mock function with given fields: ctx, repoID
func (_m *MockRepoRelationsStore) To(ctx context.Context, repoID int64) ([]*database.RepoRelation, error) {
	ret := _m.Called(ctx, repoID)

	if len(ret) == 0 {
		panic("no return value specified for To")
	}

	var r0 []*database.RepoRelation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*database.RepoRelation, error)); ok {
		return rf(ctx, repoID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*database.RepoRelation); ok {
		r0 = rf(ctx, repoID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*database.RepoRelation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, repoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepoRelationsStore_To_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'To'
type MockRepoRelationsStore_To_Call struct {
	*mock.Call
}

// To is a helper method to define mock.On call
//   - ctx context.Context
//   - repoID int64
func (_e *MockRepoRelationsStore_Expecter) To(ctx interface{}, repoID interface{}) *MockRepoRelationsStore_To_Call {
	return &MockRepoRelationsStore_To_Call{Call: _e.mock.On("To", ctx, repoID)}
}

func (_c *MockRepoRelationsStore_To_Call) Run(run func(ctx context.Context, repoID int64)) *MockRepoRelationsStore_To_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockRepoRelationsStore_To_Call) Return(_a0 []*database.RepoRelation, _a1 error) *MockRepoRelationsStore_To_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepoRelationsStore_To_Call) RunAndReturn(run func(context.Context, int64) ([]*database.RepoRelation, error)) *MockRepoRelationsStore_To_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepoRelationsStore creates a new instance of MockRepoRelationsStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepoRelationsStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepoRelationsStore {
	mock := &MockRepoRelationsStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
