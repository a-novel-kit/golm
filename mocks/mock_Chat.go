// Code generated by mockery v2.53.1. DO NOT EDIT.

package mocks

import (
	context "context"

	golm "github.com/a-novel-kit/golm"
	mock "github.com/stretchr/testify/mock"

	utils "github.com/a-novel-kit/golm/utils"
)

// MockChat is an autogenerated mock type for the Chat type
type MockChat[RawRequest any, RawResponse any] struct {
	mock.Mock
}

type MockChat_Expecter[RawRequest any, RawResponse any] struct {
	mock *mock.Mock
}

func (_m *MockChat[RawRequest, RawResponse]) EXPECT() *MockChat_Expecter[RawRequest, RawResponse] {
	return &MockChat_Expecter[RawRequest, RawResponse]{mock: &_m.Mock}
}

// Completion provides a mock function with given fields: ctx, message, options
func (_m *MockChat[RawRequest, RawResponse]) Completion(ctx context.Context, message golm.UserMessage, options golm.CompletionParams) (*golm.AssistantMessage, error) {
	ret := _m.Called(ctx, message, options)

	if len(ret) == 0 {
		panic("no return value specified for Completion")
	}

	var r0 *golm.AssistantMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, golm.UserMessage, golm.CompletionParams) (*golm.AssistantMessage, error)); ok {
		return rf(ctx, message, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, golm.UserMessage, golm.CompletionParams) *golm.AssistantMessage); ok {
		r0 = rf(ctx, message, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*golm.AssistantMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, golm.UserMessage, golm.CompletionParams) error); ok {
		r1 = rf(ctx, message, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChat_Completion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Completion'
type MockChat_Completion_Call[RawRequest any, RawResponse any] struct {
	*mock.Call
}

// Completion is a helper method to define mock.On call
//   - ctx context.Context
//   - message golm.UserMessage
//   - options golm.CompletionParams
func (_e *MockChat_Expecter[RawRequest, RawResponse]) Completion(ctx interface{}, message interface{}, options interface{}) *MockChat_Completion_Call[RawRequest, RawResponse] {
	return &MockChat_Completion_Call[RawRequest, RawResponse]{Call: _e.mock.On("Completion", ctx, message, options)}
}

func (_c *MockChat_Completion_Call[RawRequest, RawResponse]) Run(run func(ctx context.Context, message golm.UserMessage, options golm.CompletionParams)) *MockChat_Completion_Call[RawRequest, RawResponse] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(golm.UserMessage), args[2].(golm.CompletionParams))
	})
	return _c
}

func (_c *MockChat_Completion_Call[RawRequest, RawResponse]) Return(response *golm.AssistantMessage, err error) *MockChat_Completion_Call[RawRequest, RawResponse] {
	_c.Call.Return(response, err)
	return _c
}

func (_c *MockChat_Completion_Call[RawRequest, RawResponse]) RunAndReturn(run func(context.Context, golm.UserMessage, golm.CompletionParams) (*golm.AssistantMessage, error)) *MockChat_Completion_Call[RawRequest, RawResponse] {
	_c.Call.Return(run)
	return _c
}

