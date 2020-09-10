// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ukfast/sdk-go/pkg/service/ssl (interfaces: SSLService)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	connection "github.com/ukfast/sdk-go/pkg/connection"
	ssl "github.com/ukfast/sdk-go/pkg/service/ssl"
	reflect "reflect"
)

// MockSSLService is a mock of SSLService interface
type MockSSLService struct {
	ctrl     *gomock.Controller
	recorder *MockSSLServiceMockRecorder
}

// MockSSLServiceMockRecorder is the mock recorder for MockSSLService
type MockSSLServiceMockRecorder struct {
	mock *MockSSLService
}

// NewMockSSLService creates a new mock instance
func NewMockSSLService(ctrl *gomock.Controller) *MockSSLService {
	mock := &MockSSLService{ctrl: ctrl}
	mock.recorder = &MockSSLServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSSLService) EXPECT() *MockSSLServiceMockRecorder {
	return m.recorder
}

// GetCertificate mocks base method
func (m *MockSSLService) GetCertificate(arg0 int) (ssl.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificate", arg0)
	ret0, _ := ret[0].(ssl.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificate indicates an expected call of GetCertificate
func (mr *MockSSLServiceMockRecorder) GetCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificate", reflect.TypeOf((*MockSSLService)(nil).GetCertificate), arg0)
}

// GetCertificateContent mocks base method
func (m *MockSSLService) GetCertificateContent(arg0 int) (ssl.CertificateContent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificateContent", arg0)
	ret0, _ := ret[0].(ssl.CertificateContent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificateContent indicates an expected call of GetCertificateContent
func (mr *MockSSLServiceMockRecorder) GetCertificateContent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateContent", reflect.TypeOf((*MockSSLService)(nil).GetCertificateContent), arg0)
}

// GetCertificatePrivateKey mocks base method
func (m *MockSSLService) GetCertificatePrivateKey(arg0 int) (ssl.CertificatePrivateKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificatePrivateKey", arg0)
	ret0, _ := ret[0].(ssl.CertificatePrivateKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificatePrivateKey indicates an expected call of GetCertificatePrivateKey
func (mr *MockSSLServiceMockRecorder) GetCertificatePrivateKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificatePrivateKey", reflect.TypeOf((*MockSSLService)(nil).GetCertificatePrivateKey), arg0)
}

// GetCertificates mocks base method
func (m *MockSSLService) GetCertificates(arg0 connection.APIRequestParameters) ([]ssl.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificates", arg0)
	ret0, _ := ret[0].([]ssl.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificates indicates an expected call of GetCertificates
func (mr *MockSSLServiceMockRecorder) GetCertificates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificates", reflect.TypeOf((*MockSSLService)(nil).GetCertificates), arg0)
}

// GetCertificatesPaginated mocks base method
func (m *MockSSLService) GetCertificatesPaginated(arg0 connection.APIRequestParameters) (*ssl.PaginatedCertificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificatesPaginated", arg0)
	ret0, _ := ret[0].(*ssl.PaginatedCertificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificatesPaginated indicates an expected call of GetCertificatesPaginated
func (mr *MockSSLServiceMockRecorder) GetCertificatesPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificatesPaginated", reflect.TypeOf((*MockSSLService)(nil).GetCertificatesPaginated), arg0)
}

// GetRecommendations mocks base method
func (m *MockSSLService) GetRecommendations(arg0 string) (ssl.Recommendations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecommendations", arg0)
	ret0, _ := ret[0].(ssl.Recommendations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecommendations indicates an expected call of GetRecommendations
func (mr *MockSSLServiceMockRecorder) GetRecommendations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendations", reflect.TypeOf((*MockSSLService)(nil).GetRecommendations), arg0)
}

// GetReport mocks base method
func (m *MockSSLService) GetReport(arg0 string) (ssl.Report, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReport", arg0)
	ret0, _ := ret[0].(ssl.Report)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReport indicates an expected call of GetReport
func (mr *MockSSLServiceMockRecorder) GetReport(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReport", reflect.TypeOf((*MockSSLService)(nil).GetReport), arg0)
}

// ValidateCertificate mocks base method
func (m *MockSSLService) ValidateCertificate(arg0 ssl.ValidateRequest) (ssl.CertificateValidation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateCertificate", arg0)
	ret0, _ := ret[0].(ssl.CertificateValidation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateCertificate indicates an expected call of ValidateCertificate
func (mr *MockSSLServiceMockRecorder) ValidateCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateCertificate", reflect.TypeOf((*MockSSLService)(nil).ValidateCertificate), arg0)
}
