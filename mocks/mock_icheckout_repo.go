// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/icheckout_repo.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	context "context"
	models "discount-service/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICheckoutRepo is a mock of ICheckoutRepo interface.
type MockICheckoutRepo struct {
	ctrl     *gomock.Controller
	recorder *MockICheckoutRepoMockRecorder
}

// MockICheckoutRepoMockRecorder is the mock recorder for MockICheckoutRepo.
type MockICheckoutRepoMockRecorder struct {
	mock *MockICheckoutRepo
}

// NewMockICheckoutRepo creates a new mock instance.
func NewMockICheckoutRepo(ctrl *gomock.Controller) *MockICheckoutRepo {
	mock := &MockICheckoutRepo{ctrl: ctrl}
	mock.recorder = &MockICheckoutRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICheckoutRepo) EXPECT() *MockICheckoutRepoMockRecorder {
	return m.recorder
}

// BeginTrx mocks base method.
func (m *MockICheckoutRepo) BeginTrx(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTrx", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeginTrx indicates an expected call of BeginTrx.
func (mr *MockICheckoutRepoMockRecorder) BeginTrx(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTrx", reflect.TypeOf((*MockICheckoutRepo)(nil).BeginTrx), ctx)
}

// CommitTrx mocks base method.
func (m *MockICheckoutRepo) CommitTrx(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitTrx", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitTrx indicates an expected call of CommitTrx.
func (mr *MockICheckoutRepoMockRecorder) CommitTrx(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitTrx", reflect.TypeOf((*MockICheckoutRepo)(nil).CommitTrx), ctx)
}

// RepoGetItemsList mocks base method.
func (m *MockICheckoutRepo) RepoGetItemsList(ctx context.Context, sku []string) ([]models.Items, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RepoGetItemsList", ctx, sku)
	ret0, _ := ret[0].([]models.Items)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RepoGetItemsList indicates an expected call of RepoGetItemsList.
func (mr *MockICheckoutRepoMockRecorder) RepoGetItemsList(ctx, sku interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RepoGetItemsList", reflect.TypeOf((*MockICheckoutRepo)(nil).RepoGetItemsList), ctx, sku)
}

// RepoInsertOrder mocks base method.
func (m *MockICheckoutRepo) RepoInsertOrder(ctx context.Context, data []models.Orders) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RepoInsertOrder", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// RepoInsertOrder indicates an expected call of RepoInsertOrder.
func (mr *MockICheckoutRepoMockRecorder) RepoInsertOrder(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RepoInsertOrder", reflect.TypeOf((*MockICheckoutRepo)(nil).RepoInsertOrder), ctx, data)
}

// RepoPromotionActiveList mocks base method.
func (m *MockICheckoutRepo) RepoPromotionActiveList(ctx context.Context) ([]models.PromotionItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RepoPromotionActiveList", ctx)
	ret0, _ := ret[0].([]models.PromotionItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RepoPromotionActiveList indicates an expected call of RepoPromotionActiveList.
func (mr *MockICheckoutRepoMockRecorder) RepoPromotionActiveList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RepoPromotionActiveList", reflect.TypeOf((*MockICheckoutRepo)(nil).RepoPromotionActiveList), ctx)
}

// RepoUpdateItems mocks base method.
func (m *MockICheckoutRepo) RepoUpdateItems(ctx context.Context, data models.Orders) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RepoUpdateItems", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// RepoUpdateItems indicates an expected call of RepoUpdateItems.
func (mr *MockICheckoutRepoMockRecorder) RepoUpdateItems(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RepoUpdateItems", reflect.TypeOf((*MockICheckoutRepo)(nil).RepoUpdateItems), ctx, data)
}

// RoolBackTrx mocks base method.
func (m *MockICheckoutRepo) RoolBackTrx(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RoolBackTrx", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RoolBackTrx indicates an expected call of RoolBackTrx.
func (mr *MockICheckoutRepoMockRecorder) RoolBackTrx(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RoolBackTrx", reflect.TypeOf((*MockICheckoutRepo)(nil).RoolBackTrx), ctx)
}
