// Code generated by mockery v2.53.0. DO NOT EDIT.

package parquet

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "opencsg.com/csghub-server/common/types"
)

// MockReader is an autogenerated mock type for the Reader type
type MockReader struct {
	mock.Mock
}

type MockReader_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReader) EXPECT() *MockReader_Expecter {
	return &MockReader_Expecter{mock: &_m.Mock}
}

// FetchRows provides a mock function with given fields: ctx, objNames, req, lfs
func (_m *MockReader) FetchRows(ctx context.Context, objNames []string, req types.QueryReq, lfs bool) ([]string, []string, [][]interface{}, error) {
	ret := _m.Called(ctx, objNames, req, lfs)

	if len(ret) == 0 {
		panic("no return value specified for FetchRows")
	}

	var r0 []string
	var r1 []string
	var r2 [][]interface{}
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, types.QueryReq, bool) ([]string, []string, [][]interface{}, error)); ok {
		return rf(ctx, objNames, req, lfs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, types.QueryReq, bool) []string); ok {
		r0 = rf(ctx, objNames, req, lfs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, types.QueryReq, bool) []string); ok {
		r1 = rf(ctx, objNames, req, lfs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, []string, types.QueryReq, bool) [][]interface{}); ok {
		r2 = rf(ctx, objNames, req, lfs)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([][]interface{})
		}
	}

	if rf, ok := ret.Get(3).(func(context.Context, []string, types.QueryReq, bool) error); ok {
		r3 = rf(ctx, objNames, req, lfs)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// MockReader_FetchRows_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchRows'
type MockReader_FetchRows_Call struct {
	*mock.Call
}

// FetchRows is a helper method to define mock.On call
//   - ctx context.Context
//   - objNames []string
//   - req types.QueryReq
//   - lfs bool
func (_e *MockReader_Expecter) FetchRows(ctx interface{}, objNames interface{}, req interface{}, lfs interface{}) *MockReader_FetchRows_Call {
	return &MockReader_FetchRows_Call{Call: _e.mock.On("FetchRows", ctx, objNames, req, lfs)}
}

func (_c *MockReader_FetchRows_Call) Run(run func(ctx context.Context, objNames []string, req types.QueryReq, lfs bool)) *MockReader_FetchRows_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string), args[2].(types.QueryReq), args[3].(bool))
	})
	return _c
}

func (_c *MockReader_FetchRows_Call) Return(columns []string, columnsType []string, rows [][]interface{}, err error) *MockReader_FetchRows_Call {
	_c.Call.Return(columns, columnsType, rows, err)
	return _c
}

func (_c *MockReader_FetchRows_Call) RunAndReturn(run func(context.Context, []string, types.QueryReq, bool) ([]string, []string, [][]interface{}, error)) *MockReader_FetchRows_Call {
	_c.Call.Return(run)
	return _c
}

// RowCount provides a mock function with given fields: ctx, objNames, req, lfs
func (_m *MockReader) RowCount(ctx context.Context, objNames []string, req types.QueryReq, lfs bool) (int, error) {
	ret := _m.Called(ctx, objNames, req, lfs)

	if len(ret) == 0 {
		panic("no return value specified for RowCount")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, types.QueryReq, bool) (int, error)); ok {
		return rf(ctx, objNames, req, lfs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, types.QueryReq, bool) int); ok {
		r0 = rf(ctx, objNames, req, lfs)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, types.QueryReq, bool) error); ok {
		r1 = rf(ctx, objNames, req, lfs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReader_RowCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RowCount'
type MockReader_RowCount_Call struct {
	*mock.Call
}

// RowCount is a helper method to define mock.On call
//   - ctx context.Context
//   - objNames []string
//   - req types.QueryReq
//   - lfs bool
func (_e *MockReader_Expecter) RowCount(ctx interface{}, objNames interface{}, req interface{}, lfs interface{}) *MockReader_RowCount_Call {
	return &MockReader_RowCount_Call{Call: _e.mock.On("RowCount", ctx, objNames, req, lfs)}
}

func (_c *MockReader_RowCount_Call) Run(run func(ctx context.Context, objNames []string, req types.QueryReq, lfs bool)) *MockReader_RowCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string), args[2].(types.QueryReq), args[3].(bool))
	})
	return _c
}

func (_c *MockReader_RowCount_Call) Return(count int, err error) *MockReader_RowCount_Call {
	_c.Call.Return(count, err)
	return _c
}

func (_c *MockReader_RowCount_Call) RunAndReturn(run func(context.Context, []string, types.QueryReq, bool) (int, error)) *MockReader_RowCount_Call {
	_c.Call.Return(run)
	return _c
}

// TopN provides a mock function with given fields: ctx, objName, count
func (_m *MockReader) TopN(ctx context.Context, objName string, count int) ([]string, []string, [][]interface{}, error) {
	ret := _m.Called(ctx, objName, count)

	if len(ret) == 0 {
		panic("no return value specified for TopN")
	}

	var r0 []string
	var r1 []string
	var r2 [][]interface{}
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) ([]string, []string, [][]interface{}, error)); ok {
		return rf(ctx, objName, count)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int) []string); ok {
		r0 = rf(ctx, objName, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int) []string); ok {
		r1 = rf(ctx, objName, count)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, int) [][]interface{}); ok {
		r2 = rf(ctx, objName, count)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([][]interface{})
		}
	}

	if rf, ok := ret.Get(3).(func(context.Context, string, int) error); ok {
		r3 = rf(ctx, objName, count)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// MockReader_TopN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TopN'
type MockReader_TopN_Call struct {
	*mock.Call
}

// TopN is a helper method to define mock.On call
//   - ctx context.Context
//   - objName string
//   - count int
func (_e *MockReader_Expecter) TopN(ctx interface{}, objName interface{}, count interface{}) *MockReader_TopN_Call {
	return &MockReader_TopN_Call{Call: _e.mock.On("TopN", ctx, objName, count)}
}

func (_c *MockReader_TopN_Call) Run(run func(ctx context.Context, objName string, count int)) *MockReader_TopN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int))
	})
	return _c
}

func (_c *MockReader_TopN_Call) Return(columns []string, columnsType []string, rows [][]interface{}, err error) *MockReader_TopN_Call {
	_c.Call.Return(columns, columnsType, rows, err)
	return _c
}

func (_c *MockReader_TopN_Call) RunAndReturn(run func(context.Context, string, int) ([]string, []string, [][]interface{}, error)) *MockReader_TopN_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockReader creates a new instance of MockReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReader {
	mock := &MockReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
