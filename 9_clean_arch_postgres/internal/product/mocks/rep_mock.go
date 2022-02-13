// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/product/repository.go

// Package mock_product is a generated GoMock package.
package mock_product

import (
	models "go_practice/9_clean_arch_db/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// DeleteById mocks base method.
func (m *MockProductRepository) DeleteById(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteById", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteById indicates an expected call of DeleteById.
func (mr *MockProductRepositoryMockRecorder) DeleteById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteById", reflect.TypeOf((*MockProductRepository)(nil).DeleteById), id)
}

// Insert mocks base method.
func (m *MockProductRepository) Insert(product *models.Product) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", product)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockProductRepositoryMockRecorder) Insert(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockProductRepository)(nil).Insert), product)
}

// SelectAll mocks base method.
func (m *MockProductRepository) SelectAll() ([]*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAll")
	ret0, _ := ret[0].([]*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAll indicates an expected call of SelectAll.
func (mr *MockProductRepositoryMockRecorder) SelectAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAll", reflect.TypeOf((*MockProductRepository)(nil).SelectAll))
}

// SelectById mocks base method.
func (m *MockProductRepository) SelectById(id uint64) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectById", id)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectById indicates an expected call of SelectById.
func (mr *MockProductRepositoryMockRecorder) SelectById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectById", reflect.TypeOf((*MockProductRepository)(nil).SelectById), id)
}

// UpdateById mocks base method.
func (m *MockProductRepository) UpdateById(productId uint64, updatedProduct *models.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateById", productId, updatedProduct)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateById indicates an expected call of UpdateById.
func (mr *MockProductRepositoryMockRecorder) UpdateById(productId, updatedProduct interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateById", reflect.TypeOf((*MockProductRepository)(nil).UpdateById), productId, updatedProduct)
}
