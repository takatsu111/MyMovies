// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/queryService/IndexPlayListsInMyPage.go

// Package mock_queryService is a generated GoMock package.
package mock_queryService

import (
	model "MyPIPE/domain/model"
	queryService "MyPIPE/domain/queryService"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIndexPlayListsInMyPageQueryService is a mock of IndexPlayListsInMyPageQueryService interface
type MockIndexPlayListsInMyPageQueryService struct {
	ctrl     *gomock.Controller
	recorder *MockIndexPlayListsInMyPageQueryServiceMockRecorder
}

// MockIndexPlayListsInMyPageQueryServiceMockRecorder is the mock recorder for MockIndexPlayListsInMyPageQueryService
type MockIndexPlayListsInMyPageQueryServiceMockRecorder struct {
	mock *MockIndexPlayListsInMyPageQueryService
}

// NewMockIndexPlayListsInMyPageQueryService creates a new mock instance
func NewMockIndexPlayListsInMyPageQueryService(ctrl *gomock.Controller) *MockIndexPlayListsInMyPageQueryService {
	mock := &MockIndexPlayListsInMyPageQueryService{ctrl: ctrl}
	mock.recorder = &MockIndexPlayListsInMyPageQueryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexPlayListsInMyPageQueryService) EXPECT() *MockIndexPlayListsInMyPageQueryServiceMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockIndexPlayListsInMyPageQueryService) All(userId model.UserID) *queryService.IndexPlayListsInMyPageDTO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", userId)
	ret0, _ := ret[0].(*queryService.IndexPlayListsInMyPageDTO)
	return ret0
}

// All indicates an expected call of All
func (mr *MockIndexPlayListsInMyPageQueryServiceMockRecorder) All(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockIndexPlayListsInMyPageQueryService)(nil).All), userId)
}