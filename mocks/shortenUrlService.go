// Code generated by MockGen. DO NOT EDIT.
// Source: service/shortenurlService.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	common "tinyurl_api/common"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockShortenUrlService is a mock of ShortenUrlService interface.
type MockShortenUrlService struct {
	ctrl     *gomock.Controller
	recorder *MockShortenUrlServiceMockRecorder
}

// MockShortenUrlServiceMockRecorder is the mock recorder for MockShortenUrlService.
type MockShortenUrlServiceMockRecorder struct {
	mock *MockShortenUrlService
}

// NewMockShortenUrlService creates a new mock instance.
func NewMockShortenUrlService(ctrl *gomock.Controller) *MockShortenUrlService {
	mock := &MockShortenUrlService{ctrl: ctrl}
	mock.recorder = &MockShortenUrlServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShortenUrlService) EXPECT() *MockShortenUrlServiceMockRecorder {
	return m.recorder
}

// CreateShortenUrl mocks base method.
func (m *MockShortenUrlService) CreateShortenUrl(ctx *gin.Context) (common.Response, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShortenUrl", ctx)
	ret0, _ := ret[0].(common.Response)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// CreateShortenUrl indicates an expected call of CreateShortenUrl.
func (mr *MockShortenUrlServiceMockRecorder) CreateShortenUrl(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShortenUrl", reflect.TypeOf((*MockShortenUrlService)(nil).CreateShortenUrl), ctx)
}
