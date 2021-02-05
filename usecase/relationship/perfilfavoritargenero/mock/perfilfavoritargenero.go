// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/relationship/perfilfavoritargenero/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/yohanalexander/deezefy-music/entity"
	reflect "reflect"
)

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

// Favoritar mocks base method
func (m *MockUseCase) Favoritar(g *entity.Genero, p *entity.Perfil) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Favoritar", g, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// Favoritar indicates an expected call of Favoritar
func (mr *MockUseCaseMockRecorder) Favoritar(g, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Favoritar", reflect.TypeOf((*MockUseCase)(nil).Favoritar), g, p)
}

// Desfavoritar mocks base method
func (m *MockUseCase) Desfavoritar(g *entity.Genero, p *entity.Perfil) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Desfavoritar", g, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// Desfavoritar indicates an expected call of Desfavoritar
func (mr *MockUseCaseMockRecorder) Desfavoritar(g, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Desfavoritar", reflect.TypeOf((*MockUseCase)(nil).Desfavoritar), g, p)
}
