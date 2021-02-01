// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/generosfavoritos/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	generosfavoritos "github.com/yohanalexander/deezefy-music/entity/generosfavoritos"
	reflect "reflect"
)

// MockGenerosFavoritos is a mock of GenerosFavoritos interface
type MockGenerosFavoritos struct {
	ctrl     *gomock.Controller
	recorder *MockGenerosFavoritosMockRecorder
}

// MockGenerosFavoritosMockRecorder is the mock recorder for MockGenerosFavoritos
type MockGenerosFavoritosMockRecorder struct {
	mock *MockGenerosFavoritos
}

// NewMockGenerosFavoritos creates a new mock instance
func NewMockGenerosFavoritos(ctrl *gomock.Controller) *MockGenerosFavoritos {
	mock := &MockGenerosFavoritos{ctrl: ctrl}
	mock.recorder = &MockGenerosFavoritosMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGenerosFavoritos) EXPECT() *MockGenerosFavoritosMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockGenerosFavoritos) Get(perfil int, genero, ouvinte string) (*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", perfil, genero, ouvinte)
	ret0, _ := ret[0].(*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockGenerosFavoritosMockRecorder) Get(perfil, genero, ouvinte interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGenerosFavoritos)(nil).Get), perfil, genero, ouvinte)
}

// GetByPerfil mocks base method
func (m *MockGenerosFavoritos) GetByPerfil(perfil int) (*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPerfil", perfil)
	ret0, _ := ret[0].(*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPerfil indicates an expected call of GetByPerfil
func (mr *MockGenerosFavoritosMockRecorder) GetByPerfil(perfil interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPerfil", reflect.TypeOf((*MockGenerosFavoritos)(nil).GetByPerfil), perfil)
}

// GetByGenero mocks base method
func (m *MockGenerosFavoritos) GetByGenero(genero string) (*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByGenero", genero)
	ret0, _ := ret[0].(*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByGenero indicates an expected call of GetByGenero
func (mr *MockGenerosFavoritosMockRecorder) GetByGenero(genero interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByGenero", reflect.TypeOf((*MockGenerosFavoritos)(nil).GetByGenero), genero)
}

// Search mocks base method
func (m *MockGenerosFavoritos) Search(query string) ([]*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockGenerosFavoritosMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockGenerosFavoritos)(nil).Search), query)
}

// List mocks base method
func (m *MockGenerosFavoritos) List() ([]*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockGenerosFavoritosMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockGenerosFavoritos)(nil).List))
}

// Create mocks base method
func (m *MockGenerosFavoritos) Create(e *generosfavoritos.GenerosFavoritos) (int, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Create indicates an expected call of Create
func (mr *MockGenerosFavoritosMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGenerosFavoritos)(nil).Create), e)
}

// Update mocks base method
func (m *MockGenerosFavoritos) Update(e *generosfavoritos.GenerosFavoritos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockGenerosFavoritosMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGenerosFavoritos)(nil).Update), e)
}

// Delete mocks base method
func (m *MockGenerosFavoritos) Delete(perfil int, genero, ouvinte string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", perfil, genero, ouvinte)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockGenerosFavoritosMockRecorder) Delete(perfil, genero, ouvinte interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGenerosFavoritos)(nil).Delete), perfil, genero, ouvinte)
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
func (m *MockRepository) Get(perfil int, genero, ouvinte string) (*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", perfil, genero, ouvinte)
	ret0, _ := ret[0].(*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(perfil, genero, ouvinte interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), perfil, genero, ouvinte)
}

// GetByPerfil mocks base method
func (m *MockRepository) GetByPerfil(perfil int) (*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPerfil", perfil)
	ret0, _ := ret[0].(*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPerfil indicates an expected call of GetByPerfil
func (mr *MockRepositoryMockRecorder) GetByPerfil(perfil interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPerfil", reflect.TypeOf((*MockRepository)(nil).GetByPerfil), perfil)
}

// GetByGenero mocks base method
func (m *MockRepository) GetByGenero(genero string) (*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByGenero", genero)
	ret0, _ := ret[0].(*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByGenero indicates an expected call of GetByGenero
func (mr *MockRepositoryMockRecorder) GetByGenero(genero interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByGenero", reflect.TypeOf((*MockRepository)(nil).GetByGenero), genero)
}

// Search mocks base method
func (m *MockRepository) Search(query string) ([]*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockRepositoryMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), query)
}

// List mocks base method
func (m *MockRepository) List() ([]*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List))
}

// Create mocks base method
func (m *MockRepository) Create(e *generosfavoritos.GenerosFavoritos) (int, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Create indicates an expected call of Create
func (mr *MockRepositoryMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), e)
}

