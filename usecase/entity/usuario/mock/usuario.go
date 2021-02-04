// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/entity/usuario/interface.go

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
func (m *MockRead) Get(email string) (*entity.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", email)
	ret0, _ := ret[0].(*entity.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockReadMockRecorder) Get(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRead)(nil).Get), email)
}

// Search mocks base method
func (m *MockRead) Search(query string) ([]*entity.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*entity.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockReadMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRead)(nil).Search), query)
}

// List mocks base method
func (m *MockRead) List() ([]*entity.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Usuario)
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
func (m *MockWrite) Create(e *entity.Usuario) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockWriteMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWrite)(nil).Create), e)
}

// Update mocks base method
func (m *MockWrite) Update(e *entity.Usuario) error {
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
func (m *MockWrite) Delete(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockWriteMockRecorder) Delete(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWrite)(nil).Delete), email)
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
func (m *MockRepository) Get(email string) (*entity.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", email)
	ret0, _ := ret[0].(*entity.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), email)
}

// Search mocks base method
func (m *MockRepository) Search(query string) ([]*entity.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*entity.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockRepositoryMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), query)
}

// List mocks base method
func (m *MockRepository) List() ([]*entity.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List))
}

// Create mocks base method
func (m *MockRepository) Create(e *entity.Usuario) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRepositoryMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), e)
}

// Update mocks base method
func (m *MockRepository) Update(e *entity.Usuario) error {
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
func (m *MockRepository) Delete(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), email)
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

// GetUsuario mocks base method
func (m *MockUseCase) GetUsuario(email string) (*entity.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsuario", email)
	ret0, _ := ret[0].(*entity.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsuario indicates an expected call of GetUsuario
func (mr *MockUseCaseMockRecorder) GetUsuario(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsuario", reflect.TypeOf((*MockUseCase)(nil).GetUsuario), email)
}

// SearchUsuarios mocks base method
func (m *MockUseCase) SearchUsuarios(query string) ([]*entity.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUsuarios", query)
	ret0, _ := ret[0].([]*entity.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUsuarios indicates an expected call of SearchUsuarios
func (mr *MockUseCaseMockRecorder) SearchUsuarios(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUsuarios", reflect.TypeOf((*MockUseCase)(nil).SearchUsuarios), query)
}

// ListUsuarios mocks base method
func (m *MockUseCase) ListUsuarios() ([]*entity.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsuarios")
	ret0, _ := ret[0].([]*entity.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsuarios indicates an expected call of ListUsuarios
func (mr *MockUseCaseMockRecorder) ListUsuarios() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsuarios", reflect.TypeOf((*MockUseCase)(nil).ListUsuarios))
}

// CreateUsuario mocks base method
func (m *MockUseCase) CreateUsuario(email, password, birthday string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUsuario", email, password, birthday)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUsuario indicates an expected call of CreateUsuario
func (mr *MockUseCaseMockRecorder) CreateUsuario(email, password, birthday interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUsuario", reflect.TypeOf((*MockUseCase)(nil).CreateUsuario), email, password, birthday)
}

// UpdateUsuario mocks base method
func (m *MockUseCase) UpdateUsuario(e *entity.Usuario) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUsuario", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUsuario indicates an expected call of UpdateUsuario
func (mr *MockUseCaseMockRecorder) UpdateUsuario(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUsuario", reflect.TypeOf((*MockUseCase)(nil).UpdateUsuario), e)
}

// DeleteUsuario mocks base method
func (m *MockUseCase) DeleteUsuario(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsuario", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUsuario indicates an expected call of DeleteUsuario
func (mr *MockUseCaseMockRecorder) DeleteUsuario(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsuario", reflect.TypeOf((*MockUseCase)(nil).DeleteUsuario), email)
}