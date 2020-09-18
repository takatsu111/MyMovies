// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/CommentRepository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "MyPIPE/domain/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCommentRepository is a mock of CommentRepository interface
type MockCommentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCommentRepositoryMockRecorder
}

// MockCommentRepositoryMockRecorder is the mock recorder for MockCommentRepository
type MockCommentRepositoryMockRecorder struct {
	mock *MockCommentRepository
}

// NewMockCommentRepository creates a new mock instance
func NewMockCommentRepository(ctrl *gomock.Controller) *MockCommentRepository {
	mock := &MockCommentRepository{ctrl: ctrl}
	mock.recorder = &MockCommentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommentRepository) EXPECT() *MockCommentRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockCommentRepository) GetAll() ([]model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockCommentRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCommentRepository)(nil).GetAll))
}

// FindById mocks base method
func (m *MockCommentRepository) FindById(commentID model.CommentID) (*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", commentID)
	ret0, _ := ret[0].(*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById
func (mr *MockCommentRepositoryMockRecorder) FindById(commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockCommentRepository)(nil).FindById), commentID)
}

// FindByUserId mocks base method
func (m *MockCommentRepository) FindByUserId(userId model.UserID) ([]model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserId", userId)
	ret0, _ := ret[0].([]model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserId indicates an expected call of FindByUserId
func (mr *MockCommentRepositoryMockRecorder) FindByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserId", reflect.TypeOf((*MockCommentRepository)(nil).FindByUserId), userId)
}

// FindByMovieId mocks base method
func (m *MockCommentRepository) FindByMovieId(movieId model.MovieID) ([]model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMovieId", movieId)
	ret0, _ := ret[0].([]model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMovieId indicates an expected call of FindByMovieId
func (mr *MockCommentRepositoryMockRecorder) FindByMovieId(movieId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMovieId", reflect.TypeOf((*MockCommentRepository)(nil).FindByMovieId), movieId)
}

// Save mocks base method
func (m *MockCommentRepository) Save(arg0 *model.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockCommentRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCommentRepository)(nil).Save), arg0)
}
