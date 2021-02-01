// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/artista/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	artista "github.com/yohanalexander/deezefy-music/entity/artista"
	reflect "reflect"
)

// MockArtista is a mock of Artista interface
type MockArtista struct {
	ctrl     *gomock.Controller
	recorder *MockArtistaMockRecorder
}

// MockArtistaMockRecorder is the mock recorder for MockArtista
type MockArtistaMockRecorder struct {
	mock *MockArtista
}

// NewMockArtista creates a new mock instance
func NewMockArtista(ctrl *gomock.Controller) *MockArtista {
	mock := &MockArtista{ctrl: ctrl}
	mock.recorder = &MockArtistaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockArtista) EXPECT() *MockArtistaMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockArtista) Get(email string) (*artista.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", email)
	ret0, _ := ret[0].(*artista.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockArtistaMockRecorder) Get(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockArtista)(nil).Get), email)
}

// Search mocks base method
func (m *MockArtista) Search(query string) ([]*artista.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*artista.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockArtistaMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockArtista)(nil).Search), query)
}

// List mocks base method
func (m *MockArtista) List() ([]*artista.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*artista.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockArtistaMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockArtista)(nil).List))
}

// Create mocks base method
func (m *MockArtista) Create(e *artista.Artista) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockArtistaMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArtista)(nil).Create), e)
}

// Update mocks base method
func (m *MockArtista) Update(e *artista.Artista) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockArtistaMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArtista)(nil).Update), e)
}

// Delete mocks base method
func (m *MockArtista) Delete(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockArtistaMockRecorder) Delete(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArtista)(nil).Delete), email)
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
func (m *MockRepository) Get(email string) (*artista.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", email)
	ret0, _ := ret[0].(*artista.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), email)
}

// Search mocks base method
func (m *MockRepository) Search(query string) ([]*artista.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*artista.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockRepositoryMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), query)
}

// List mocks base method
func (m *MockRepository) List() ([]*artista.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*artista.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List))
}

// Create mocks base method
func (m *MockRepository) Create(e *artista.Artista) (string, error) {
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
func (m *MockRepository) Update(e *artista.Artista) error {
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

// GetArtista mocks base method
func (m *MockUseCase) GetArtista(email string) (*artista.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArtista", email)
	ret0, _ := ret[0].(*artista.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArtista indicates an expected call of GetArtista
func (mr *MockUseCaseMockRecorder) GetArtista(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArtista", reflect.TypeOf((*MockUseCase)(nil).GetArtista), email)
}

// SearchArtistas mocks base method
func (m *MockUseCase) SearchArtistas(query string) ([]*artista.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchArtistas", query)
	ret0, _ := ret[0].([]*artista.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchArtistas indicates an expected call of SearchArtistas
func (mr *MockUseCaseMockRecorder) SearchArtistas(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchArtistas", reflect.TypeOf((*MockUseCase)(nil).SearchArtistas), query)
}

// ListArtistas mocks base method
func (m *MockUseCase) ListArtistas() ([]*artista.Artista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListArtistas")
	ret0, _ := ret[0].([]*artista.Artista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListArtistas indicates an expected call of ListArtistas
func (mr *MockUseCaseMockRecorder) ListArtistas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListArtistas", reflect.TypeOf((*MockUseCase)(nil).ListArtistas))
}

// CreateArtista mocks base method
func (m *MockUseCase) CreateArtista(usuario, nomeartistico, biografia string, anoformacao int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArtista", usuario, nomeartistico, biografia, anoformacao)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArtista indicates an expected call of CreateArtista
func (mr *MockUseCaseMockRecorder) CreateArtista(usuario, nomeartistico, biografia, anoformacao interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArtista", reflect.TypeOf((*MockUseCase)(nil).CreateArtista), usuario, nomeartistico, biografia, anoformacao)
}

// UpdateArtista mocks base method
func (m *MockUseCase) UpdateArtista(e *artista.Artista) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArtista", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateArtista indicates an expected call of UpdateArtista
func (mr *MockUseCaseMockRecorder) UpdateArtista(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArtista", reflect.TypeOf((*MockUseCase)(nil).UpdateArtista), e)
}

// DeleteArtista mocks base method
func (m *MockUseCase) DeleteArtista(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArtista", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteArtista indicates an expected call of DeleteArtista
func (mr *MockUseCaseMockRecorder) DeleteArtista(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArtista", reflect.TypeOf((*MockUseCase)(nil).DeleteArtista), email)
}