// Update mocks base method
func (m *MockRepository) Update(e *generosfavoritos.GenerosFavoritos) error {
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
func (m *MockRepository) Delete(perfil int, genero, ouvinte string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", perfil, genero, ouvinte)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(perfil, genero, ouvinte interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), perfil, genero, ouvinte)
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

// GetGenerosFavoritos mocks base method
func (m *MockUseCase) GetGenerosFavoritos(perfil int, genero, ouvinte string) (*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenerosFavoritos", perfil, genero, ouvinte)
	ret0, _ := ret[0].(*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenerosFavoritos indicates an expected call of GetGenerosFavoritos
func (mr *MockUseCaseMockRecorder) GetGenerosFavoritos(perfil, genero, ouvinte interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenerosFavoritos", reflect.TypeOf((*MockUseCase)(nil).GetGenerosFavoritos), perfil, genero, ouvinte)
}

// GetGenerosFavoritosByPerfil mocks base method
func (m *MockUseCase) GetGenerosFavoritosByPerfil(perfil int) (*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenerosFavoritosByPerfil", perfil)
	ret0, _ := ret[0].(*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenerosFavoritosByPerfil indicates an expected call of GetGenerosFavoritosByPerfil
func (mr *MockUseCaseMockRecorder) GetGenerosFavoritosByPerfil(perfil interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenerosFavoritosByPerfil", reflect.TypeOf((*MockUseCase)(nil).GetGenerosFavoritosByPerfil), perfil)
}

// GetGenerosFavoritosByGenero mocks base method
func (m *MockUseCase) GetGenerosFavoritosByGenero(genero string) (*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenerosFavoritosByGenero", genero)
	ret0, _ := ret[0].(*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenerosFavoritosByGenero indicates an expected call of GetGenerosFavoritosByGenero
func (mr *MockUseCaseMockRecorder) GetGenerosFavoritosByGenero(genero interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenerosFavoritosByGenero", reflect.TypeOf((*MockUseCase)(nil).GetGenerosFavoritosByGenero), genero)
}

// SearchGenerosFavoritoss mocks base method
func (m *MockUseCase) SearchGenerosFavoritoss(query string) ([]*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchGenerosFavoritoss", query)
	ret0, _ := ret[0].([]*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchGenerosFavoritoss indicates an expected call of SearchGenerosFavoritoss
func (mr *MockUseCaseMockRecorder) SearchGenerosFavoritoss(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchGenerosFavoritoss", reflect.TypeOf((*MockUseCase)(nil).SearchGenerosFavoritoss), query)
}

// ListGenerosFavoritoss mocks base method
func (m *MockUseCase) ListGenerosFavoritoss() ([]*generosfavoritos.GenerosFavoritos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGenerosFavoritoss")
	ret0, _ := ret[0].([]*generosfavoritos.GenerosFavoritos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGenerosFavoritoss indicates an expected call of ListGenerosFavoritoss
func (mr *MockUseCaseMockRecorder) ListGenerosFavoritoss() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGenerosFavoritoss", reflect.TypeOf((*MockUseCase)(nil).ListGenerosFavoritoss))
}

// CreateGenerosFavoritos mocks base method
func (m *MockUseCase) CreateGenerosFavoritos(perfil int, genero, ouvinte string) (int, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGenerosFavoritos", perfil, genero, ouvinte)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// CreateGenerosFavoritos indicates an expected call of CreateGenerosFavoritos
func (mr *MockUseCaseMockRecorder) CreateGenerosFavoritos(perfil, genero, ouvinte interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGenerosFavoritos", reflect.TypeOf((*MockUseCase)(nil).CreateGenerosFavoritos), perfil, genero, ouvinte)
}

// UpdateGenerosFavoritos mocks base method
func (m *MockUseCase) UpdateGenerosFavoritos(e *generosfavoritos.GenerosFavoritos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGenerosFavoritos", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGenerosFavoritos indicates an expected call of UpdateGenerosFavoritos
func (mr *MockUseCaseMockRecorder) UpdateGenerosFavoritos(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGenerosFavoritos", reflect.TypeOf((*MockUseCase)(nil).UpdateGenerosFavoritos), e)
}

// DeleteGenerosFavoritos mocks base method
func (m *MockUseCase) DeleteGenerosFavoritos(perfil int, genero, ouvinte string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGenerosFavoritos", perfil, genero, ouvinte)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGenerosFavoritos indicates an expected call of DeleteGenerosFavoritos
func (mr *MockUseCaseMockRecorder) DeleteGenerosFavoritos(perfil, genero, ouvinte interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGenerosFavoritos", reflect.TypeOf((*MockUseCase)(nil).DeleteGenerosFavoritos), perfil, genero, ouvinte)
}