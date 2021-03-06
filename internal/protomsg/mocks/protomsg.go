// Code generated by MockGen. DO NOT EDIT.
// Source: protomsg.go

// Package mocks is a generated GoMock package.
package mocks

import (
	net "net"
	protomsg "power/internal/protomsg"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProtocoler is a mock of Protocoler interface.
type MockProtocoler struct {
	ctrl     *gomock.Controller
	recorder *MockProtocolerMockRecorder
}

// MockProtocolerMockRecorder is the mock recorder for MockProtocoler.
type MockProtocolerMockRecorder struct {
	mock *MockProtocoler
}

// NewMockProtocoler creates a new mock instance.
func NewMockProtocoler(ctrl *gomock.Controller) *MockProtocoler {
	mock := &MockProtocoler{ctrl: ctrl}
	mock.recorder = &MockProtocolerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProtocoler) EXPECT() *MockProtocolerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockProtocoler) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockProtocolerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockProtocoler)(nil).Close))
}

// Message mocks base method.
func (m *MockProtocoler) Message(req *protomsg.Message) (*protomsg.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message", req)
	ret0, _ := ret[0].(*protomsg.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Message indicates an expected call of Message.
func (mr *MockProtocolerMockRecorder) Message(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockProtocoler)(nil).Message), req)
}

// MessageUDP mocks base method.
func (m *MockProtocoler) MessageUDP(addr *net.UDPAddr, req *protomsg.Message) (*protomsg.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageUDP", addr, req)
	ret0, _ := ret[0].(*protomsg.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MessageUDP indicates an expected call of MessageUDP.
func (mr *MockProtocolerMockRecorder) MessageUDP(addr, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageUDP", reflect.TypeOf((*MockProtocoler)(nil).MessageUDP), addr, req)
}

// Receiver mocks base method.
func (m *MockProtocoler) Receiver() chan *protomsg.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receiver")
	ret0, _ := ret[0].(chan *protomsg.Message)
	return ret0
}

// Receiver indicates an expected call of Receiver.
func (mr *MockProtocolerMockRecorder) Receiver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receiver", reflect.TypeOf((*MockProtocoler)(nil).Receiver))
}

// Send mocks base method.
func (m *MockProtocoler) Send(addr net.Addr, req *protomsg.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", addr, req)
}

// Send indicates an expected call of Send.
func (mr *MockProtocolerMockRecorder) Send(addr, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockProtocoler)(nil).Send), addr, req)
}
