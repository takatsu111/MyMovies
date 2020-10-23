// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/UpdateMovie.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	model "MyPIPE/domain/model"
	usecase "MyPIPE/usecase"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIUpdateMovie is a mock of IUpdateMovie interface
type MockIUpdateMovie struct {
	ctrl     *gomock.Controller
	recorder *MockIUpdateMovieMockRecorder
}

// MockIUpdateMovieMockRecorder is the mock recorder for MockIUpdateMovie
type MockIUpdateMovieMockRecorder struct {
	mock *MockIUpdateMovie
}

// NewMockIUpdateMovie creates a new mock instance
func NewMockIUpdateMovie(ctrl *gomock.Controller) *MockIUpdateMovie {
	mock := &MockIUpdateMovie{ctrl: ctrl}
	mock.recorder = &MockIUpdateMovieMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUpdateMovie) EXPECT() *MockIUpdateMovieMockRecorder {
	return m.recorder
}

// Update mocks base method
func (m *MockIUpdateMovie) Update(updateDTO *usecase.UpdateDTO) (*model.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", updateDTO)
	ret0, _ := ret[0].(*model.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockIUpdateMovieMockRecorder) Update(updateDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUpdateMovie)(nil).Update), updateDTO)
}

// UpdateStatus mocks base method
func (m *MockIUpdateMovie) UpdateStatus(updateStatusDTO usecase.UpdateStatusDTO) (*model.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", updateStatusDTO)
	ret0, _ := ret[0].(*model.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockIUpdateMovieMockRecorder) UpdateStatus(updateStatusDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockIUpdateMovie)(nil).UpdateStatus), updateStatusDTO)
}
