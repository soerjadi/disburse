// Code generated by MockGen. DO NOT EDIT.
// Source: type.go
//
// Generated by this command:
//
//	mockgen -package=mocks -mock_names=Repository=MockTransactionRepository -destination=../../mocks/transaction_repo_mock.go -source=type.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/soerjadi/brick/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockTransactionRepository is a mock of Repository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// Callback mocks base method.
func (m *MockTransactionRepository) Callback(ctx context.Context, req model.CallbackRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Callback", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Callback indicates an expected call of Callback.
func (mr *MockTransactionRepositoryMockRecorder) Callback(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Callback", reflect.TypeOf((*MockTransactionRepository)(nil).Callback), ctx, req)
}

// CheckAccount mocks base method.
func (m *MockTransactionRepository) CheckAccount(ctx context.Context, req model.CheckAccountRequest) (*model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAccount", ctx, req)
	ret0, _ := ret[0].(*model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAccount indicates an expected call of CheckAccount.
func (mr *MockTransactionRepositoryMockRecorder) CheckAccount(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccount", reflect.TypeOf((*MockTransactionRepository)(nil).CheckAccount), ctx, req)
}

// Disbursement mocks base method.
func (m *MockTransactionRepository) Disbursement(ctx context.Context, req model.DisbursementRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disbursement", ctx, req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Disbursement indicates an expected call of Disbursement.
func (mr *MockTransactionRepositoryMockRecorder) Disbursement(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disbursement", reflect.TypeOf((*MockTransactionRepository)(nil).Disbursement), ctx, req)
}

// GetAccountByNumberOrigin mocks base method.
func (m *MockTransactionRepository) GetAccountByNumberOrigin(ctx context.Context, Number, OriginBank string) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByNumberOrigin", ctx, Number, OriginBank)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByNumberOrigin indicates an expected call of GetAccountByNumberOrigin.
func (mr *MockTransactionRepositoryMockRecorder) GetAccountByNumberOrigin(ctx, Number, OriginBank any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByNumberOrigin", reflect.TypeOf((*MockTransactionRepository)(nil).GetAccountByNumberOrigin), ctx, Number, OriginBank)
}

// InsertAccount mocks base method.
func (m *MockTransactionRepository) InsertAccount(ctx context.Context, req model.Account) (model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertAccount", ctx, req)
	ret0, _ := ret[0].(model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertAccount indicates an expected call of InsertAccount.
func (mr *MockTransactionRepositoryMockRecorder) InsertAccount(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertAccount", reflect.TypeOf((*MockTransactionRepository)(nil).InsertAccount), ctx, req)
}

// InsertTrx mocks base method.
func (m *MockTransactionRepository) InsertTrx(ctx context.Context, req model.Transaction) (model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTrx", ctx, req)
	ret0, _ := ret[0].(model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertTrx indicates an expected call of InsertTrx.
func (mr *MockTransactionRepositoryMockRecorder) InsertTrx(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTrx", reflect.TypeOf((*MockTransactionRepository)(nil).InsertTrx), ctx, req)
}

// UpdateTrxID mocks base method.
func (m *MockTransactionRepository) UpdateTrxID(ctx context.Context, id int64, trxID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrxID", ctx, id, trxID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTrxID indicates an expected call of UpdateTrxID.
func (mr *MockTransactionRepositoryMockRecorder) UpdateTrxID(ctx, id, trxID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrxID", reflect.TypeOf((*MockTransactionRepository)(nil).UpdateTrxID), ctx, id, trxID)
}