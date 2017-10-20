// Code generated by MockGen. DO NOT EDIT.
// Source: geth/common/geth/node_constructor.go

// Package geth is a generated GoMock package.
package geth

import (
	gomock "github.com/golang/mock/gomock"
	params "github.com/status-im/status-go/geth/params"
	reflect "reflect"
)

// MockNodeConstructor is a mock of NodeConstructor interface
type MockNodeConstructor struct {
	ctrl     *gomock.Controller
	recorder *MockNodeConstructorMockRecorder
}

// MockNodeConstructorMockRecorder is the mock recorder for MockNodeConstructor
type MockNodeConstructorMockRecorder struct {
	mock *MockNodeConstructor
}

// NewMockNodeConstructor creates a new mock instance
func NewMockNodeConstructor(ctrl *gomock.Controller) *MockNodeConstructor {
	mock := &MockNodeConstructor{ctrl: ctrl}
	mock.recorder = &MockNodeConstructorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeConstructor) EXPECT() *MockNodeConstructorMockRecorder {
	return m.recorder
}

// Make mocks base method
func (m *MockNodeConstructor) Make() (Node, error) {
	ret := m.ctrl.Call(m, "Make")
	ret0, _ := ret[0].(Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Make indicates an expected call of Make
func (mr *MockNodeConstructorMockRecorder) Make() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Make", reflect.TypeOf((*MockNodeConstructor)(nil).Make))
}

// Config mocks base method
func (m *MockNodeConstructor) Config() *params.NodeConfig {
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*params.NodeConfig)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockNodeConstructorMockRecorder) Config() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockNodeConstructor)(nil).Config))
}

// SetConfig mocks base method
func (m *MockNodeConstructor) SetConfig(config *params.NodeConfig) {
	m.ctrl.Call(m, "SetConfig", config)
}

// SetConfig indicates an expected call of SetConfig
func (mr *MockNodeConstructorMockRecorder) SetConfig(config interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockNodeConstructor)(nil).SetConfig), config)
}