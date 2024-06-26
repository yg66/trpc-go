// Code generated by MockGen. DO NOT EDIT.
// Source: helloworld.trpc.go

// Package stream is a generated GoMock package.
package stream

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "trpc.group/trpc-go/trpc-go/client"
)

// MockTestStreamService is a mock of TestStreamService interface.
type MockTestStreamService struct {
	ctrl     *gomock.Controller
	recorder *MockTestStreamServiceMockRecorder
}

// MockTestStreamServiceMockRecorder is the mock recorder for MockTestStreamService.
type MockTestStreamServiceMockRecorder struct {
	mock *MockTestStreamService
}

// NewMockTestStreamService creates a new mock instance.
func NewMockTestStreamService(ctrl *gomock.Controller) *MockTestStreamService {
	mock := &MockTestStreamService{ctrl: ctrl}
	mock.recorder = &MockTestStreamServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestStreamService) EXPECT() *MockTestStreamServiceMockRecorder {
	return m.recorder
}

// UploadFileStream mocks base method.
func (m *MockTestStreamService) UploadFileStream(arg0 TestStream_UploadFileStreamServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFileStream", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFileStream indicates an expected call of UploadFileStream.
func (mr *MockTestStreamServiceMockRecorder) UploadFileStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFileStream", reflect.TypeOf((*MockTestStreamService)(nil).UploadFileStream), arg0)
}

// MockTestStream_UploadFileStreamServer is a mock of TestStream_UploadFileStreamServer interface.
type MockTestStream_UploadFileStreamServer struct {
	ctrl     *gomock.Controller
	recorder *MockTestStream_UploadFileStreamServerMockRecorder
}

// MockTestStream_UploadFileStreamServerMockRecorder is the mock recorder for MockTestStream_UploadFileStreamServer.
type MockTestStream_UploadFileStreamServerMockRecorder struct {
	mock *MockTestStream_UploadFileStreamServer
}

// NewMockTestStream_UploadFileStreamServer creates a new mock instance.
func NewMockTestStream_UploadFileStreamServer(ctrl *gomock.Controller) *MockTestStream_UploadFileStreamServer {
	mock := &MockTestStream_UploadFileStreamServer{ctrl: ctrl}
	mock.recorder = &MockTestStream_UploadFileStreamServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestStream_UploadFileStreamServer) EXPECT() *MockTestStream_UploadFileStreamServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockTestStream_UploadFileStreamServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockTestStream_UploadFileStreamServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTestStream_UploadFileStreamServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockTestStream_UploadFileStreamServer) Recv() (*UploadFileReq, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*UploadFileReq)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockTestStream_UploadFileStreamServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockTestStream_UploadFileStreamServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockTestStream_UploadFileStreamServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockTestStream_UploadFileStreamServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTestStream_UploadFileStreamServer)(nil).RecvMsg), m)
}

// SendAndClose mocks base method.
func (m *MockTestStream_UploadFileStreamServer) SendAndClose(arg0 *UploadFileResp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAndClose", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAndClose indicates an expected call of SendAndClose.
func (mr *MockTestStream_UploadFileStreamServerMockRecorder) SendAndClose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAndClose", reflect.TypeOf((*MockTestStream_UploadFileStreamServer)(nil).SendAndClose), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockTestStream_UploadFileStreamServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockTestStream_UploadFileStreamServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTestStream_UploadFileStreamServer)(nil).SendMsg), m)
}

// MockTestStreamClientProxy is a mock of TestStreamClientProxy interface.
type MockTestStreamClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockTestStreamClientProxyMockRecorder
}

// MockTestStreamClientProxyMockRecorder is the mock recorder for MockTestStreamClientProxy.
type MockTestStreamClientProxyMockRecorder struct {
	mock *MockTestStreamClientProxy
}

// NewMockTestStreamClientProxy creates a new mock instance.
func NewMockTestStreamClientProxy(ctrl *gomock.Controller) *MockTestStreamClientProxy {
	mock := &MockTestStreamClientProxy{ctrl: ctrl}
	mock.recorder = &MockTestStreamClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestStreamClientProxy) EXPECT() *MockTestStreamClientProxyMockRecorder {
	return m.recorder
}

// UploadFileStream mocks base method.
func (m *MockTestStreamClientProxy) UploadFileStream(ctx context.Context, opts ...client.Option) (TestStream_UploadFileStreamClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadFileStream", varargs...)
	ret0, _ := ret[0].(TestStream_UploadFileStreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFileStream indicates an expected call of UploadFileStream.
func (mr *MockTestStreamClientProxyMockRecorder) UploadFileStream(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFileStream", reflect.TypeOf((*MockTestStreamClientProxy)(nil).UploadFileStream), varargs...)
}

// MockTestStream_UploadFileStreamClient is a mock of TestStream_UploadFileStreamClient interface.
type MockTestStream_UploadFileStreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockTestStream_UploadFileStreamClientMockRecorder
}

// MockTestStream_UploadFileStreamClientMockRecorder is the mock recorder for MockTestStream_UploadFileStreamClient.
type MockTestStream_UploadFileStreamClientMockRecorder struct {
	mock *MockTestStream_UploadFileStreamClient
}

// NewMockTestStream_UploadFileStreamClient creates a new mock instance.
func NewMockTestStream_UploadFileStreamClient(ctrl *gomock.Controller) *MockTestStream_UploadFileStreamClient {
	mock := &MockTestStream_UploadFileStreamClient{ctrl: ctrl}
	mock.recorder = &MockTestStream_UploadFileStreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestStream_UploadFileStreamClient) EXPECT() *MockTestStream_UploadFileStreamClientMockRecorder {
	return m.recorder
}

// CloseAndRecv mocks base method.
func (m *MockTestStream_UploadFileStreamClient) CloseAndRecv() (*UploadFileResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndRecv")
	ret0, _ := ret[0].(*UploadFileResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAndRecv indicates an expected call of CloseAndRecv.
func (mr *MockTestStream_UploadFileStreamClientMockRecorder) CloseAndRecv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndRecv", reflect.TypeOf((*MockTestStream_UploadFileStreamClient)(nil).CloseAndRecv))
}

// CloseSend mocks base method.
func (m *MockTestStream_UploadFileStreamClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockTestStream_UploadFileStreamClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockTestStream_UploadFileStreamClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockTestStream_UploadFileStreamClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockTestStream_UploadFileStreamClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTestStream_UploadFileStreamClient)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockTestStream_UploadFileStreamClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockTestStream_UploadFileStreamClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTestStream_UploadFileStreamClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockTestStream_UploadFileStreamClient) Send(arg0 *UploadFileReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockTestStream_UploadFileStreamClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTestStream_UploadFileStreamClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockTestStream_UploadFileStreamClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockTestStream_UploadFileStreamClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTestStream_UploadFileStreamClient)(nil).SendMsg), m)
}
