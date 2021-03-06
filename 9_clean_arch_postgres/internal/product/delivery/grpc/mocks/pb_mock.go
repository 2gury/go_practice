// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/product/delivery/grpc/product_grpc.pb.go

// Package mock_grpc is a generated GoMock package.
package mock_grpc

import (
	context "context"
	grpc "go_practice/9_clean_arch_db/internal/product/delivery/grpc"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc0 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockProductServiceClient is a mock of ProductServiceClient interface.
type MockProductServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceClientMockRecorder
}

// MockProductServiceClientMockRecorder is the mock recorder for MockProductServiceClient.
type MockProductServiceClientMockRecorder struct {
	mock *MockProductServiceClient
}

// NewMockProductServiceClient creates a new mock instance.
func NewMockProductServiceClient(ctrl *gomock.Controller) *MockProductServiceClient {
	mock := &MockProductServiceClient{ctrl: ctrl}
	mock.recorder = &MockProductServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductServiceClient) EXPECT() *MockProductServiceClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductServiceClient) Create(ctx context.Context, in *grpc.Product, opts ...grpc0.CallOption) (*grpc.ProductIdValue, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*grpc.ProductIdValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductServiceClientMockRecorder) Create(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductServiceClient)(nil).Create), varargs...)
}

// DeleteById mocks base method.
func (m *MockProductServiceClient) DeleteById(ctx context.Context, in *grpc.ProductIdValue, opts ...grpc0.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteById", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockProductServiceClientMockRecorder) DeleteById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockProductServiceClient)(nil).DeleteById), varargs...)
}

// GetById mocks base method.
func (m *MockProductServiceClient) GetById(ctx context.Context, in *grpc.ProductIdValue, opts ...grpc0.CallOption) (*grpc.Product, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetById", varargs...)
	ret0, _ := ret[0].(*grpc.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockProductServiceClientMockRecorder) GetById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockProductServiceClient)(nil).GetById), varargs...)
}

// List mocks base method.
func (m *MockProductServiceClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc0.CallOption) (*grpc.ArrayProducts, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*grpc.ArrayProducts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProductServiceClientMockRecorder) List(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProductServiceClient)(nil).List), varargs...)
}

// UpdateById mocks base method.
func (m *MockProductServiceClient) UpdateById(ctx context.Context, in *grpc.UpdateInfoProduct, opts ...grpc0.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateById", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateById indicates an expected call of UpdateById.
func (mr *MockProductServiceClientMockRecorder) UpdateById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateById", reflect.TypeOf((*MockProductServiceClient)(nil).UpdateById), varargs...)
}

// MockProductServiceServer is a mock of ProductServiceServer interface.
type MockProductServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceServerMockRecorder
}

// MockProductServiceServerMockRecorder is the mock recorder for MockProductServiceServer.
type MockProductServiceServerMockRecorder struct {
	mock *MockProductServiceServer
}

// NewMockProductServiceServer creates a new mock instance.
func NewMockProductServiceServer(ctrl *gomock.Controller) *MockProductServiceServer {
	mock := &MockProductServiceServer{ctrl: ctrl}
	mock.recorder = &MockProductServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductServiceServer) EXPECT() *MockProductServiceServerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductServiceServer) Create(arg0 context.Context, arg1 *grpc.Product) (*grpc.ProductIdValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*grpc.ProductIdValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductServiceServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductServiceServer)(nil).Create), arg0, arg1)
}

// DeleteById mocks base method.
func (m *MockProductServiceServer) DeleteById(arg0 context.Context, arg1 *grpc.ProductIdValue) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockProductServiceServerMockRecorder) DeleteById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockProductServiceServer)(nil).DeleteById), arg0, arg1)
}

// GetById mocks base method.
func (m *MockProductServiceServer) GetById(arg0 context.Context, arg1 *grpc.ProductIdValue) (*grpc.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*grpc.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockProductServiceServerMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockProductServiceServer)(nil).GetById), arg0, arg1)
}

// List mocks base method.
func (m *MockProductServiceServer) List(arg0 context.Context, arg1 *emptypb.Empty) (*grpc.ArrayProducts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*grpc.ArrayProducts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProductServiceServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProductServiceServer)(nil).List), arg0, arg1)
}

// UpdateById mocks base method.
func (m *MockProductServiceServer) UpdateById(arg0 context.Context, arg1 *grpc.UpdateInfoProduct) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateById", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateById indicates an expected call of UpdateById.
func (mr *MockProductServiceServerMockRecorder) UpdateById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateById", reflect.TypeOf((*MockProductServiceServer)(nil).UpdateById), arg0, arg1)
}

// mustEmbedUnimplementedProductServiceServer mocks base method.
func (m *MockProductServiceServer) mustEmbedUnimplementedProductServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedProductServiceServer")
}

// mustEmbedUnimplementedProductServiceServer indicates an expected call of mustEmbedUnimplementedProductServiceServer.
func (mr *MockProductServiceServerMockRecorder) mustEmbedUnimplementedProductServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedProductServiceServer", reflect.TypeOf((*MockProductServiceServer)(nil).mustEmbedUnimplementedProductServiceServer))
}

// MockUnsafeProductServiceServer is a mock of UnsafeProductServiceServer interface.
type MockUnsafeProductServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeProductServiceServerMockRecorder
}

// MockUnsafeProductServiceServerMockRecorder is the mock recorder for MockUnsafeProductServiceServer.
type MockUnsafeProductServiceServerMockRecorder struct {
	mock *MockUnsafeProductServiceServer
}

// NewMockUnsafeProductServiceServer creates a new mock instance.
func NewMockUnsafeProductServiceServer(ctrl *gomock.Controller) *MockUnsafeProductServiceServer {
	mock := &MockUnsafeProductServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeProductServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeProductServiceServer) EXPECT() *MockUnsafeProductServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedProductServiceServer mocks base method.
func (m *MockUnsafeProductServiceServer) mustEmbedUnimplementedProductServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedProductServiceServer")
}

// mustEmbedUnimplementedProductServiceServer indicates an expected call of mustEmbedUnimplementedProductServiceServer.
func (mr *MockUnsafeProductServiceServerMockRecorder) mustEmbedUnimplementedProductServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedProductServiceServer", reflect.TypeOf((*MockUnsafeProductServiceServer)(nil).mustEmbedUnimplementedProductServiceServer))
}
