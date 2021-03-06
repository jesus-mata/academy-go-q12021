// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jesus-mata/academy-go-q12021/application/repository (interfaces: NewsArticleRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/jesus-mata/academy-go-q12021/domain"
)

// MockNewsArticleRepository is a mock of NewsArticleRepository interface.
type MockNewsArticleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockNewsArticleRepositoryMockRecorder
}

// MockNewsArticleRepositoryMockRecorder is the mock recorder for MockNewsArticleRepository.
type MockNewsArticleRepositoryMockRecorder struct {
	mock *MockNewsArticleRepository
}

// NewMockNewsArticleRepository creates a new mock instance.
func NewMockNewsArticleRepository(ctrl *gomock.Controller) *MockNewsArticleRepository {
	mock := &MockNewsArticleRepository{ctrl: ctrl}
	mock.recorder = &MockNewsArticleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNewsArticleRepository) EXPECT() *MockNewsArticleRepositoryMockRecorder {
	return m.recorder
}

// FetchCurrent mocks base method.
func (m *MockNewsArticleRepository) FetchCurrent() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCurrent")
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchCurrent indicates an expected call of FetchCurrent.
func (mr *MockNewsArticleRepositoryMockRecorder) FetchCurrent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCurrent", reflect.TypeOf((*MockNewsArticleRepository)(nil).FetchCurrent))
}

// FindAll mocks base method.
func (m *MockNewsArticleRepository) FindAll() ([]*domain.NewsArticle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*domain.NewsArticle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockNewsArticleRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockNewsArticleRepository)(nil).FindAll))
}

// FindByID mocks base method.
func (m *MockNewsArticleRepository) FindByID(arg0 string) (*domain.NewsArticle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(*domain.NewsArticle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockNewsArticleRepositoryMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockNewsArticleRepository)(nil).FindByID), arg0)
}

// GetIterator mocks base method.
func (m *MockNewsArticleRepository) GetIterator() (domain.NewsIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIterator")
	ret0, _ := ret[0].(domain.NewsIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIterator indicates an expected call of GetIterator.
func (mr *MockNewsArticleRepositoryMockRecorder) GetIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIterator", reflect.TypeOf((*MockNewsArticleRepository)(nil).GetIterator))
}
