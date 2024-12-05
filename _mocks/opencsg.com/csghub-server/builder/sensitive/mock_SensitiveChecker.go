// Code generated by mockery v2.49.1. DO NOT EDIT.

package sensitive

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	sensitive "opencsg.com/csghub-server/builder/sensitive"
)

// MockSensitiveChecker is an autogenerated mock type for the SensitiveChecker type
type MockSensitiveChecker struct {
	mock.Mock
}

type MockSensitiveChecker_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSensitiveChecker) EXPECT() *MockSensitiveChecker_Expecter {
	return &MockSensitiveChecker_Expecter{mock: &_m.Mock}
}

// PassImageCheck provides a mock function with given fields: ctx, scenario, ossBucketName, ossObjectName
func (_m *MockSensitiveChecker) PassImageCheck(ctx context.Context, scenario sensitive.Scenario, ossBucketName string, ossObjectName string) (*sensitive.CheckResult, error) {
	ret := _m.Called(ctx, scenario, ossBucketName, ossObjectName)

	if len(ret) == 0 {
		panic("no return value specified for PassImageCheck")
	}

	var r0 *sensitive.CheckResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sensitive.Scenario, string, string) (*sensitive.CheckResult, error)); ok {
		return rf(ctx, scenario, ossBucketName, ossObjectName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sensitive.Scenario, string, string) *sensitive.CheckResult); ok {
		r0 = rf(ctx, scenario, ossBucketName, ossObjectName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sensitive.CheckResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sensitive.Scenario, string, string) error); ok {
		r1 = rf(ctx, scenario, ossBucketName, ossObjectName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSensitiveChecker_PassImageCheck_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PassImageCheck'
type MockSensitiveChecker_PassImageCheck_Call struct {
	*mock.Call
}

// PassImageCheck is a helper method to define mock.On call
//   - ctx context.Context
//   - scenario sensitive.Scenario
//   - ossBucketName string
//   - ossObjectName string
func (_e *MockSensitiveChecker_Expecter) PassImageCheck(ctx interface{}, scenario interface{}, ossBucketName interface{}, ossObjectName interface{}) *MockSensitiveChecker_PassImageCheck_Call {
	return &MockSensitiveChecker_PassImageCheck_Call{Call: _e.mock.On("PassImageCheck", ctx, scenario, ossBucketName, ossObjectName)}
}

func (_c *MockSensitiveChecker_PassImageCheck_Call) Run(run func(ctx context.Context, scenario sensitive.Scenario, ossBucketName string, ossObjectName string)) *MockSensitiveChecker_PassImageCheck_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sensitive.Scenario), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockSensitiveChecker_PassImageCheck_Call) Return(_a0 *sensitive.CheckResult, _a1 error) *MockSensitiveChecker_PassImageCheck_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSensitiveChecker_PassImageCheck_Call) RunAndReturn(run func(context.Context, sensitive.Scenario, string, string) (*sensitive.CheckResult, error)) *MockSensitiveChecker_PassImageCheck_Call {
	_c.Call.Return(run)
	return _c
}

// PassTextCheck provides a mock function with given fields: ctx, scenario, text
func (_m *MockSensitiveChecker) PassTextCheck(ctx context.Context, scenario sensitive.Scenario, text string) (*sensitive.CheckResult, error) {
	ret := _m.Called(ctx, scenario, text)

	if len(ret) == 0 {
		panic("no return value specified for PassTextCheck")
	}

	var r0 *sensitive.CheckResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sensitive.Scenario, string) (*sensitive.CheckResult, error)); ok {
		return rf(ctx, scenario, text)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sensitive.Scenario, string) *sensitive.CheckResult); ok {
		r0 = rf(ctx, scenario, text)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sensitive.CheckResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sensitive.Scenario, string) error); ok {
		r1 = rf(ctx, scenario, text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSensitiveChecker_PassTextCheck_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PassTextCheck'
type MockSensitiveChecker_PassTextCheck_Call struct {
	*mock.Call
}

// PassTextCheck is a helper method to define mock.On call
//   - ctx context.Context
//   - scenario sensitive.Scenario
//   - text string
func (_e *MockSensitiveChecker_Expecter) PassTextCheck(ctx interface{}, scenario interface{}, text interface{}) *MockSensitiveChecker_PassTextCheck_Call {
	return &MockSensitiveChecker_PassTextCheck_Call{Call: _e.mock.On("PassTextCheck", ctx, scenario, text)}
}

func (_c *MockSensitiveChecker_PassTextCheck_Call) Run(run func(ctx context.Context, scenario sensitive.Scenario, text string)) *MockSensitiveChecker_PassTextCheck_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sensitive.Scenario), args[2].(string))
	})
	return _c
}

func (_c *MockSensitiveChecker_PassTextCheck_Call) Return(_a0 *sensitive.CheckResult, _a1 error) *MockSensitiveChecker_PassTextCheck_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSensitiveChecker_PassTextCheck_Call) RunAndReturn(run func(context.Context, sensitive.Scenario, string) (*sensitive.CheckResult, error)) *MockSensitiveChecker_PassTextCheck_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSensitiveChecker creates a new instance of MockSensitiveChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSensitiveChecker(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSensitiveChecker {
	mock := &MockSensitiveChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
