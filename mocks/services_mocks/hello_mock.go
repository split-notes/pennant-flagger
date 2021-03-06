// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/hello/interface.go

// Package services_mocks is a generated GoMock package.
package services_mocks

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/split-notes/pennant-flagger/db/models"
	reflect "reflect"
)

// MockHello is a mock of Service interface
type MockHello struct {
	ctrl     *gomock.Controller
	recorder *MockHelloMockRecorder
}

// MockHelloMockRecorder is the mock recorder for MockHello
type MockHelloMockRecorder struct {
	mock *MockHello
}

// NewMockHello creates a new mock instance
func NewMockHello(ctrl *gomock.Controller) *MockHello {
	mock := &MockHello{ctrl: ctrl}
	mock.recorder = &MockHelloMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHello) EXPECT() *MockHelloMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockHello) Create(greetingModel models.Greetings) (*models.Greetings, error) {
	ret := m.ctrl.Call(m, "Create", greetingModel)
	ret0, _ := ret[0].(*models.Greetings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockHelloMockRecorder) Create(greetingModel interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockHello)(nil).Create), greetingModel)
}

// Get mocks base method
func (m *MockHello) Get() ([]models.Greetings, error) {
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].([]models.Greetings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockHelloMockRecorder) Get() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHello)(nil).Get))
}
