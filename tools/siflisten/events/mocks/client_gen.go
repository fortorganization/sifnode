// Code generated by MockGen. DO NOT EDIT.
// Source: ../api.go

// Package eventmocks is a generated GoMock package.
package eventmocks

import (
	context "context"
	reflect "reflect"

	events "github.com/Sifchain/sifnode/tools/siflisten/events"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ConsumeForever mocks base method.
func (m *MockClient) ConsumeForever(ctx context.Context, cursorName string, f events.Consumer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ConsumeForever", ctx, cursorName, f)
}

// ConsumeForever indicates an expected call of ConsumeForever.
func (mr *MockClientMockRecorder) ConsumeForever(ctx, cursorName, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeForever", reflect.TypeOf((*MockClient)(nil).ConsumeForever), ctx, cursorName, f)
}

// CreateEvent mocks base method.
func (m *MockClient) CreateEvent(ctx context.Context, ev *events.Event) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", ctx, ev)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockClientMockRecorder) CreateEvent(ctx, ev interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockClient)(nil).CreateEvent), ctx, ev)
}

// GetCursor mocks base method.
func (m *MockClient) GetCursor(ctx context.Context, name string) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCursor", ctx, name)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetCursor indicates an expected call of GetCursor.
func (mr *MockClientMockRecorder) GetCursor(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCursor", reflect.TypeOf((*MockClient)(nil).GetCursor), ctx, name)
}

// GetEvent mocks base method.
func (m *MockClient) GetEvent(ctx context.Context, cursorPosition int64) (*events.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvent", ctx, cursorPosition)
	ret0, _ := ret[0].(*events.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvent indicates an expected call of GetEvent.
func (mr *MockClientMockRecorder) GetEvent(ctx, cursorPosition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockClient)(nil).GetEvent), ctx, cursorPosition)
}

// GetEvents mocks base method.
func (m *MockClient) GetEvents(ctx context.Context, cursorPosition int64) ([]*events.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvents", ctx, cursorPosition)
	ret0, _ := ret[0].([]*events.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvents indicates an expected call of GetEvents.
func (mr *MockClientMockRecorder) GetEvents(ctx, cursorPosition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvents", reflect.TypeOf((*MockClient)(nil).GetEvents), ctx, cursorPosition)
}

// SetCursor mocks base method.
func (m *MockClient) SetCursor(ctx context.Context, name string, position int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCursor", ctx, name, position)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCursor indicates an expected call of SetCursor.
func (mr *MockClientMockRecorder) SetCursor(ctx, name, position interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCursor", reflect.TypeOf((*MockClient)(nil).SetCursor), ctx, name, position)
}
