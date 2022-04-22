// Code generated by MockGen. DO NOT EDIT.
// Source: jwk/registry.go

// Package jwk_test is a generated GoMock package.
package jwk_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	herodot "github.com/ory/herodot"
	config "github.com/ory/hydra/driver/config"
	jwk "github.com/ory/hydra/jwk"
	logrusx "github.com/ory/x/logrusx"
)

// MockInternalRegistry is a mock of InternalRegistry interface.
type MockInternalRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockInternalRegistryMockRecorder
}

// MockInternalRegistryMockRecorder is the mock recorder for MockInternalRegistry.
type MockInternalRegistryMockRecorder struct {
	mock *MockInternalRegistry
}

// NewMockInternalRegistry creates a new mock instance.
func NewMockInternalRegistry(ctrl *gomock.Controller) *MockInternalRegistry {
	mock := &MockInternalRegistry{ctrl: ctrl}
	mock.recorder = &MockInternalRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInternalRegistry) EXPECT() *MockInternalRegistryMockRecorder {
	return m.recorder
}

// AuditLogger mocks base method.
func (m *MockInternalRegistry) AuditLogger() *logrusx.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuditLogger")
	ret0, _ := ret[0].(*logrusx.Logger)
	return ret0
}

// AuditLogger indicates an expected call of AuditLogger.
func (mr *MockInternalRegistryMockRecorder) AuditLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuditLogger", reflect.TypeOf((*MockInternalRegistry)(nil).AuditLogger))
}

// Config mocks base method.
func (m *MockInternalRegistry) Config(ctx context.Context) *config.Provider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config", ctx)
	ret0, _ := ret[0].(*config.Provider)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockInternalRegistryMockRecorder) Config(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockInternalRegistry)(nil).Config), ctx)
}

// KeyCipher mocks base method.
func (m *MockInternalRegistry) KeyCipher() *jwk.AEAD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyCipher")
	ret0, _ := ret[0].(*jwk.AEAD)
	return ret0
}

// KeyCipher indicates an expected call of KeyCipher.
func (mr *MockInternalRegistryMockRecorder) KeyCipher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyCipher", reflect.TypeOf((*MockInternalRegistry)(nil).KeyCipher))
}

// KeyGenerators mocks base method.
func (m *MockInternalRegistry) KeyGenerators() map[string]jwk.KeyGenerator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyGenerators")
	ret0, _ := ret[0].(map[string]jwk.KeyGenerator)
	return ret0
}

// KeyGenerators indicates an expected call of KeyGenerators.
func (mr *MockInternalRegistryMockRecorder) KeyGenerators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyGenerators", reflect.TypeOf((*MockInternalRegistry)(nil).KeyGenerators))
}

// KeyManager mocks base method.
func (m *MockInternalRegistry) KeyManager() jwk.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyManager")
	ret0, _ := ret[0].(jwk.Manager)
	return ret0
}

// KeyManager indicates an expected call of KeyManager.
func (mr *MockInternalRegistryMockRecorder) KeyManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyManager", reflect.TypeOf((*MockInternalRegistry)(nil).KeyManager))
}

// Logger mocks base method.
func (m *MockInternalRegistry) Logger() *logrusx.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(*logrusx.Logger)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockInternalRegistryMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockInternalRegistry)(nil).Logger))
}

// SoftwareKeyManager mocks base method.
func (m *MockInternalRegistry) SoftwareKeyManager() jwk.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SoftwareKeyManager")
	ret0, _ := ret[0].(jwk.Manager)
	return ret0
}

// SoftwareKeyManager indicates an expected call of SoftwareKeyManager.
func (mr *MockInternalRegistryMockRecorder) SoftwareKeyManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftwareKeyManager", reflect.TypeOf((*MockInternalRegistry)(nil).SoftwareKeyManager))
}

// Writer mocks base method.
func (m *MockInternalRegistry) Writer() herodot.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Writer")
	ret0, _ := ret[0].(herodot.Writer)
	return ret0
}

// Writer indicates an expected call of Writer.
func (mr *MockInternalRegistryMockRecorder) Writer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Writer", reflect.TypeOf((*MockInternalRegistry)(nil).Writer))
}

// MockRegistry is a mock of Registry interface.
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry.
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance.
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockRegistry) Config(ctx context.Context) *config.Provider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config", ctx)
	ret0, _ := ret[0].(*config.Provider)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockRegistryMockRecorder) Config(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockRegistry)(nil).Config), ctx)
}

// KeyCipher mocks base method.
func (m *MockRegistry) KeyCipher() *jwk.AEAD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyCipher")
	ret0, _ := ret[0].(*jwk.AEAD)
	return ret0
}

// KeyCipher indicates an expected call of KeyCipher.
func (mr *MockRegistryMockRecorder) KeyCipher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyCipher", reflect.TypeOf((*MockRegistry)(nil).KeyCipher))
}

// KeyGenerators mocks base method.
func (m *MockRegistry) KeyGenerators() map[string]jwk.KeyGenerator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyGenerators")
	ret0, _ := ret[0].(map[string]jwk.KeyGenerator)
	return ret0
}

// KeyGenerators indicates an expected call of KeyGenerators.
func (mr *MockRegistryMockRecorder) KeyGenerators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyGenerators", reflect.TypeOf((*MockRegistry)(nil).KeyGenerators))
}

// KeyManager mocks base method.
func (m *MockRegistry) KeyManager() jwk.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyManager")
	ret0, _ := ret[0].(jwk.Manager)
	return ret0
}

// KeyManager indicates an expected call of KeyManager.
func (mr *MockRegistryMockRecorder) KeyManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyManager", reflect.TypeOf((*MockRegistry)(nil).KeyManager))
}

// SoftwareKeyManager mocks base method.
func (m *MockRegistry) SoftwareKeyManager() jwk.Manager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SoftwareKeyManager")
	ret0, _ := ret[0].(jwk.Manager)
	return ret0
}

// SoftwareKeyManager indicates an expected call of SoftwareKeyManager.
func (mr *MockRegistryMockRecorder) SoftwareKeyManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftwareKeyManager", reflect.TypeOf((*MockRegistry)(nil).SoftwareKeyManager))
}
