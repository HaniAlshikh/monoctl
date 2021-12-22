// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/finleap-connect/monoskope/pkg/api/domain (interfaces: CertificateClient)

// Package domain is a generated GoMock package.
package domain

import (
	context "context"
	reflect "reflect"

	domain "github.com/finleap-connect/monoskope/pkg/api/domain"
	projections "github.com/finleap-connect/monoskope/pkg/api/domain/projections"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockCertificateClient is a mock of CertificateClient interface.
type MockCertificateClient struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateClientMockRecorder
}

// MockCertificateClientMockRecorder is the mock recorder for MockCertificateClient.
type MockCertificateClientMockRecorder struct {
	mock *MockCertificateClient
}

// NewMockCertificateClient creates a new mock instance.
func NewMockCertificateClient(ctrl *gomock.Controller) *MockCertificateClient {
	mock := &MockCertificateClient{ctrl: ctrl}
	mock.recorder = &MockCertificateClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCertificateClient) EXPECT() *MockCertificateClientMockRecorder {
	return m.recorder
}

// GetCertificate mocks base method.
func (m *MockCertificateClient) GetCertificate(arg0 context.Context, arg1 *domain.GetCertificateRequest, arg2 ...grpc.CallOption) (*projections.Certificate, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCertificate", varargs...)
	ret0, _ := ret[0].(*projections.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificate indicates an expected call of GetCertificate.
func (mr *MockCertificateClientMockRecorder) GetCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificate", reflect.TypeOf((*MockCertificateClient)(nil).GetCertificate), varargs...)
}
