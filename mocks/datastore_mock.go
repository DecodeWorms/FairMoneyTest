// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	model "fairmoneytest/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDataStore is a mock of DataStore interface.
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore.
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance.
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// GetTransactionByID mocks base method.
func (m *MockDataStore) GetTransactionByID(ID string) (*model.TransactionRecords, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByID", ID)
	ret0, _ := ret[0].(*model.TransactionRecords)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByID indicates an expected call of GetTransactionByID.
func (mr *MockDataStoreMockRecorder) GetTransactionByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByID", reflect.TypeOf((*MockDataStore)(nil).GetTransactionByID), ID)
}

// GetTransactionByReference mocks base method.
func (m *MockDataStore) GetTransactionByReference(ref string) (*model.TransactionRecords, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByReference", ref)
	ret0, _ := ret[0].(*model.TransactionRecords)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByReference indicates an expected call of GetTransactionByReference.
func (mr *MockDataStoreMockRecorder) GetTransactionByReference(ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByReference", reflect.TypeOf((*MockDataStore)(nil).GetTransactionByReference), ref)
}

// RecordTransaction mocks base method.
func (m *MockDataStore) RecordTransaction(data *model.TransactionRecords) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordTransaction", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordTransaction indicates an expected call of RecordTransaction.
func (mr *MockDataStoreMockRecorder) RecordTransaction(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordTransaction", reflect.TypeOf((*MockDataStore)(nil).RecordTransaction), data)
}

// UpdateAccountBalance mocks base method.
func (m *MockDataStore) UpdateAccountBalance(ID string, data *model.TransactionRecords) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccountBalance", ID, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccountBalance indicates an expected call of UpdateAccountBalance.
func (mr *MockDataStoreMockRecorder) UpdateAccountBalance(ID, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountBalance", reflect.TypeOf((*MockDataStore)(nil).UpdateAccountBalance), ID, data)
}

// MockTransaction is a mock of Transaction interface.
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction.
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance.
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// GetTransactionByID mocks base method.
func (m *MockTransaction) GetTransactionByID(ID string) (*model.TransactionRecords, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByID", ID)
	ret0, _ := ret[0].(*model.TransactionRecords)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByID indicates an expected call of GetTransactionByID.
func (mr *MockTransactionMockRecorder) GetTransactionByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByID", reflect.TypeOf((*MockTransaction)(nil).GetTransactionByID), ID)
}

// GetTransactionByReference mocks base method.
func (m *MockTransaction) GetTransactionByReference(ref string) (*model.TransactionRecords, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByReference", ref)
	ret0, _ := ret[0].(*model.TransactionRecords)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionByReference indicates an expected call of GetTransactionByReference.
func (mr *MockTransactionMockRecorder) GetTransactionByReference(ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByReference", reflect.TypeOf((*MockTransaction)(nil).GetTransactionByReference), ref)
}

// RecordTransaction mocks base method.
func (m *MockTransaction) RecordTransaction(data *model.TransactionRecords) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordTransaction", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordTransaction indicates an expected call of RecordTransaction.
func (mr *MockTransactionMockRecorder) RecordTransaction(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordTransaction", reflect.TypeOf((*MockTransaction)(nil).RecordTransaction), data)
}

// UpdateAccountBalance mocks base method.
func (m *MockTransaction) UpdateAccountBalance(ID string, data *model.TransactionRecords) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccountBalance", ID, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccountBalance indicates an expected call of UpdateAccountBalance.
func (mr *MockTransactionMockRecorder) UpdateAccountBalance(ID, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountBalance", reflect.TypeOf((*MockTransaction)(nil).UpdateAccountBalance), ID, data)
}