// CompletionJSON provides a mock function with given fields: ctx, message, options, dest
func (_m *MockChat[RawRequest, RawResponse]) CompletionJSON(ctx context.Context, message golm.UserMessage, options golm.CompletionParams, dest any) error {
	ret := _m.Called(ctx, message, options, dest)

	if len(ret) == 0 {
		panic("no return value specified for CompletionJSON")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, golm.UserMessage, golm.CompletionParams, any) error); ok {
		r0 = rf(ctx, message, options, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockChat_CompletionJSON_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompletionJSON'
type MockChat_CompletionJSON_Call[RawRequest any, RawResponse any] struct {
	*mock.Call
}

// CompletionJSON is a helper method to define mock.On call
//   - ctx context.Context
//   - message golm.UserMessage
//   - options golm.CompletionParams
//   - dest any
func (_e *MockChat_Expecter[RawRequest, RawResponse]) CompletionJSON(ctx interface{}, message interface{}, options interface{}, dest interface{}) *MockChat_CompletionJSON_Call[RawRequest, RawResponse] {
	return &MockChat_CompletionJSON_Call[RawRequest, RawResponse]{Call: _e.mock.On("CompletionJSON", ctx, message, options, dest)}
}

func (_c *MockChat_CompletionJSON_Call[RawRequest, RawResponse]) Run(run func(ctx context.Context, message golm.UserMessage, options golm.CompletionParams, dest any)) *MockChat_CompletionJSON_Call[RawRequest, RawResponse] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(golm.UserMessage), args[2].(golm.CompletionParams), args[3].(any))
	})
	return _c
}

func (_c *MockChat_CompletionJSON_Call[RawRequest, RawResponse]) Return(err error) *MockChat_CompletionJSON_Call[RawRequest, RawResponse] {
	_c.Call.Return(err)
	return _c
}

func (_c *MockChat_CompletionJSON_Call[RawRequest, RawResponse]) RunAndReturn(run func(context.Context, golm.UserMessage, golm.CompletionParams, any) error) *MockChat_CompletionJSON_Call[RawRequest, RawResponse] {
	_c.Call.Return(run)
	return _c
}

// CompletionStream provides a mock function with given fields: ctx, message, options
func (_m *MockChat[RawRequest, RawResponse]) CompletionStream(ctx context.Context, message golm.UserMessage, options golm.CompletionParams) (<-chan string, utils.StreamWaitFn) {
	ret := _m.Called(ctx, message, options)

	if len(ret) == 0 {
		panic("no return value specified for CompletionStream")
	}

	var r0 <-chan string
	var r1 utils.StreamWaitFn
	if rf, ok := ret.Get(0).(func(context.Context, golm.UserMessage, golm.CompletionParams) (<-chan string, utils.StreamWaitFn)); ok {
		return rf(ctx, message, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, golm.UserMessage, golm.CompletionParams) <-chan string); ok {
		r0 = rf(ctx, message, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, golm.UserMessage, golm.CompletionParams) utils.StreamWaitFn); ok {
		r1 = rf(ctx, message, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(utils.StreamWaitFn)
		}
	}

	return r0, r1
}

// MockChat_CompletionStream_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompletionStream'
type MockChat_CompletionStream_Call[RawRequest any, RawResponse any] struct {
	*mock.Call
}

// CompletionStream is a helper method to define mock.On call
//   - ctx context.Context
//   - message golm.UserMessage
//   - options golm.CompletionParams
func (_e *MockChat_Expecter[RawRequest, RawResponse]) CompletionStream(ctx interface{}, message interface{}, options interface{}) *MockChat_CompletionStream_Call[RawRequest, RawResponse] {
	return &MockChat_CompletionStream_Call[RawRequest, RawResponse]{Call: _e.mock.On("CompletionStream", ctx, message, options)}
}

func (_c *MockChat_CompletionStream_Call[RawRequest, RawResponse]) Run(run func(ctx context.Context, message golm.UserMessage, options golm.CompletionParams)) *MockChat_CompletionStream_Call[RawRequest, RawResponse] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(golm.UserMessage), args[2].(golm.CompletionParams))
	})
	return _c
}

func (_c *MockChat_CompletionStream_Call[RawRequest, RawResponse]) Return(response <-chan string, wait utils.StreamWaitFn) *MockChat_CompletionStream_Call[RawRequest, RawResponse] {
	_c.Call.Return(response, wait)
	return _c
}

func (_c *MockChat_CompletionStream_Call[RawRequest, RawResponse]) RunAndReturn(run func(context.Context, golm.UserMessage, golm.CompletionParams) (<-chan string, utils.StreamWaitFn)) *MockChat_CompletionStream_Call[RawRequest, RawResponse] {
	_c.Call.Return(run)
	return _c
}

