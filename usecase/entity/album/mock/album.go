// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/entity/album/interface.go

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
func (m *MockRead) Get(id int) (*entity.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockReadMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRead)(nil).Get), id)
}

// Search mocks base method
func (m *MockRead) Search(query string) ([]*entity.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*entity.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockReadMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRead)(nil).Search), query)
}

// List mocks base method
func (m *MockRead) List() ([]*entity.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Album)
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
func (m *MockWrite) Create(e *entity.Album) (int, error) {
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
func (m *MockWrite) Update(e *entity.Album) error {
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
func (m *MockRepository) Get(id int) (*entity.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), id)
}

// Search mocks base method
func (m *MockRepository) Search(query string) ([]*entity.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*entity.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockRepositoryMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), query)
}

// List mocks base method
func (m *MockRepository) List() ([]*entity.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*entity.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List))
}

// Create mocks base method
func (m *MockRepository) Create(e *entity.Album) (int, error) {
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
func (m *MockRepository) Update(e *entity.Album) error {
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

// GetAlbum mocks base method
func (m *MockUseCase) GetAlbum(id int) (*entity.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbum", id)
	ret0, _ := ret[0].(*entity.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlbum indicates an expected call of GetAlbum
func (mr *MockUseCaseMockRecorder) GetAlbum(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbum", reflect.TypeOf((*MockUseCase)(nil).GetAlbum), id)
}

// SearchAlbums mocks base method
func (m *MockUseCase) SearchAlbums(query string) ([]*entity.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchAlbums", query)
	ret0, _ := ret[0].([]*entity.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAlbums indicates an expected call of SearchAlbums
func (mr *MockUseCaseMockRecorder) SearchAlbums(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAlbums", reflect.TypeOf((*MockUseCase)(nil).SearchAlbums), query)
}

// ListAlbums mocks base method
func (m *MockUseCase) ListAlbums() ([]*entity.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAlbums")
	ret0, _ := ret[0].([]*entity.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAlbums indicates an expected call of ListAlbums
func (mr *MockUseCaseMockRecorder) ListAlbums() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlbums", reflect.TypeOf((*MockUseCase)(nil).ListAlbums))
}

// CreateAlbum mocks base method
func (m *MockUseCase) CreateAlbum(email, password, birthday, nomeartistico, biografia, titulo string, anoformacao, anolancamento, id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAlbum", email, password, birthday, nomeartistico, biografia, titulo, anoformacao, anolancamento, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAlbum indicates an expected call of CreateAlbum
func (mr *MockUseCaseMockRecorder) CreateAlbum(email, password, birthday, nomeartistico, biografia, titulo, anoformacao, anolancamento, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAlbum", reflect.TypeOf((*MockUseCase)(nil).CreateAlbum), email, password, birthday, nomeartistico, biografia, titulo, anoformacao, anolancamento, id)
}

// UpdateAlbum mocks base method
func (m *MockUseCase) UpdateAlbum(e *entity.Album) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAlbum", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAlbum indicates an expected call of UpdateAlbum
func (mr *MockUseCaseMockRecorder) UpdateAlbum(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlbum", reflect.TypeOf((*MockUseCase)(nil).UpdateAlbum), e)
}

// DeleteAlbum mocks base method
func (m *MockUseCase) DeleteAlbum(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAlbum", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAlbum indicates an expected call of DeleteAlbum
func (mr *MockUseCaseMockRecorder) DeleteAlbum(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlbum", reflect.TypeOf((*MockUseCase)(nil).DeleteAlbum), id)
}
