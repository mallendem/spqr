// Code generated by MockGen. DO NOT EDIT.
// Source: ./router/poolmgr/pool_mgr.go
//
// Generated by this command:
//
//	mockgen -source=./router/poolmgr/pool_mgr.go -destination=./router/mock/poolmgr/mock_pool_mgr.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	kr "github.com/pg-sharding/spqr/pkg/models/kr"
	txstatus "github.com/pg-sharding/spqr/pkg/txstatus"
	client "github.com/pg-sharding/spqr/router/client"
	poolmgr "github.com/pg-sharding/spqr/router/poolmgr"
	gomock "go.uber.org/mock/gomock"
)

// MockConnectionKeeper is a mock of ConnectionKeeper interface.
type MockConnectionKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionKeeperMockRecorder
	isgomock struct{}
}

// MockConnectionKeeperMockRecorder is the mock recorder for MockConnectionKeeper.
type MockConnectionKeeperMockRecorder struct {
	mock *MockConnectionKeeper
}

// NewMockConnectionKeeper creates a new mock instance.
func NewMockConnectionKeeper(ctrl *gomock.Controller) *MockConnectionKeeper {
	mock := &MockConnectionKeeper{ctrl: ctrl}
	mock.recorder = &MockConnectionKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionKeeper) EXPECT() *MockConnectionKeeperMockRecorder {
	return m.recorder
}

// ActiveShards mocks base method.
func (m *MockConnectionKeeper) ActiveShards() []kr.ShardKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveShards")
	ret0, _ := ret[0].([]kr.ShardKey)
	return ret0
}

// ActiveShards indicates an expected call of ActiveShards.
func (mr *MockConnectionKeeperMockRecorder) ActiveShards() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveShards", reflect.TypeOf((*MockConnectionKeeper)(nil).ActiveShards))
}

// ActiveShardsReset mocks base method.
func (m *MockConnectionKeeper) ActiveShardsReset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ActiveShardsReset")
}

// ActiveShardsReset indicates an expected call of ActiveShardsReset.
func (mr *MockConnectionKeeperMockRecorder) ActiveShardsReset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveShardsReset", reflect.TypeOf((*MockConnectionKeeper)(nil).ActiveShardsReset))
}

// Client mocks base method.
func (m *MockConnectionKeeper) Client() client.RouterClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(client.RouterClient)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockConnectionKeeperMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockConnectionKeeper)(nil).Client))
}

// DataPending mocks base method.
func (m *MockConnectionKeeper) DataPending() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataPending")
	ret0, _ := ret[0].(bool)
	return ret0
}

// DataPending indicates an expected call of DataPending.
func (mr *MockConnectionKeeperMockRecorder) DataPending() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataPending", reflect.TypeOf((*MockConnectionKeeper)(nil).DataPending))
}

// SetTxStatus mocks base method.
func (m *MockConnectionKeeper) SetTxStatus(status txstatus.TXStatus) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTxStatus", status)
}

// SetTxStatus indicates an expected call of SetTxStatus.
func (mr *MockConnectionKeeperMockRecorder) SetTxStatus(status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTxStatus", reflect.TypeOf((*MockConnectionKeeper)(nil).SetTxStatus), status)
}

// SyncCount mocks base method.
func (m *MockConnectionKeeper) SyncCount() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncCount")
	ret0, _ := ret[0].(int64)
	return ret0
}

// SyncCount indicates an expected call of SyncCount.
func (mr *MockConnectionKeeperMockRecorder) SyncCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncCount", reflect.TypeOf((*MockConnectionKeeper)(nil).SyncCount))
}

// TxStatus mocks base method.
func (m *MockConnectionKeeper) TxStatus() txstatus.TXStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxStatus")
	ret0, _ := ret[0].(txstatus.TXStatus)
	return ret0
}

// TxStatus indicates an expected call of TxStatus.
func (mr *MockConnectionKeeperMockRecorder) TxStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxStatus", reflect.TypeOf((*MockConnectionKeeper)(nil).TxStatus))
}

// MockPoolMgr is a mock of PoolMgr interface.
type MockPoolMgr struct {
	ctrl     *gomock.Controller
	recorder *MockPoolMgrMockRecorder
	isgomock struct{}
}

// MockPoolMgrMockRecorder is the mock recorder for MockPoolMgr.
type MockPoolMgrMockRecorder struct {
	mock *MockPoolMgr
}

// NewMockPoolMgr creates a new mock instance.
func NewMockPoolMgr(ctrl *gomock.Controller) *MockPoolMgr {
	mock := &MockPoolMgr{ctrl: ctrl}
	mock.recorder = &MockPoolMgrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPoolMgr) EXPECT() *MockPoolMgrMockRecorder {
	return m.recorder
}

// ConnectionActive mocks base method.
func (m *MockPoolMgr) ConnectionActive(rst poolmgr.ConnectionKeeper) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectionActive", rst)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ConnectionActive indicates an expected call of ConnectionActive.
func (mr *MockPoolMgrMockRecorder) ConnectionActive(rst any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionActive", reflect.TypeOf((*MockPoolMgr)(nil).ConnectionActive), rst)
}

// TXEndCB mocks base method.
func (m *MockPoolMgr) TXEndCB(rst poolmgr.ConnectionKeeper) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TXEndCB", rst)
	ret0, _ := ret[0].(error)
	return ret0
}

// TXEndCB indicates an expected call of TXEndCB.
func (mr *MockPoolMgrMockRecorder) TXEndCB(rst any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TXEndCB", reflect.TypeOf((*MockPoolMgr)(nil).TXEndCB), rst)
}

// UnRouteCB mocks base method.
func (m *MockPoolMgr) UnRouteCB(client client.RouterClient, sh []kr.ShardKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnRouteCB", client, sh)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnRouteCB indicates an expected call of UnRouteCB.
func (mr *MockPoolMgrMockRecorder) UnRouteCB(client, sh any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnRouteCB", reflect.TypeOf((*MockPoolMgr)(nil).UnRouteCB), client, sh)
}

// UnRouteWithError mocks base method.
func (m *MockPoolMgr) UnRouteWithError(client client.RouterClient, sh []kr.ShardKey, errmsg error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnRouteWithError", client, sh, errmsg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnRouteWithError indicates an expected call of UnRouteWithError.
func (mr *MockPoolMgrMockRecorder) UnRouteWithError(client, sh, errmsg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnRouteWithError", reflect.TypeOf((*MockPoolMgr)(nil).UnRouteWithError), client, sh, errmsg)
}

// ValidateSliceChange mocks base method.
func (m *MockPoolMgr) ValidateSliceChange(rst poolmgr.ConnectionKeeper) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateSliceChange", rst)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidateSliceChange indicates an expected call of ValidateSliceChange.
func (mr *MockPoolMgrMockRecorder) ValidateSliceChange(rst any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateSliceChange", reflect.TypeOf((*MockPoolMgr)(nil).ValidateSliceChange), rst)
}
