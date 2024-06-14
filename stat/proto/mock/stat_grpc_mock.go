// Code generated by MockGen. DO NOT EDIT.
// Source: proto/stat_grpc.pb.go

// Package mock_statpb is a generated GoMock package.
package mock_statpb

import (
	context "context"
	reflect "reflect"
	proto "stat/proto"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockStatServiceClient is a mock of StatServiceClient interface.
type MockStatServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockStatServiceClientMockRecorder
}

// MockStatServiceClientMockRecorder is the mock recorder for MockStatServiceClient.
type MockStatServiceClientMockRecorder struct {
	mock *MockStatServiceClient
}

// NewMockStatServiceClient creates a new mock instance.
func NewMockStatServiceClient(ctrl *gomock.Controller) *MockStatServiceClient {
	mock := &MockStatServiceClient{ctrl: ctrl}
	mock.recorder = &MockStatServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatServiceClient) EXPECT() *MockStatServiceClientMockRecorder {
	return m.recorder
}

// GetPostStats mocks base method.
func (m *MockStatServiceClient) GetPostStats(ctx context.Context, in *proto.GetPostStatsReq, opts ...grpc.CallOption) (*proto.GetPostStatsResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPostStats", varargs...)
	ret0, _ := ret[0].(*proto.GetPostStatsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostStats indicates an expected call of GetPostStats.
func (mr *MockStatServiceClientMockRecorder) GetPostStats(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostStats", reflect.TypeOf((*MockStatServiceClient)(nil).GetPostStats), varargs...)
}

// GetTopPosts mocks base method.
func (m *MockStatServiceClient) GetTopPosts(ctx context.Context, in *proto.GetTopPostsReq, opts ...grpc.CallOption) (*proto.GetTopPostsResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTopPosts", varargs...)
	ret0, _ := ret[0].(*proto.GetTopPostsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopPosts indicates an expected call of GetTopPosts.
func (mr *MockStatServiceClientMockRecorder) GetTopPosts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopPosts", reflect.TypeOf((*MockStatServiceClient)(nil).GetTopPosts), varargs...)
}

// GetTopUsers mocks base method.
func (m *MockStatServiceClient) GetTopUsers(ctx context.Context, in *proto.GetTopUsersReq, opts ...grpc.CallOption) (*proto.GetTopUsersResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTopUsers", varargs...)
	ret0, _ := ret[0].(*proto.GetTopUsersResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopUsers indicates an expected call of GetTopUsers.
func (mr *MockStatServiceClientMockRecorder) GetTopUsers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopUsers", reflect.TypeOf((*MockStatServiceClient)(nil).GetTopUsers), varargs...)
}

// MockStatServiceServer is a mock of StatServiceServer interface.
type MockStatServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockStatServiceServerMockRecorder
}

// MockStatServiceServerMockRecorder is the mock recorder for MockStatServiceServer.
type MockStatServiceServerMockRecorder struct {
	mock *MockStatServiceServer
}

// NewMockStatServiceServer creates a new mock instance.
func NewMockStatServiceServer(ctrl *gomock.Controller) *MockStatServiceServer {
	mock := &MockStatServiceServer{ctrl: ctrl}
	mock.recorder = &MockStatServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatServiceServer) EXPECT() *MockStatServiceServerMockRecorder {
	return m.recorder
}

// GetPostStats mocks base method.
func (m *MockStatServiceServer) GetPostStats(arg0 context.Context, arg1 *proto.GetPostStatsReq) (*proto.GetPostStatsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostStats", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetPostStatsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostStats indicates an expected call of GetPostStats.
func (mr *MockStatServiceServerMockRecorder) GetPostStats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostStats", reflect.TypeOf((*MockStatServiceServer)(nil).GetPostStats), arg0, arg1)
}

// GetTopPosts mocks base method.
func (m *MockStatServiceServer) GetTopPosts(arg0 context.Context, arg1 *proto.GetTopPostsReq) (*proto.GetTopPostsResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopPosts", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetTopPostsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopPosts indicates an expected call of GetTopPosts.
func (mr *MockStatServiceServerMockRecorder) GetTopPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopPosts", reflect.TypeOf((*MockStatServiceServer)(nil).GetTopPosts), arg0, arg1)
}

// GetTopUsers mocks base method.
func (m *MockStatServiceServer) GetTopUsers(arg0 context.Context, arg1 *proto.GetTopUsersReq) (*proto.GetTopUsersResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopUsers", arg0, arg1)
	ret0, _ := ret[0].(*proto.GetTopUsersResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopUsers indicates an expected call of GetTopUsers.
func (mr *MockStatServiceServerMockRecorder) GetTopUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopUsers", reflect.TypeOf((*MockStatServiceServer)(nil).GetTopUsers), arg0, arg1)
}

// mustEmbedUnimplementedStatServiceServer mocks base method.
func (m *MockStatServiceServer) mustEmbedUnimplementedStatServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedStatServiceServer")
}

// mustEmbedUnimplementedStatServiceServer indicates an expected call of mustEmbedUnimplementedStatServiceServer.
func (mr *MockStatServiceServerMockRecorder) mustEmbedUnimplementedStatServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedStatServiceServer", reflect.TypeOf((*MockStatServiceServer)(nil).mustEmbedUnimplementedStatServiceServer))
}

// MockUnsafeStatServiceServer is a mock of UnsafeStatServiceServer interface.
type MockUnsafeStatServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeStatServiceServerMockRecorder
}

// MockUnsafeStatServiceServerMockRecorder is the mock recorder for MockUnsafeStatServiceServer.
type MockUnsafeStatServiceServerMockRecorder struct {
	mock *MockUnsafeStatServiceServer
}

// NewMockUnsafeStatServiceServer creates a new mock instance.
func NewMockUnsafeStatServiceServer(ctrl *gomock.Controller) *MockUnsafeStatServiceServer {
	mock := &MockUnsafeStatServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeStatServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeStatServiceServer) EXPECT() *MockUnsafeStatServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedStatServiceServer mocks base method.
func (m *MockUnsafeStatServiceServer) mustEmbedUnimplementedStatServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedStatServiceServer")
}

// mustEmbedUnimplementedStatServiceServer indicates an expected call of mustEmbedUnimplementedStatServiceServer.
func (mr *MockUnsafeStatServiceServerMockRecorder) mustEmbedUnimplementedStatServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedStatServiceServer", reflect.TypeOf((*MockUnsafeStatServiceServer)(nil).mustEmbedUnimplementedStatServiceServer))
}