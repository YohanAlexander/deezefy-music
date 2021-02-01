// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/playlist/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	playlist "github.com/yohanalexander/deezefy-music/entity/playlist"
	reflect "reflect"
)

// MockPlaylist is a mock of Playlist interface
type MockPlaylist struct {
	ctrl     *gomock.Controller
	recorder *MockPlaylistMockRecorder
}

// MockPlaylistMockRecorder is the mock recorder for MockPlaylist
type MockPlaylistMockRecorder struct {
	mock *MockPlaylist
}

// NewMockPlaylist creates a new mock instance
func NewMockPlaylist(ctrl *gomock.Controller) *MockPlaylist {
	mock := &MockPlaylist{ctrl: ctrl}
	mock.recorder = &MockPlaylistMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlaylist) EXPECT() *MockPlaylistMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockPlaylist) Get(nome string) (*playlist.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", nome)
	ret0, _ := ret[0].(*playlist.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPlaylistMockRecorder) Get(nome interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPlaylist)(nil).Get), nome)
}

// Search mocks base method
func (m *MockPlaylist) Search(query string) ([]*playlist.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*playlist.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockPlaylistMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockPlaylist)(nil).Search), query)
}

// List mocks base method
func (m *MockPlaylist) List() ([]*playlist.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*playlist.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockPlaylistMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPlaylist)(nil).List))
}

// Create mocks base method
func (m *MockPlaylist) Create(e *playlist.Playlist) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockPlaylistMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPlaylist)(nil).Create), e)
}

// Update mocks base method
func (m *MockPlaylist) Update(e *playlist.Playlist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockPlaylistMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPlaylist)(nil).Update), e)
}

// Delete mocks base method
func (m *MockPlaylist) Delete(nome string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", nome)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPlaylistMockRecorder) Delete(nome interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPlaylist)(nil).Delete), nome)
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
func (m *MockRepository) Get(nome string) (*playlist.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", nome)
	ret0, _ := ret[0].(*playlist.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(nome interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), nome)
}

// Search mocks base method
func (m *MockRepository) Search(query string) ([]*playlist.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*playlist.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockRepositoryMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), query)
}

// List mocks base method
func (m *MockRepository) List() ([]*playlist.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*playlist.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List))
}

// Create mocks base method
func (m *MockRepository) Create(e *playlist.Playlist) (string, error) {
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
func (m *MockRepository) Update(e *playlist.Playlist) error {
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
func (m *MockRepository) Delete(nome string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", nome)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(nome interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), nome)
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

// GetPlaylist mocks base method
func (m *MockUseCase) GetPlaylist(nome string) (*playlist.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaylist", nome)
	ret0, _ := ret[0].(*playlist.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaylist indicates an expected call of GetPlaylist
func (mr *MockUseCaseMockRecorder) GetPlaylist(nome interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaylist", reflect.TypeOf((*MockUseCase)(nil).GetPlaylist), nome)
}

// SearchPlaylists mocks base method
func (m *MockUseCase) SearchPlaylists(query string) ([]*playlist.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPlaylists", query)
	ret0, _ := ret[0].([]*playlist.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPlaylists indicates an expected call of SearchPlaylists
func (mr *MockUseCaseMockRecorder) SearchPlaylists(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPlaylists", reflect.TypeOf((*MockUseCase)(nil).SearchPlaylists), query)
}

// ListPlaylists mocks base method
func (m *MockUseCase) ListPlaylists() ([]*playlist.Playlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPlaylists")
	ret0, _ := ret[0].([]*playlist.Playlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPlaylists indicates an expected call of ListPlaylists
func (mr *MockUseCaseMockRecorder) ListPlaylists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPlaylists", reflect.TypeOf((*MockUseCase)(nil).ListPlaylists))
}

// CreatePlaylist mocks base method
func (m *MockUseCase) CreatePlaylist(nome, status string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlaylist", nome, status)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePlaylist indicates an expected call of CreatePlaylist
func (mr *MockUseCaseMockRecorder) CreatePlaylist(nome, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlaylist", reflect.TypeOf((*MockUseCase)(nil).CreatePlaylist), nome, status)
}

// UpdatePlaylist mocks base method
func (m *MockUseCase) UpdatePlaylist(e *playlist.Playlist) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlaylist", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePlaylist indicates an expected call of UpdatePlaylist
func (mr *MockUseCaseMockRecorder) UpdatePlaylist(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlaylist", reflect.TypeOf((*MockUseCase)(nil).UpdatePlaylist), e)
}

// DeletePlaylist mocks base method
func (m *MockUseCase) DeletePlaylist(nome string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlaylist", nome)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlaylist indicates an expected call of DeletePlaylist
func (mr *MockUseCaseMockRecorder) DeletePlaylist(nome interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlaylist", reflect.TypeOf((*MockUseCase)(nil).DeletePlaylist), nome)
}
