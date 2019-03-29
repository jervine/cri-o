// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubernetes-sigs/cri-o/lib/sandbox (interfaces: NetNsIface)

// Package sandboxmock is a generated GoMock package.
package sandboxmock

import (
	gomock "github.com/golang/mock/gomock"
	sandbox "github.com/kubernetes-sigs/cri-o/lib/sandbox"
	reflect "reflect"
)

// MockNetNsIface is a mock of NetNsIface interface
type MockNetNsIface struct {
	ctrl     *gomock.Controller
	recorder *MockNetNsIfaceMockRecorder
}

// MockNetNsIfaceMockRecorder is the mock recorder for MockNetNsIface
type MockNetNsIfaceMockRecorder struct {
	mock *MockNetNsIface
}

// NewMockNetNsIface creates a new mock instance
func NewMockNetNsIface(ctrl *gomock.Controller) *MockNetNsIface {
	mock := &MockNetNsIface{ctrl: ctrl}
	mock.recorder = &MockNetNsIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetNsIface) EXPECT() *MockNetNsIfaceMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockNetNsIface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockNetNsIfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockNetNsIface)(nil).Close))
}

// Get mocks base method
func (m *MockNetNsIface) Get() *sandbox.NetNs {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*sandbox.NetNs)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockNetNsIfaceMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNetNsIface)(nil).Get))
}

// Initialize mocks base method
func (m *MockNetNsIface) Initialize() (sandbox.NetNsIface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize")
	ret0, _ := ret[0].(sandbox.NetNsIface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Initialize indicates an expected call of Initialize
func (mr *MockNetNsIfaceMockRecorder) Initialize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockNetNsIface)(nil).Initialize))
}

// Initialized mocks base method
func (m *MockNetNsIface) Initialized() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialized")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Initialized indicates an expected call of Initialized
func (mr *MockNetNsIfaceMockRecorder) Initialized() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialized", reflect.TypeOf((*MockNetNsIface)(nil).Initialized))
}

// Remove mocks base method
func (m *MockNetNsIface) Remove() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove")
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockNetNsIfaceMockRecorder) Remove() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockNetNsIface)(nil).Remove))
}

// SymlinkCreate mocks base method
func (m *MockNetNsIface) SymlinkCreate(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SymlinkCreate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SymlinkCreate indicates an expected call of SymlinkCreate
func (mr *MockNetNsIfaceMockRecorder) SymlinkCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SymlinkCreate", reflect.TypeOf((*MockNetNsIface)(nil).SymlinkCreate), arg0)
}