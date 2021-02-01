// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/usuario/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	usuario "github.com/yohanalexander/deezefy-music/entity/usuario"
	reflect "reflect"
)

// MockUsuario is a mock of Usuario interface
type MockUsuario struct {
	ctrl     *gomock.Controller
	recorder *MockUsuarioMockRecorder
}

// MockUsuarioMockRecorder is the mock recorder for MockUsuario
type MockUsuarioMockRecorder struct {
	mock *MockUsuario
}

// NewMockUsuario creates a new mock instance
func NewMockUsuario(ctrl *gomock.Controller) *MockUsuario {
	mock := &MockUsuario{ctrl: ctrl}
	mock.recorder = &MockUsuarioMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsuario) EXPECT() *MockUsuarioMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockUsuario) Get(email string) (*usuario.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", email)
	ret0, _ := ret[0].(*usuario.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUsuarioMockRecorder) Get(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsuario)(nil).Get), email)
}

// Search mocks base method
func (m *MockUsuario) Search(query string) ([]*usuario.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*usuario.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockUsuarioMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockUsuario)(nil).Search), query)
}

// List mocks base method
func (m *MockUsuario) List() ([]*usuario.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*usuario.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockUsuarioMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUsuario)(nil).List))
}

// Create mocks base method
func (m *MockUsuario) Create(e *usuario.Usuario) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUsuarioMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsuario)(nil).Create), e)
}

// Update mocks base method
func (m *MockUsuario) Update(e *usuario.Usuario) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockUsuarioMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUsuario)(nil).Update), e)
}

// Delete mocks base method
func (m *MockUsuario) Delete(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockUsuarioMockRecorder) Delete(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUsuario)(nil).Delete), email)
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
func (m *MockRepository) Get(email string) (*usuario.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", email)
	ret0, _ := ret[0].(*usuario.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), email)
}

// Search mocks base method
func (m *MockRepository) Search(query string) ([]*usuario.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*usuario.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockRepositoryMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), query)
}

// List mocks base method
func (m *MockRepository) List() ([]*usuario.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*usuario.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List))
}

// Create mocks base method
func (m *MockRepository) Create(e *usuario.Usuario) (string, error) {
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
func (m *MockRepository) Update(e *usuario.Usuario) error {
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
func (m *MockUseCase) GetUsuario(email string) (*usuario.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsuario", email)
	ret0, _ := ret[0].(*usuario.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsuario indicates an expected call of GetUsuario
func (mr *MockUseCaseMockRecorder) GetUsuario(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsuario", reflect.TypeOf((*MockUseCase)(nil).GetUsuario), email)
}

// SearchUsuarios mocks base method
func (m *MockUseCase) SearchUsuarios(query string) ([]*usuario.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUsuarios", query)
	ret0, _ := ret[0].([]*usuario.Usuario)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUsuarios indicates an expected call of SearchUsuarios
func (mr *MockUseCaseMockRecorder) SearchUsuarios(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUsuarios", reflect.TypeOf((*MockUseCase)(nil).SearchUsuarios), query)
}

// ListUsuarios mocks base method
func (m *MockUseCase) ListUsuarios() ([]*usuario.Usuario, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsuarios")
	ret0, _ := ret[0].([]*usuario.Usuario)
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
func (m *MockUseCase) UpdateUsuario(e *usuario.Usuario) error {
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