// Code generated by MockGen. DO NOT EDIT.
// Source: category.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	entity "quickeat/pkg/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCategoryService is a mock of CategoryService interface.
type MockCategoryService struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryServiceMockRecorder
}

// MockCategoryServiceMockRecorder is the mock recorder for MockCategoryService.
type MockCategoryServiceMockRecorder struct {
	mock *MockCategoryService
}

// NewMockCategoryService creates a new mock instance.
func NewMockCategoryService(ctrl *gomock.Controller) *MockCategoryService {
	mock := &MockCategoryService{ctrl: ctrl}
	mock.recorder = &MockCategoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryService) EXPECT() *MockCategoryServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCategoryService) Create(ctx context.Context, category *entity.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, category)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCategoryServiceMockRecorder) Create(ctx, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCategoryService)(nil).Create), ctx, category)
}

// DeleteCategoryById mocks base method.
func (m *MockCategoryService) DeleteCategoryById(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategoryById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategoryById indicates an expected call of DeleteCategoryById.
func (mr *MockCategoryServiceMockRecorder) DeleteCategoryById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategoryById", reflect.TypeOf((*MockCategoryService)(nil).DeleteCategoryById), ctx, id)
}

// DeleteCategoryByName mocks base method.
func (m *MockCategoryService) DeleteCategoryByName(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategoryByName", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategoryByName indicates an expected call of DeleteCategoryByName.
func (mr *MockCategoryServiceMockRecorder) DeleteCategoryByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategoryByName", reflect.TypeOf((*MockCategoryService)(nil).DeleteCategoryByName), ctx, name)
}

// Get mocks base method.
func (m *MockCategoryService) Get(ctx context.Context, id *int) ([]*entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].([]*entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCategoryServiceMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCategoryService)(nil).Get), ctx, id)
}

// GetByDish mocks base method.
func (m *MockCategoryService) GetByDish(ctx context.Context, dishId int) (*entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDish", ctx, dishId)
	ret0, _ := ret[0].(*entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDish indicates an expected call of GetByDish.
func (mr *MockCategoryServiceMockRecorder) GetByDish(ctx, dishId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDish", reflect.TypeOf((*MockCategoryService)(nil).GetByDish), ctx, dishId)
}

// Update mocks base method.
func (m *MockCategoryService) Update(ctx context.Context, category *entity.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, category)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCategoryServiceMockRecorder) Update(ctx, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCategoryService)(nil).Update), ctx, category)
}
