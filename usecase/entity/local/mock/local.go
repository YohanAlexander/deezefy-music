// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/entity/local/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/yohanalexander/deezefy-music/entity"
	reflect "reflect"
)

// MockRead is a mock of Read interface
type MockRead struct {
	ctrl     *gomock.Controller
	recorder *MockReadMockRecorder
}

// MockReadMockRecorder is the mock recorder for MockRead
type MockReadMockRecorder struct {
	mock *MockRead
}

// NewMockRead creates a new mock instance
func NewMockRead(ctrl *gomock.Controller) *MockRead {
	mock := &MockRead{ctrl: ctrl}
	mock.recorder = &MockReadMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRead) EXPECT() *MockReadMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockRead) Get(id int) (*entity.Local, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Local)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockReadMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRead)(nil).Get), id)
}

// Search mocks base method
func (m *MockRead) Search(query string) ([]*entity.Local, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*entity.Local)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockReadMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRead)(nil).Search), query)
}

// List mocks base method
func (m *MockRead) List() ([]*entity.Local, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Local)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockReadMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRead)(nil).List))
}

// MockWrite is a mock of Write interface
type MockWrite struct {
	ctrl     *gomock.Controller
	recorder *MockWriteMockRecorder
}

// MockWriteMockRecorder is the mock recorder for MockWrite
type MockWriteMockRecorder struct {
	mock *MockWrite
}

// NewMockWrite creates a new mock instance
func NewMockWrite(ctrl *gomock.Controller) *MockWrite {
	mock := &MockWrite{ctrl: ctrl}
	mock.recorder = &MockWriteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWrite) EXPECT() *MockWriteMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockWrite) Create(e *entity.Local) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockWriteMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWrite)(nil).Create), e)
}

// Update mocks base method
func (m *MockWrite) Update(e *entity.Local) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockWriteMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWrite)(nil).Update), e)
}

// Delete mocks base method
func (m *MockWrite) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockWriteMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWrite)(nil).Delete), id)
}

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockRepository) Get(id int) (*entity.Local, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Local)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), id)
}

// Search mocks base method
func (m *MockRepository) Search(query string) ([]*entity.Local, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*entity.Local)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockRepositoryMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), query)
}

// List mocks base method
func (m *MockRepository) List() ([]*entity.Local, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Local)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List))
}

// Create mocks base method
func (m *MockRepository) Create(e *entity.Local) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRepositoryMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), e)
}

// Update mocks base method
func (m *MockRepository) Update(e *entity.Local) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRepositoryMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), e)
}

// Delete mocks base method
func (m *MockRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), id)
}

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// GetLocal mocks base method
func (m *MockUseCase) GetLocal(id int) (*entity.Local, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocal", id)
	ret0, _ := ret[0].(*entity.Local)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocal indicates an expected call of GetLocal
func (mr *MockUseCaseMockRecorder) GetLocal(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocal", reflect.TypeOf((*MockUseCase)(nil).GetLocal), id)
}

// SearchLocals mocks base method
func (m *MockUseCase) SearchLocals(query string) ([]*entity.Local, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchLocals", query)
	ret0, _ := ret[0].([]*entity.Local)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchLocals indicates an expected call of SearchLocals
func (mr *MockUseCaseMockRecorder) SearchLocals(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchLocals", reflect.TypeOf((*MockUseCase)(nil).SearchLocals), query)
}

// ListLocals mocks base method
func (m *MockUseCase) ListLocals() ([]*entity.Local, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLocals")
	ret0, _ := ret[0].([]*entity.Local)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLocals indicates an expected call of ListLocals
func (mr *MockUseCaseMockRecorder) ListLocals() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLocals", reflect.TypeOf((*MockUseCase)(nil).ListLocals))
}

// CreateLocal mocks base method
func (m *MockUseCase) CreateLocal(cidade, pais string, id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLocal", cidade, pais, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLocal indicates an expected call of CreateLocal
func (mr *MockUseCaseMockRecorder) CreateLocal(cidade, pais, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLocal", reflect.TypeOf((*MockUseCase)(nil).CreateLocal), cidade, pais, id)
}

// UpdateLocal mocks base method
func (m *MockUseCase) UpdateLocal(e *entity.Local) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLocal", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLocal indicates an expected call of UpdateLocal
func (mr *MockUseCaseMockRecorder) UpdateLocal(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLocal", reflect.TypeOf((*MockUseCase)(nil).UpdateLocal), e)
}

// DeleteLocal mocks base method
func (m *MockUseCase) DeleteLocal(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLocal", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLocal indicates an expected call of DeleteLocal
func (mr *MockUseCaseMockRecorder) DeleteLocal(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLocal", reflect.TypeOf((*MockUseCase)(nil).DeleteLocal), id)
}