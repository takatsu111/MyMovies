// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/repository/MovieStatusRepository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "MyPIPE/domain/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMovieStatusRepository is a mock of MovieStatusRepository interface
type MockMovieStatusRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMovieStatusRepositoryMockRecorder
}

// MockMovieStatusRepositoryMockRecorder is the mock recorder for MockMovieStatusRepository
type MockMovieStatusRepositoryMockRecorder struct {
	mock *MockMovieStatusRepository
}

// NewMockMovieStatusRepository creates a new mock instance
func NewMockMovieStatusRepository(ctrl *gomock.Controller) *MockMovieStatusRepository {
	mock := &MockMovieStatusRepository{ctrl: ctrl}
	mock.recorder = &MockMovieStatusRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMovieStatusRepository) EXPECT() *MockMovieStatusRepositoryMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockMovieStatusRepository) Find(movieId model.MovieID) (*model.MovieStatusModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", movieId)
	ret0, _ := ret[0].(*model.MovieStatusModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockMovieStatusRepositoryMockRecorder) Find(movieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMovieStatusRepository)(nil).Find), movieId)
}

// Save mocks base method
func (m *MockMovieStatusRepository) Save(movieStatusModel *model.MovieStatusModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", movieStatusModel)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockMovieStatusRepositoryMockRecorder) Save(movieStatusModel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockMovieStatusRepository)(nil).Save), movieStatusModel)
}
