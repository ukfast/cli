// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ukfast/sdk-go/pkg/service/storage (interfaces: StorageService)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	connection "github.com/ukfast/sdk-go/pkg/connection"
	storage "github.com/ukfast/sdk-go/pkg/service/storage"
	reflect "reflect"
)

// MockStorageService is a mock of StorageService interface
type MockStorageService struct {
	ctrl     *gomock.Controller
	recorder *MockStorageServiceMockRecorder
}

// MockStorageServiceMockRecorder is the mock recorder for MockStorageService
type MockStorageServiceMockRecorder struct {
	mock *MockStorageService
}

// NewMockStorageService creates a new mock instance
func NewMockStorageService(ctrl *gomock.Controller) *MockStorageService {
	mock := &MockStorageService{ctrl: ctrl}
	mock.recorder = &MockStorageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageService) EXPECT() *MockStorageServiceMockRecorder {
	return m.recorder
}

// GetHost mocks base method
func (m *MockStorageService) GetHost(arg0 int) (storage.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHost", arg0)
	ret0, _ := ret[0].(storage.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHost indicates an expected call of GetHost
func (mr *MockStorageServiceMockRecorder) GetHost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHost", reflect.TypeOf((*MockStorageService)(nil).GetHost), arg0)
}

// GetHosts mocks base method
func (m *MockStorageService) GetHosts(arg0 connection.APIRequestParameters) ([]storage.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHosts", arg0)
	ret0, _ := ret[0].([]storage.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHosts indicates an expected call of GetHosts
func (mr *MockStorageServiceMockRecorder) GetHosts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHosts", reflect.TypeOf((*MockStorageService)(nil).GetHosts), arg0)
}

// GetHostsPaginated mocks base method
func (m *MockStorageService) GetHostsPaginated(arg0 connection.APIRequestParameters) (*storage.PaginatedHost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostsPaginated", arg0)
	ret0, _ := ret[0].(*storage.PaginatedHost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostsPaginated indicates an expected call of GetHostsPaginated
func (mr *MockStorageServiceMockRecorder) GetHostsPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostsPaginated", reflect.TypeOf((*MockStorageService)(nil).GetHostsPaginated), arg0)
}

// GetSolution mocks base method
func (m *MockStorageService) GetSolution(arg0 int) (storage.Solution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolution", arg0)
	ret0, _ := ret[0].(storage.Solution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolution indicates an expected call of GetSolution
func (mr *MockStorageServiceMockRecorder) GetSolution(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolution", reflect.TypeOf((*MockStorageService)(nil).GetSolution), arg0)
}

// GetSolutions mocks base method
func (m *MockStorageService) GetSolutions(arg0 connection.APIRequestParameters) ([]storage.Solution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutions", arg0)
	ret0, _ := ret[0].([]storage.Solution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutions indicates an expected call of GetSolutions
func (mr *MockStorageServiceMockRecorder) GetSolutions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutions", reflect.TypeOf((*MockStorageService)(nil).GetSolutions), arg0)
}

// GetSolutionsPaginated mocks base method
func (m *MockStorageService) GetSolutionsPaginated(arg0 connection.APIRequestParameters) (*storage.PaginatedSolution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionsPaginated", arg0)
	ret0, _ := ret[0].(*storage.PaginatedSolution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionsPaginated indicates an expected call of GetSolutionsPaginated
func (mr *MockStorageServiceMockRecorder) GetSolutionsPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionsPaginated", reflect.TypeOf((*MockStorageService)(nil).GetSolutionsPaginated), arg0)
}

// GetVolume mocks base method
func (m *MockStorageService) GetVolume(arg0 int) (storage.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolume", arg0)
	ret0, _ := ret[0].(storage.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolume indicates an expected call of GetVolume
func (mr *MockStorageServiceMockRecorder) GetVolume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolume", reflect.TypeOf((*MockStorageService)(nil).GetVolume), arg0)
}

// GetVolumes mocks base method
func (m *MockStorageService) GetVolumes(arg0 connection.APIRequestParameters) ([]storage.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumes", arg0)
	ret0, _ := ret[0].([]storage.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumes indicates an expected call of GetVolumes
func (mr *MockStorageServiceMockRecorder) GetVolumes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumes", reflect.TypeOf((*MockStorageService)(nil).GetVolumes), arg0)
}

// GetVolumesPaginated mocks base method
func (m *MockStorageService) GetVolumesPaginated(arg0 connection.APIRequestParameters) (*storage.PaginatedVolume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumesPaginated", arg0)
	ret0, _ := ret[0].(*storage.PaginatedVolume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumesPaginated indicates an expected call of GetVolumesPaginated
func (mr *MockStorageServiceMockRecorder) GetVolumesPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumesPaginated", reflect.TypeOf((*MockStorageService)(nil).GetVolumesPaginated), arg0)
}