// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ukfast/sdk-go/pkg/service/account (interfaces: AccountService)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	connection "github.com/ukfast/sdk-go/pkg/connection"
	account "github.com/ukfast/sdk-go/pkg/service/account"
	reflect "reflect"
)

// MockAccountService is a mock of AccountService interface
type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

// MockAccountServiceMockRecorder is the mock recorder for MockAccountService
type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

// NewMockAccountService creates a new mock instance
func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

// CreateInvoiceQuery mocks base method
func (m *MockAccountService) CreateInvoiceQuery(arg0 account.CreateInvoiceQueryRequest) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvoiceQuery", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInvoiceQuery indicates an expected call of CreateInvoiceQuery
func (mr *MockAccountServiceMockRecorder) CreateInvoiceQuery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvoiceQuery", reflect.TypeOf((*MockAccountService)(nil).CreateInvoiceQuery), arg0)
}

// GetContact mocks base method
func (m *MockAccountService) GetContact(arg0 int) (account.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContact", arg0)
	ret0, _ := ret[0].(account.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContact indicates an expected call of GetContact
func (mr *MockAccountServiceMockRecorder) GetContact(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContact", reflect.TypeOf((*MockAccountService)(nil).GetContact), arg0)
}

// GetContacts mocks base method
func (m *MockAccountService) GetContacts(arg0 connection.APIRequestParameters) ([]account.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContacts", arg0)
	ret0, _ := ret[0].([]account.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContacts indicates an expected call of GetContacts
func (mr *MockAccountServiceMockRecorder) GetContacts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContacts", reflect.TypeOf((*MockAccountService)(nil).GetContacts), arg0)
}

// GetContactsPaginated mocks base method
func (m *MockAccountService) GetContactsPaginated(arg0 connection.APIRequestParameters) (*account.PaginatedContact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContactsPaginated", arg0)
	ret0, _ := ret[0].(*account.PaginatedContact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactsPaginated indicates an expected call of GetContactsPaginated
func (mr *MockAccountServiceMockRecorder) GetContactsPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactsPaginated", reflect.TypeOf((*MockAccountService)(nil).GetContactsPaginated), arg0)
}

// GetCredits mocks base method
func (m *MockAccountService) GetCredits(arg0 connection.APIRequestParameters) ([]account.Credit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredits", arg0)
	ret0, _ := ret[0].([]account.Credit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredits indicates an expected call of GetCredits
func (mr *MockAccountServiceMockRecorder) GetCredits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredits", reflect.TypeOf((*MockAccountService)(nil).GetCredits), arg0)
}

// GetDetails mocks base method
func (m *MockAccountService) GetDetails() (account.Details, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetails")
	ret0, _ := ret[0].(account.Details)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetails indicates an expected call of GetDetails
func (mr *MockAccountServiceMockRecorder) GetDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetails", reflect.TypeOf((*MockAccountService)(nil).GetDetails))
}

// GetInvoice mocks base method
func (m *MockAccountService) GetInvoice(arg0 int) (account.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvoice", arg0)
	ret0, _ := ret[0].(account.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvoice indicates an expected call of GetInvoice
func (mr *MockAccountServiceMockRecorder) GetInvoice(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvoice", reflect.TypeOf((*MockAccountService)(nil).GetInvoice), arg0)
}

// GetInvoiceQueries mocks base method
func (m *MockAccountService) GetInvoiceQueries(arg0 connection.APIRequestParameters) ([]account.InvoiceQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvoiceQueries", arg0)
	ret0, _ := ret[0].([]account.InvoiceQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvoiceQueries indicates an expected call of GetInvoiceQueries
func (mr *MockAccountServiceMockRecorder) GetInvoiceQueries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvoiceQueries", reflect.TypeOf((*MockAccountService)(nil).GetInvoiceQueries), arg0)
}

// GetInvoiceQueriesPaginated mocks base method
func (m *MockAccountService) GetInvoiceQueriesPaginated(arg0 connection.APIRequestParameters) (*account.PaginatedInvoiceQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvoiceQueriesPaginated", arg0)
	ret0, _ := ret[0].(*account.PaginatedInvoiceQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvoiceQueriesPaginated indicates an expected call of GetInvoiceQueriesPaginated
func (mr *MockAccountServiceMockRecorder) GetInvoiceQueriesPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvoiceQueriesPaginated", reflect.TypeOf((*MockAccountService)(nil).GetInvoiceQueriesPaginated), arg0)
}

// GetInvoiceQuery mocks base method
func (m *MockAccountService) GetInvoiceQuery(arg0 int) (account.InvoiceQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvoiceQuery", arg0)
	ret0, _ := ret[0].(account.InvoiceQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvoiceQuery indicates an expected call of GetInvoiceQuery
func (mr *MockAccountServiceMockRecorder) GetInvoiceQuery(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvoiceQuery", reflect.TypeOf((*MockAccountService)(nil).GetInvoiceQuery), arg0)
}

// GetInvoices mocks base method
func (m *MockAccountService) GetInvoices(arg0 connection.APIRequestParameters) ([]account.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvoices", arg0)
	ret0, _ := ret[0].([]account.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvoices indicates an expected call of GetInvoices
func (mr *MockAccountServiceMockRecorder) GetInvoices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvoices", reflect.TypeOf((*MockAccountService)(nil).GetInvoices), arg0)
}

// GetInvoicesPaginated mocks base method
func (m *MockAccountService) GetInvoicesPaginated(arg0 connection.APIRequestParameters) (*account.PaginatedInvoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvoicesPaginated", arg0)
	ret0, _ := ret[0].(*account.PaginatedInvoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvoicesPaginated indicates an expected call of GetInvoicesPaginated
func (mr *MockAccountServiceMockRecorder) GetInvoicesPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvoicesPaginated", reflect.TypeOf((*MockAccountService)(nil).GetInvoicesPaginated), arg0)
}
