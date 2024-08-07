// Code generated by MockGen. DO NOT EDIT.
// Source: x/btcstaking/types/expected_keepers.go

// Package types is a generated GoMock package.
package types

import (
	context "context"
	big "math/big"
	reflect "reflect"

	types "github.com/babylonlabs-io/babylon/types"
	types0 "github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
	types1 "github.com/babylonlabs-io/babylon/x/btclightclient/types"
	types2 "github.com/babylonlabs-io/babylon/x/epoching/types"
	gomock "github.com/golang/mock/gomock"
)

// MockBTCLightClientKeeper is a mock of BTCLightClientKeeper interface.
type MockBTCLightClientKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBTCLightClientKeeperMockRecorder
}

// MockBTCLightClientKeeperMockRecorder is the mock recorder for MockBTCLightClientKeeper.
type MockBTCLightClientKeeperMockRecorder struct {
	mock *MockBTCLightClientKeeper
}

// NewMockBTCLightClientKeeper creates a new mock instance.
func NewMockBTCLightClientKeeper(ctrl *gomock.Controller) *MockBTCLightClientKeeper {
	mock := &MockBTCLightClientKeeper{ctrl: ctrl}
	mock.recorder = &MockBTCLightClientKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBTCLightClientKeeper) EXPECT() *MockBTCLightClientKeeperMockRecorder {
	return m.recorder
}

// GetBaseBTCHeader mocks base method.
func (m *MockBTCLightClientKeeper) GetBaseBTCHeader(ctx context.Context) *types1.BTCHeaderInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBaseBTCHeader", ctx)
	ret0, _ := ret[0].(*types1.BTCHeaderInfo)
	return ret0
}

// GetBaseBTCHeader indicates an expected call of GetBaseBTCHeader.
func (mr *MockBTCLightClientKeeperMockRecorder) GetBaseBTCHeader(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBaseBTCHeader", reflect.TypeOf((*MockBTCLightClientKeeper)(nil).GetBaseBTCHeader), ctx)
}

// GetHeaderByHash mocks base method.
func (m *MockBTCLightClientKeeper) GetHeaderByHash(ctx context.Context, hash *types.BTCHeaderHashBytes) *types1.BTCHeaderInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaderByHash", ctx, hash)
	ret0, _ := ret[0].(*types1.BTCHeaderInfo)
	return ret0
}

// GetHeaderByHash indicates an expected call of GetHeaderByHash.
func (mr *MockBTCLightClientKeeperMockRecorder) GetHeaderByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByHash", reflect.TypeOf((*MockBTCLightClientKeeper)(nil).GetHeaderByHash), ctx, hash)
}

// GetTipInfo mocks base method.
func (m *MockBTCLightClientKeeper) GetTipInfo(ctx context.Context) *types1.BTCHeaderInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTipInfo", ctx)
	ret0, _ := ret[0].(*types1.BTCHeaderInfo)
	return ret0
}

// GetTipInfo indicates an expected call of GetTipInfo.
func (mr *MockBTCLightClientKeeperMockRecorder) GetTipInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTipInfo", reflect.TypeOf((*MockBTCLightClientKeeper)(nil).GetTipInfo), ctx)
}

// MockBtcCheckpointKeeper is a mock of BtcCheckpointKeeper interface.
type MockBtcCheckpointKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBtcCheckpointKeeperMockRecorder
}

// MockBtcCheckpointKeeperMockRecorder is the mock recorder for MockBtcCheckpointKeeper.
type MockBtcCheckpointKeeperMockRecorder struct {
	mock *MockBtcCheckpointKeeper
}

// NewMockBtcCheckpointKeeper creates a new mock instance.
func NewMockBtcCheckpointKeeper(ctrl *gomock.Controller) *MockBtcCheckpointKeeper {
	mock := &MockBtcCheckpointKeeper{ctrl: ctrl}
	mock.recorder = &MockBtcCheckpointKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBtcCheckpointKeeper) EXPECT() *MockBtcCheckpointKeeperMockRecorder {
	return m.recorder
}