// GetHistory provides a mock function with no fields
func (_m *MockChat[RawRequest, RawResponse]) GetHistory() golm.ChatHistory {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetHistory")
	}

	var r0 golm.ChatHistory
	if rf, ok := ret.Get(0).(func() golm.ChatHistory); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(golm.ChatHistory)
	}

	return r0
}

// MockChat_GetHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHistory'
type MockChat_GetHistory_Call[RawRequest any, RawResponse any] struct {
	*mock.Call
}

// GetHistory is a helper method to define mock.On call
func (_e *MockChat_Expecter[RawRequest, RawResponse]) GetHistory() *MockChat_GetHistory_Call[RawRequest, RawResponse] {
	return &MockChat_GetHistory_Call[RawRequest, RawResponse]{Call: _e.mock.On("GetHistory")}
}

func (_c *MockChat_GetHistory_Call[RawRequest, RawResponse]) Run(run func()) *MockChat_GetHistory_Call[RawRequest, RawResponse] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockChat_GetHistory_Call[RawRequest, RawResponse]) Return(_a0 golm.ChatHistory) *MockChat_GetHistory_Call[RawRequest, RawResponse] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChat_GetHistory_Call[RawRequest, RawResponse]) RunAndReturn(run func() golm.ChatHistory) *MockChat_GetHistory_Call[RawRequest, RawResponse] {
	_c.Call.Return(run)
	return _c
}

// PushHistory provides a mock function with given fields: messages
func (_m *MockChat[RawRequest, RawResponse]) PushHistory(messages ...golm.MessageWithRole) {
	_va := make([]interface{}, len(messages))
	for _i := range messages {
		_va[_i] = messages[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockChat_PushHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PushHistory'
type MockChat_PushHistory_Call[RawRequest any, RawResponse any] struct {
	*mock.Call
}

// PushHistory is a helper method to define mock.On call
//   - messages ...golm.MessageWithRole
func (_e *MockChat_Expecter[RawRequest, RawResponse]) PushHistory(messages ...interface{}) *MockChat_PushHistory_Call[RawRequest, RawResponse] {
	return &MockChat_PushHistory_Call[RawRequest, RawResponse]{Call: _e.mock.On("PushHistory",
		append([]interface{}{}, messages...)...)}
}

func (_c *MockChat_PushHistory_Call[RawRequest, RawResponse]) Run(run func(messages ...golm.MessageWithRole)) *MockChat_PushHistory_Call[RawRequest, RawResponse] {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]golm.MessageWithRole, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(golm.MessageWithRole)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockChat_PushHistory_Call[RawRequest, RawResponse]) Return() *MockChat_PushHistory_Call[RawRequest, RawResponse] {
	_c.Call.Return()
	return _c
}

func (_c *MockChat_PushHistory_Call[RawRequest, RawResponse]) RunAndReturn(run func(...golm.MessageWithRole)) *MockChat_PushHistory_Call[RawRequest, RawResponse] {
	_c.Run(run)
	return _c
}

// RawQuery provides a mock function with given fields: ctx, request
func (_m *MockChat[RawRequest, RawResponse]) RawQuery(ctx context.Context, request RawRequest) (RawResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for RawQuery")
	}

	var r0 RawResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, RawRequest) (RawResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, RawRequest) RawResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(RawResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, RawRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChat_RawQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RawQuery'
type MockChat_RawQuery_Call[RawRequest any, RawResponse any] struct {
	*mock.Call
}

// RawQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - request RawRequest
func (_e *MockChat_Expecter[RawRequest, RawResponse]) RawQuery(ctx interface{}, request interface{}) *MockChat_RawQuery_Call[RawRequest, RawResponse] {
	return &MockChat_RawQuery_Call[RawRequest, RawResponse]{Call: _e.mock.On("RawQuery", ctx, request)}
}

func (_c *MockChat_RawQuery_Call[RawRequest, RawResponse]) Run(run func(ctx context.Context, request RawRequest)) *MockChat_RawQuery_Call[RawRequest, RawResponse] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(RawRequest))
	})
	return _c
}

