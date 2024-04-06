// Package mock_password_hash is a generated GoMock package.
package mock_password_hash

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPasswordHash is a mock of PasswordHash interface.
type MockPasswordHash struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordHashMockRecorder
}

// MockPasswordHashMockRecorder is the mock recorder for MockPasswordHash.
type MockPasswordHashMockRecorder struct {
	mock *MockPasswordHash
}

// NewMockPasswordHash creates a new mock instance.
func NewMockPasswordHash(ctrl *gomock.Controller) *MockPasswordHash {
	mock := &MockPasswordHash{ctrl: ctrl}
	mock.recorder = &MockPasswordHashMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordHash) EXPECT() *MockPasswordHashMockRecorder {
	return m.recorder
}

// CompareHashAndPassword mocks base method.
func (m *MockPasswordHash) CompareHashAndPassword(hashedPassword, plainPassword string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareHashAndPassword", hashedPassword, plainPassword)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CompareHashAndPassword indicates an expected call of CompareHashAndPassword.
func (mr *MockPasswordHashMockRecorder) CompareHashAndPassword(hashedPassword, plainPassword any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareHashAndPassword", reflect.TypeOf((*MockPasswordHash)(nil).CompareHashAndPassword), hashedPassword, plainPassword)
}

// GetHash mocks base method.
func (m *MockPasswordHash) GetHash(stringToHash string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHash", stringToHash)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHash indicates an expected call of GetHash.
func (mr *MockPasswordHashMockRecorder) GetHash(stringToHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHash", reflect.TypeOf((*MockPasswordHash)(nil).GetHash), stringToHash)
}