// GetParams mocks base method.
func (m *MockBtcCheckpointKeeper) GetParams(ctx context.Context) types0.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types0.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockBtcCheckpointKeeperMockRecorder) GetParams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockBtcCheckpointKeeper)(nil).GetParams), ctx)
}

// GetPowLimit mocks base method.
func (m *MockBtcCheckpointKeeper) GetPowLimit() *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPowLimit")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GetPowLimit indicates an expected call of GetPowLimit.
func (mr *MockBtcCheckpointKeeperMockRecorder) GetPowLimit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPowLimit", reflect.TypeOf((*MockBtcCheckpointKeeper)(nil).GetPowLimit))
}

// MockCheckpointingKeeper is a mock of CheckpointingKeeper interface.
type MockCheckpointingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockCheckpointingKeeperMockRecorder
}

// MockCheckpointingKeeperMockRecorder is the mock recorder for MockCheckpointingKeeper.
type MockCheckpointingKeeperMockRecorder struct {
	mock *MockCheckpointingKeeper
}

// NewMockCheckpointingKeeper creates a new mock instance.
func NewMockCheckpointingKeeper(ctrl *gomock.Controller) *MockCheckpointingKeeper {
	mock := &MockCheckpointingKeeper{ctrl: ctrl}
	mock.recorder = &MockCheckpointingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckpointingKeeper) EXPECT() *MockCheckpointingKeeperMockRecorder {
	return m.recorder
}

// GetEpoch mocks base method.
func (m *MockCheckpointingKeeper) GetEpoch(ctx context.Context) *types2.Epoch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEpoch", ctx)
	ret0, _ := ret[0].(*types2.Epoch)
	return ret0
}

// GetEpoch indicates an expected call of GetEpoch.
func (mr *MockCheckpointingKeeperMockRecorder) GetEpoch(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpoch", reflect.TypeOf((*MockCheckpointingKeeper)(nil).GetEpoch), ctx)
}

// GetLastFinalizedEpoch mocks base method.
func (m *MockCheckpointingKeeper) GetLastFinalizedEpoch(ctx context.Context) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastFinalizedEpoch", ctx)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetLastFinalizedEpoch indicates an expected call of GetLastFinalizedEpoch.
func (mr *MockCheckpointingKeeperMockRecorder) GetLastFinalizedEpoch(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastFinalizedEpoch", reflect.TypeOf((*MockCheckpointingKeeper)(nil).GetLastFinalizedEpoch), ctx)
}

// MockBtcStakingHooks is a mock of BtcStakingHooks interface.
type MockBtcStakingHooks struct {
	ctrl     *gomock.Controller
	recorder *MockBtcStakingHooksMockRecorder
}

// MockBtcStakingHooksMockRecorder is the mock recorder for MockBtcStakingHooks.
type MockBtcStakingHooksMockRecorder struct {
	mock *MockBtcStakingHooks
}

// NewMockBtcStakingHooks creates a new mock instance.
func NewMockBtcStakingHooks(ctrl *gomock.Controller) *MockBtcStakingHooks {
	mock := &MockBtcStakingHooks{ctrl: ctrl}
	mock.recorder = &MockBtcStakingHooksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBtcStakingHooks) EXPECT() *MockBtcStakingHooksMockRecorder {
	return m.recorder
}

// AfterFinalityProviderActivated mocks base method.
func (m *MockBtcStakingHooks) AfterFinalityProviderActivated(ctx context.Context, fpPk *types.BIP340PubKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterFinalityProviderActivated", ctx, fpPk)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterFinalityProviderActivated indicates an expected call of AfterFinalityProviderActivated.
func (mr *MockBtcStakingHooksMockRecorder) AfterFinalityProviderActivated(ctx, fpPk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterFinalityProviderActivated", reflect.TypeOf((*MockBtcStakingHooks)(nil).AfterFinalityProviderActivated), ctx, fpPk)
}