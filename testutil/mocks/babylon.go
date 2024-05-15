// Code generated by MockGen. DO NOT EDIT.
// Source: clientcontroller/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	math "cosmossdk.io/math"
	types "github.com/babylonchain/finality-provider/types"
	btcec "github.com/btcsuite/btcd/btcec/v2"
	gomock "github.com/golang/mock/gomock"
)

// MockClientController is a mock of ClientController interface.
type MockClientController struct {
	ctrl     *gomock.Controller
	recorder *MockClientControllerMockRecorder
}

// MockClientControllerMockRecorder is the mock recorder for MockClientController.
type MockClientControllerMockRecorder struct {
	mock *MockClientController
}

// NewMockClientController creates a new mock instance.
func NewMockClientController(ctrl *gomock.Controller) *MockClientController {
	mock := &MockClientController{ctrl: ctrl}
	mock.recorder = &MockClientControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientController) EXPECT() *MockClientControllerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClientController) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientControllerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClientController)(nil).Close))
}

// QueryActivatedHeight mocks base method.
func (m *MockClientController) QueryActivatedHeight() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryActivatedHeight")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryActivatedHeight indicates an expected call of QueryActivatedHeight.
func (mr *MockClientControllerMockRecorder) QueryActivatedHeight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryActivatedHeight", reflect.TypeOf((*MockClientController)(nil).QueryActivatedHeight))
}

// QueryBestBlock mocks base method.
func (m *MockClientController) QueryBestBlock() (*types.BlockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBestBlock")
	ret0, _ := ret[0].(*types.BlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryBestBlock indicates an expected call of QueryBestBlock.
func (mr *MockClientControllerMockRecorder) QueryBestBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBestBlock", reflect.TypeOf((*MockClientController)(nil).QueryBestBlock))
}

// QueryBlock mocks base method.
func (m *MockClientController) QueryBlock(height uint64) (*types.BlockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBlock", height)
	ret0, _ := ret[0].(*types.BlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryBlock indicates an expected call of QueryBlock.
func (mr *MockClientControllerMockRecorder) QueryBlock(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBlock", reflect.TypeOf((*MockClientController)(nil).QueryBlock), height)
}

// QueryBlocks mocks base method.
func (m *MockClientController) QueryBlocks(startHeight, endHeight, limit uint64) ([]*types.BlockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBlocks", startHeight, endHeight, limit)
	ret0, _ := ret[0].([]*types.BlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryBlocks indicates an expected call of QueryBlocks.
func (mr *MockClientControllerMockRecorder) QueryBlocks(startHeight, endHeight, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBlocks", reflect.TypeOf((*MockClientController)(nil).QueryBlocks), startHeight, endHeight, limit)
}

// QueryFinalityProviderSlashed mocks base method.
func (m *MockClientController) QueryFinalityProviderSlashed(fpPk *btcec.PublicKey) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFinalityProviderSlashed", fpPk)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFinalityProviderSlashed indicates an expected call of QueryFinalityProviderSlashed.
func (mr *MockClientControllerMockRecorder) QueryFinalityProviderSlashed(fpPk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFinalityProviderSlashed", reflect.TypeOf((*MockClientController)(nil).QueryFinalityProviderSlashed), fpPk)
}

// QueryFinalityProviderVotingPower mocks base method.
func (m *MockClientController) QueryFinalityProviderVotingPower(fpPk *btcec.PublicKey, blockHeight uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFinalityProviderVotingPower", fpPk, blockHeight)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFinalityProviderVotingPower indicates an expected call of QueryFinalityProviderVotingPower.
func (mr *MockClientControllerMockRecorder) QueryFinalityProviderVotingPower(fpPk, blockHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFinalityProviderVotingPower", reflect.TypeOf((*MockClientController)(nil).QueryFinalityProviderVotingPower), fpPk, blockHeight)
}

// QueryLastFinalizedEpoch mocks base method.
func (m *MockClientController) QueryLastFinalizedEpoch() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryLastFinalizedEpoch")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryLastFinalizedEpoch indicates an expected call of QueryLastFinalizedEpoch.
func (mr *MockClientControllerMockRecorder) QueryLastFinalizedEpoch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryLastFinalizedEpoch", reflect.TypeOf((*MockClientController)(nil).QueryLastFinalizedEpoch))
}

// QueryLatestFinalizedBlocks mocks base method.
func (m *MockClientController) QueryLatestFinalizedBlocks(count uint64) ([]*types.BlockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryLatestFinalizedBlocks", count)
	ret0, _ := ret[0].([]*types.BlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryLatestFinalizedBlocks indicates an expected call of QueryLatestFinalizedBlocks.
func (mr *MockClientControllerMockRecorder) QueryLatestFinalizedBlocks(count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryLatestFinalizedBlocks", reflect.TypeOf((*MockClientController)(nil).QueryLatestFinalizedBlocks), count)
}

// RegisterFinalityProvider mocks base method.
func (m *MockClientController) RegisterFinalityProvider(chainPk []byte, fpPk *btcec.PublicKey, pop []byte, commission *math.LegacyDec, description []byte, masterPubRand string) (*types.TxResponse, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterFinalityProvider", chainPk, fpPk, pop, commission, description, masterPubRand)
	ret0, _ := ret[0].(*types.TxResponse)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RegisterFinalityProvider indicates an expected call of RegisterFinalityProvider.
func (mr *MockClientControllerMockRecorder) RegisterFinalityProvider(chainPk, fpPk, pop, commission, description, masterPubRand interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterFinalityProvider", reflect.TypeOf((*MockClientController)(nil).RegisterFinalityProvider), chainPk, fpPk, pop, commission, description, masterPubRand)
}

// SubmitBatchFinalitySigs mocks base method.
func (m *MockClientController) SubmitBatchFinalitySigs(fpPk *btcec.PublicKey, blocks []*types.BlockInfo, sigs []*btcec.ModNScalar) (*types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitBatchFinalitySigs", fpPk, blocks, sigs)
	ret0, _ := ret[0].(*types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitBatchFinalitySigs indicates an expected call of SubmitBatchFinalitySigs.
func (mr *MockClientControllerMockRecorder) SubmitBatchFinalitySigs(fpPk, blocks, sigs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitBatchFinalitySigs", reflect.TypeOf((*MockClientController)(nil).SubmitBatchFinalitySigs), fpPk, blocks, sigs)
}

// SubmitFinalitySig mocks base method.
func (m *MockClientController) SubmitFinalitySig(fpPk *btcec.PublicKey, blockHeight uint64, blockHash []byte, sig *btcec.ModNScalar) (*types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitFinalitySig", fpPk, blockHeight, blockHash, sig)
	ret0, _ := ret[0].(*types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitFinalitySig indicates an expected call of SubmitFinalitySig.
func (mr *MockClientControllerMockRecorder) SubmitFinalitySig(fpPk, blockHeight, blockHash, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitFinalitySig", reflect.TypeOf((*MockClientController)(nil).SubmitFinalitySig), fpPk, blockHeight, blockHash, sig)
}
