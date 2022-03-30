// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/finleap-connect/monoskope/pkg/api/domain (interfaces: AuditLogClient,AuditLog_GetByDateRangeClient,AuditLog_GetUserActionsClient)

// Package domain is a generated GoMock package.
package domain

import (
	context "context"
	reflect "reflect"

	domain "github.com/finleap-connect/monoskope/pkg/api/domain"
	audit "github.com/finleap-connect/monoskope/pkg/api/domain/audit"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockAuditLogClient is a mock of AuditLogClient interface.
type MockAuditLogClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuditLogClientMockRecorder
}

// MockAuditLogClientMockRecorder is the mock recorder for MockAuditLogClient.
type MockAuditLogClientMockRecorder struct {
	mock *MockAuditLogClient
}

// NewMockAuditLogClient creates a new mock instance.
func NewMockAuditLogClient(ctrl *gomock.Controller) *MockAuditLogClient {
	mock := &MockAuditLogClient{ctrl: ctrl}
	mock.recorder = &MockAuditLogClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditLogClient) EXPECT() *MockAuditLogClientMockRecorder {
	return m.recorder
}

// GetByDateRange mocks base method.
func (m *MockAuditLogClient) GetByDateRange(arg0 context.Context, arg1 *domain.GetAuditLogByDateRangeRequest, arg2 ...grpc.CallOption) (domain.AuditLog_GetByDateRangeClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByDateRange", varargs...)
	ret0, _ := ret[0].(domain.AuditLog_GetByDateRangeClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDateRange indicates an expected call of GetByDateRange.
func (mr *MockAuditLogClientMockRecorder) GetByDateRange(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDateRange", reflect.TypeOf((*MockAuditLogClient)(nil).GetByDateRange), varargs...)
}

// GetUserActions mocks base method.
func (m *MockAuditLogClient) GetUserActions(arg0 context.Context, arg1 *domain.GetUserActionsRequest, arg2 ...grpc.CallOption) (domain.AuditLog_GetUserActionsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserActions", varargs...)
	ret0, _ := ret[0].(domain.AuditLog_GetUserActionsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserActions indicates an expected call of GetUserActions.
func (mr *MockAuditLogClientMockRecorder) GetUserActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserActions", reflect.TypeOf((*MockAuditLogClient)(nil).GetUserActions), varargs...)
}

// MockAuditLog_GetByDateRangeClient is a mock of AuditLog_GetByDateRangeClient interface.
type MockAuditLog_GetByDateRangeClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuditLog_GetByDateRangeClientMockRecorder
}

// MockAuditLog_GetByDateRangeClientMockRecorder is the mock recorder for MockAuditLog_GetByDateRangeClient.
type MockAuditLog_GetByDateRangeClientMockRecorder struct {
	mock *MockAuditLog_GetByDateRangeClient
}

// NewMockAuditLog_GetByDateRangeClient creates a new mock instance.
func NewMockAuditLog_GetByDateRangeClient(ctrl *gomock.Controller) *MockAuditLog_GetByDateRangeClient {
	mock := &MockAuditLog_GetByDateRangeClient{ctrl: ctrl}
	mock.recorder = &MockAuditLog_GetByDateRangeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditLog_GetByDateRangeClient) EXPECT() *MockAuditLog_GetByDateRangeClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockAuditLog_GetByDateRangeClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockAuditLog_GetByDateRangeClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockAuditLog_GetByDateRangeClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockAuditLog_GetByDateRangeClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAuditLog_GetByDateRangeClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAuditLog_GetByDateRangeClient)(nil).Context))
}

// Header mocks base method.
func (m *MockAuditLog_GetByDateRangeClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockAuditLog_GetByDateRangeClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockAuditLog_GetByDateRangeClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockAuditLog_GetByDateRangeClient) Recv() (*audit.HumanReadableEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*audit.HumanReadableEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockAuditLog_GetByDateRangeClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockAuditLog_GetByDateRangeClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockAuditLog_GetByDateRangeClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockAuditLog_GetByDateRangeClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAuditLog_GetByDateRangeClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockAuditLog_GetByDateRangeClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockAuditLog_GetByDateRangeClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAuditLog_GetByDateRangeClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockAuditLog_GetByDateRangeClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockAuditLog_GetByDateRangeClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockAuditLog_GetByDateRangeClient)(nil).Trailer))
}

// MockAuditLog_GetUserActionsClient is a mock of AuditLog_GetUserActionsClient interface.
type MockAuditLog_GetUserActionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuditLog_GetUserActionsClientMockRecorder
}

// MockAuditLog_GetUserActionsClientMockRecorder is the mock recorder for MockAuditLog_GetUserActionsClient.
type MockAuditLog_GetUserActionsClientMockRecorder struct {
	mock *MockAuditLog_GetUserActionsClient
}

// NewMockAuditLog_GetUserActionsClient creates a new mock instance.
func NewMockAuditLog_GetUserActionsClient(ctrl *gomock.Controller) *MockAuditLog_GetUserActionsClient {
	mock := &MockAuditLog_GetUserActionsClient{ctrl: ctrl}
	mock.recorder = &MockAuditLog_GetUserActionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditLog_GetUserActionsClient) EXPECT() *MockAuditLog_GetUserActionsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockAuditLog_GetUserActionsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockAuditLog_GetUserActionsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockAuditLog_GetUserActionsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockAuditLog_GetUserActionsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAuditLog_GetUserActionsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAuditLog_GetUserActionsClient)(nil).Context))
}

// Header mocks base method.
func (m *MockAuditLog_GetUserActionsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockAuditLog_GetUserActionsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockAuditLog_GetUserActionsClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockAuditLog_GetUserActionsClient) Recv() (*audit.HumanReadableEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*audit.HumanReadableEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockAuditLog_GetUserActionsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockAuditLog_GetUserActionsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockAuditLog_GetUserActionsClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockAuditLog_GetUserActionsClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockAuditLog_GetUserActionsClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockAuditLog_GetUserActionsClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockAuditLog_GetUserActionsClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockAuditLog_GetUserActionsClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockAuditLog_GetUserActionsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockAuditLog_GetUserActionsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockAuditLog_GetUserActionsClient)(nil).Trailer))
}