// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/albummusica/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	albummusica "github.com/yohanalexander/deezefy-music/entity/albummusica"
	reflect "reflect"
)

// MockAlbumMusica is a mock of AlbumMusica interface
type MockAlbumMusica struct {
	ctrl     *gomock.Controller
	recorder *MockAlbumMusicaMockRecorder
}

// MockAlbumMusicaMockRecorder is the mock recorder for MockAlbumMusica
type MockAlbumMusicaMockRecorder struct {
	mock *MockAlbumMusica
}

// NewMockAlbumMusica creates a new mock instance
func NewMockAlbumMusica(ctrl *gomock.Controller) *MockAlbumMusica {
	mock := &MockAlbumMusica{ctrl: ctrl}
	mock.recorder = &MockAlbumMusicaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAlbumMusica) EXPECT() *MockAlbumMusicaMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockAlbumMusica) Get(album, musica int, artista string) (*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", album, musica, artista)
	ret0, _ := ret[0].(*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockAlbumMusicaMockRecorder) Get(album, musica, artista interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAlbumMusica)(nil).Get), album, musica, artista)
}

// GetByAlbum mocks base method
func (m *MockAlbumMusica) GetByAlbum(album int) (*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAlbum", album)
	ret0, _ := ret[0].(*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAlbum indicates an expected call of GetByAlbum
func (mr *MockAlbumMusicaMockRecorder) GetByAlbum(album interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAlbum", reflect.TypeOf((*MockAlbumMusica)(nil).GetByAlbum), album)
}

// GetByMusica mocks base method
func (m *MockAlbumMusica) GetByMusica(musica int) (*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMusica", musica)
	ret0, _ := ret[0].(*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMusica indicates an expected call of GetByMusica
func (mr *MockAlbumMusicaMockRecorder) GetByMusica(musica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMusica", reflect.TypeOf((*MockAlbumMusica)(nil).GetByMusica), musica)
}

// Search mocks base method
func (m *MockAlbumMusica) Search(query string) ([]*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockAlbumMusicaMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockAlbumMusica)(nil).Search), query)
}

// List mocks base method
func (m *MockAlbumMusica) List() ([]*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockAlbumMusicaMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAlbumMusica)(nil).List))
}

// Create mocks base method
func (m *MockAlbumMusica) Create(e *albummusica.AlbumMusica) (int, int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Create indicates an expected call of Create
func (mr *MockAlbumMusicaMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAlbumMusica)(nil).Create), e)
}

// Update mocks base method
func (m *MockAlbumMusica) Update(e *albummusica.AlbumMusica) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockAlbumMusicaMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAlbumMusica)(nil).Update), e)
}

// Delete mocks base method
func (m *MockAlbumMusica) Delete(album, musica int, artista string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", album, musica, artista)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockAlbumMusicaMockRecorder) Delete(album, musica, artista interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAlbumMusica)(nil).Delete), album, musica, artista)
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
func (m *MockRepository) Get(album, musica int, artista string) (*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", album, musica, artista)
	ret0, _ := ret[0].(*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockRepositoryMockRecorder) Get(album, musica, artista interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), album, musica, artista)
}

