// Code generated by MockGen. DO NOT EDIT.
// Source: message_create.go
//
// Generated by this command:
//
//	mockgen -source=message_create.go -destination=mock/message_create.go -package=mock_stack
//

// Package mock_stack is a generated GoMock package.
package mock_stack

import (
	reflect "reflect"
	time "time"

	event "github.com/taniko/blog/internal/domain/event"
	vo "github.com/taniko/blog/internal/domain/model/stack/vo"
	gomock "go.uber.org/mock/gomock"
)

// MockCreateMessageEvent is a mock of CreateMessageEvent interface.
type MockCreateMessageEvent struct {
	ctrl     *gomock.Controller
	recorder *MockCreateMessageEventMockRecorder
}

// MockCreateMessageEventMockRecorder is the mock recorder for MockCreateMessageEvent.
type MockCreateMessageEventMockRecorder struct {
	mock *MockCreateMessageEvent
}

// NewMockCreateMessageEvent creates a new mock instance.
func NewMockCreateMessageEvent(ctrl *gomock.Controller) *MockCreateMessageEvent {
	mock := &MockCreateMessageEvent{ctrl: ctrl}
	mock.recorder = &MockCreateMessageEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateMessageEvent) EXPECT() *MockCreateMessageEventMockRecorder {
	return m.recorder
}

// ChannelID mocks base method.
func (m *MockCreateMessageEvent) ChannelID() vo.ChannelID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChannelID")
	ret0, _ := ret[0].(vo.ChannelID)
	return ret0
}

// ChannelID indicates an expected call of ChannelID.
func (mr *MockCreateMessageEventMockRecorder) ChannelID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelID", reflect.TypeOf((*MockCreateMessageEvent)(nil).ChannelID))
}

// Content mocks base method.
func (m *MockCreateMessageEvent) Content() vo.Content {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Content")
	ret0, _ := ret[0].(vo.Content)
	return ret0
}

// Content indicates an expected call of Content.
func (mr *MockCreateMessageEventMockRecorder) Content() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Content", reflect.TypeOf((*MockCreateMessageEvent)(nil).Content))
}

// CreatedAt mocks base method.
func (m *MockCreateMessageEvent) CreatedAt() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// CreatedAt indicates an expected call of CreatedAt.
func (mr *MockCreateMessageEventMockRecorder) CreatedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatedAt", reflect.TypeOf((*MockCreateMessageEvent)(nil).CreatedAt))
}

// GetEventHeader mocks base method.
func (m *MockCreateMessageEvent) GetEventHeader() event.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventHeader")
	ret0, _ := ret[0].(event.Header)
	return ret0
}

// GetEventHeader indicates an expected call of GetEventHeader.
func (mr *MockCreateMessageEventMockRecorder) GetEventHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventHeader", reflect.TypeOf((*MockCreateMessageEvent)(nil).GetEventHeader))
}

// GetEventName mocks base method.
func (m *MockCreateMessageEvent) GetEventName() event.Name {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventName")
	ret0, _ := ret[0].(event.Name)
	return ret0
}

// GetEventName indicates an expected call of GetEventName.
func (mr *MockCreateMessageEventMockRecorder) GetEventName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventName", reflect.TypeOf((*MockCreateMessageEvent)(nil).GetEventName))
}

// MessageID mocks base method.
func (m *MockCreateMessageEvent) MessageID() vo.MessageID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageID")
	ret0, _ := ret[0].(vo.MessageID)
	return ret0
}

// MessageID indicates an expected call of MessageID.
func (mr *MockCreateMessageEventMockRecorder) MessageID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageID", reflect.TypeOf((*MockCreateMessageEvent)(nil).MessageID))
}