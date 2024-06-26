// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/demo-talent/services (interfaces: ExpenseService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entities "github.com/demo-talent/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockExpenseService is a mock of ExpenseService interface.
type MockExpenseService struct {
	ctrl     *gomock.Controller
	recorder *MockExpenseServiceMockRecorder
}

// MockExpenseServiceMockRecorder is the mock recorder for MockExpenseService.
type MockExpenseServiceMockRecorder struct {
	mock *MockExpenseService
}

// NewMockExpenseService creates a new mock instance.
func NewMockExpenseService(ctrl *gomock.Controller) *MockExpenseService {
	mock := &MockExpenseService{ctrl: ctrl}
	mock.recorder = &MockExpenseServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExpenseService) EXPECT() *MockExpenseServiceMockRecorder {
	return m.recorder
}

// CreateExpense mocks base method.
func (m *MockExpenseService) CreateExpense(arg0 context.Context, arg1 *entities.Expense) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExpense", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExpense indicates an expected call of CreateExpense.
func (mr *MockExpenseServiceMockRecorder) CreateExpense(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExpense", reflect.TypeOf((*MockExpenseService)(nil).CreateExpense), arg0, arg1)
}

// DeleteExpense mocks base method.
func (m *MockExpenseService) DeleteExpense(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpense", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExpense indicates an expected call of DeleteExpense.
func (mr *MockExpenseServiceMockRecorder) DeleteExpense(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpense", reflect.TypeOf((*MockExpenseService)(nil).DeleteExpense), arg0, arg1)
}

// GetExpenseByID mocks base method.
func (m *MockExpenseService) GetExpenseByID(arg0 context.Context, arg1 string) (*entities.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpenseByID", arg0, arg1)
	ret0, _ := ret[0].(*entities.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpenseByID indicates an expected call of GetExpenseByID.
func (mr *MockExpenseServiceMockRecorder) GetExpenseByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpenseByID", reflect.TypeOf((*MockExpenseService)(nil).GetExpenseByID), arg0, arg1)
}

// UpdateExpense mocks base method.
func (m *MockExpenseService) UpdateExpense(arg0 context.Context, arg1 *entities.Expense) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExpense", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExpense indicates an expected call of UpdateExpense.
func (mr *MockExpenseServiceMockRecorder) UpdateExpense(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExpense", reflect.TypeOf((*MockExpenseService)(nil).UpdateExpense), arg0, arg1)
}