// GetByAlbum mocks base method
func (m *MockRepository) GetByAlbum(album int) (*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAlbum", album)
	ret0, _ := ret[0].(*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAlbum indicates an expected call of GetByAlbum
func (mr *MockRepositoryMockRecorder) GetByAlbum(album interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAlbum", reflect.TypeOf((*MockRepository)(nil).GetByAlbum), album)
}

// GetByMusica mocks base method
func (m *MockRepository) GetByMusica(musica int) (*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMusica", musica)
	ret0, _ := ret[0].(*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMusica indicates an expected call of GetByMusica
func (mr *MockRepositoryMockRecorder) GetByMusica(musica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMusica", reflect.TypeOf((*MockRepository)(nil).GetByMusica), musica)
}

// Search mocks base method
func (m *MockRepository) Search(query string) ([]*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", query)
	ret0, _ := ret[0].([]*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockRepositoryMockRecorder) Search(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), query)
}

// List mocks base method
func (m *MockRepository) List() ([]*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List))
}

// Create mocks base method
func (m *MockRepository) Create(e *albummusica.AlbumMusica) (int, int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
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
func (m *MockRepository) Update(e *albummusica.AlbumMusica) error {
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
func (m *MockRepository) Delete(album, musica int, artista string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", album, musica, artista)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRepositoryMockRecorder) Delete(album, musica, artista interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), album, musica, artista)
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

// GetAlbumMusica mocks base method
func (m *MockUseCase) GetAlbumMusica(album, musica int, artista string) (*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbumMusica", album, musica, artista)
	ret0, _ := ret[0].(*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlbumMusica indicates an expected call of GetAlbumMusica
func (mr *MockUseCaseMockRecorder) GetAlbumMusica(album, musica, artista interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbumMusica", reflect.TypeOf((*MockUseCase)(nil).GetAlbumMusica), album, musica, artista)
}

// GetAlbumMusicaByAlbum mocks base method
func (m *MockUseCase) GetAlbumMusicaByAlbum(album int) (*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbumMusicaByAlbum", album)
	ret0, _ := ret[0].(*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlbumMusicaByAlbum indicates an expected call of GetAlbumMusicaByAlbum
func (mr *MockUseCaseMockRecorder) GetAlbumMusicaByAlbum(album interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbumMusicaByAlbum", reflect.TypeOf((*MockUseCase)(nil).GetAlbumMusicaByAlbum), album)
}

// GetAlbumMusicaByMusica mocks base method
func (m *MockUseCase) GetAlbumMusicaByMusica(musica int) (*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbumMusicaByMusica", musica)
	ret0, _ := ret[0].(*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlbumMusicaByMusica indicates an expected call of GetAlbumMusicaByMusica
func (mr *MockUseCaseMockRecorder) GetAlbumMusicaByMusica(musica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbumMusicaByMusica", reflect.TypeOf((*MockUseCase)(nil).GetAlbumMusicaByMusica), musica)
}

// SearchAlbumMusicas mocks base method
func (m *MockUseCase) SearchAlbumMusicas(query string) ([]*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchAlbumMusicas", query)
	ret0, _ := ret[0].([]*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAlbumMusicas indicates an expected call of SearchAlbumMusicas
func (mr *MockUseCaseMockRecorder) SearchAlbumMusicas(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAlbumMusicas", reflect.TypeOf((*MockUseCase)(nil).SearchAlbumMusicas), query)
}

// ListAlbumMusicas mocks base method
func (m *MockUseCase) ListAlbumMusicas() ([]*albummusica.AlbumMusica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAlbumMusicas")
	ret0, _ := ret[0].([]*albummusica.AlbumMusica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAlbumMusicas indicates an expected call of ListAlbumMusicas
func (mr *MockUseCaseMockRecorder) ListAlbumMusicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAlbumMusicas", reflect.TypeOf((*MockUseCase)(nil).ListAlbumMusicas))
}

// CreateAlbumMusica mocks base method
func (m *MockUseCase) CreateAlbumMusica(album, musica int, artista string) (int, int, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAlbumMusica", album, musica, artista)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// CreateAlbumMusica indicates an expected call of CreateAlbumMusica
func (mr *MockUseCaseMockRecorder) CreateAlbumMusica(album, musica, artista interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAlbumMusica", reflect.TypeOf((*MockUseCase)(nil).CreateAlbumMusica), album, musica, artista)
}

// UpdateAlbumMusica mocks base method
func (m *MockUseCase) UpdateAlbumMusica(e *albummusica.AlbumMusica) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAlbumMusica", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAlbumMusica indicates an expected call of UpdateAlbumMusica
func (mr *MockUseCaseMockRecorder) UpdateAlbumMusica(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAlbumMusica", reflect.TypeOf((*MockUseCase)(nil).UpdateAlbumMusica), e)
}

// DeleteAlbumMusica mocks base method
func (m *MockUseCase) DeleteAlbumMusica(album, musica int, artista string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAlbumMusica", album, musica, artista)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAlbumMusica indicates an expected call of DeleteAlbumMusica
func (mr *MockUseCaseMockRecorder) DeleteAlbumMusica(album, musica, artista interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlbumMusica", reflect.TypeOf((*MockUseCase)(nil).DeleteAlbumMusica), album, musica, artista)
}