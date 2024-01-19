// Code generated by MockGen. DO NOT EDIT.
// Source: repository/employee.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/najibjodiansyah/mekari-employee/model/domain"
)

// MockEmployeeRepository is a mock of EmployeeRepository interface.
type MockEmployeeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeRepositoryMockRecorder
}

// MockEmployeeRepositoryMockRecorder is the mock recorder for MockEmployeeRepository.
type MockEmployeeRepositoryMockRecorder struct {
	mock *MockEmployeeRepository
}

// NewMockEmployeeRepository creates a new mock instance.
func NewMockEmployeeRepository(ctrl *gomock.Controller) *MockEmployeeRepository {
	mock := &MockEmployeeRepository{ctrl: ctrl}
	mock.recorder = &MockEmployeeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeRepository) EXPECT() *MockEmployeeRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockEmployeeRepository) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEmployeeRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEmployeeRepository)(nil).Delete), ctx, id)
}

// Insert mocks base method.
func (m *MockEmployeeRepository) Insert(ctx context.Context, employee *domain.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, employee)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockEmployeeRepositoryMockRecorder) Insert(ctx, employee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockEmployeeRepository)(nil).Insert), ctx, employee)
}

// SelectAll mocks base method.
func (m *MockEmployeeRepository) SelectAll(ctx context.Context) ([]*domain.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAll", ctx)
	ret0, _ := ret[0].([]*domain.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAll indicates an expected call of SelectAll.
func (mr *MockEmployeeRepositoryMockRecorder) SelectAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAll", reflect.TypeOf((*MockEmployeeRepository)(nil).SelectAll), ctx)
}

// SelectById mocks base method.
func (m *MockEmployeeRepository) SelectById(ctx context.Context, id int) (*domain.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectById", ctx, id)
	ret0, _ := ret[0].(*domain.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectById indicates an expected call of SelectById.
func (mr *MockEmployeeRepositoryMockRecorder) SelectById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectById", reflect.TypeOf((*MockEmployeeRepository)(nil).SelectById), ctx, id)
}

// Update mocks base method.
func (m *MockEmployeeRepository) Update(ctx context.Context, employee *domain.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, employee)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockEmployeeRepositoryMockRecorder) Update(ctx, employee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEmployeeRepository)(nil).Update), ctx, employee)
}