func (_c *MockChat_RawQuery_Call[RawRequest, RawResponse]) Return(response RawResponse, err error) *MockChat_RawQuery_Call[RawRequest, RawResponse] {
	_c.Call.Return(response, err)
	return _c
}

func (_c *MockChat_RawQuery_Call[RawRequest, RawResponse]) RunAndReturn(run func(context.Context, RawRequest) (RawResponse, error)) *MockChat_RawQuery_Call[RawRequest, RawResponse] {
	_c.Call.Return(run)
	return _c
}

// SetHistory provides a mock function with given fields: history
func (_m *MockChat[RawRequest, RawResponse]) SetHistory(history golm.ChatHistory) {
	_m.Called(history)
}

// MockChat_SetHistory_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHistory'
type MockChat_SetHistory_Call[RawRequest any, RawResponse any] struct {
	*mock.Call
}

// SetHistory is a helper method to define mock.On call
//   - history golm.ChatHistory
func (_e *MockChat_Expecter[RawRequest, RawResponse]) SetHistory(history interface{}) *MockChat_SetHistory_Call[RawRequest, RawResponse] {
	return &MockChat_SetHistory_Call[RawRequest, RawResponse]{Call: _e.mock.On("SetHistory", history)}
}

func (_c *MockChat_SetHistory_Call[RawRequest, RawResponse]) Run(run func(history golm.ChatHistory)) *MockChat_SetHistory_Call[RawRequest, RawResponse] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(golm.ChatHistory))
	})
	return _c
}

func (_c *MockChat_SetHistory_Call[RawRequest, RawResponse]) Return() *MockChat_SetHistory_Call[RawRequest, RawResponse] {
	_c.Call.Return()
	return _c
}

func (_c *MockChat_SetHistory_Call[RawRequest, RawResponse]) RunAndReturn(run func(golm.ChatHistory)) *MockChat_SetHistory_Call[RawRequest, RawResponse] {
	_c.Run(run)
	return _c
}

// SetSystem provides a mock function with given fields: system
func (_m *MockChat[RawRequest, RawResponse]) SetSystem(system *golm.SystemMessage) {
	_m.Called(system)
}

// MockChat_SetSystem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSystem'
type MockChat_SetSystem_Call[RawRequest any, RawResponse any] struct {
	*mock.Call
}

// SetSystem is a helper method to define mock.On call
//   - system *golm.SystemMessage
func (_e *MockChat_Expecter[RawRequest, RawResponse]) SetSystem(system interface{}) *MockChat_SetSystem_Call[RawRequest, RawResponse] {
	return &MockChat_SetSystem_Call[RawRequest, RawResponse]{Call: _e.mock.On("SetSystem", system)}
}

func (_c *MockChat_SetSystem_Call[RawRequest, RawResponse]) Run(run func(system *golm.SystemMessage)) *MockChat_SetSystem_Call[RawRequest, RawResponse] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*golm.SystemMessage))
	})
	return _c
}

func (_c *MockChat_SetSystem_Call[RawRequest, RawResponse]) Return() *MockChat_SetSystem_Call[RawRequest, RawResponse] {
	_c.Call.Return()
	return _c
}

func (_c *MockChat_SetSystem_Call[RawRequest, RawResponse]) RunAndReturn(run func(*golm.SystemMessage)) *MockChat_SetSystem_Call[RawRequest, RawResponse] {
	_c.Run(run)
	return _c
}

// NewMockChat creates a new instance of MockChat. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockChat[RawRequest any, RawResponse any](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockChat[RawRequest, RawResponse] {
	mock := &MockChat[RawRequest, RawResponse]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
