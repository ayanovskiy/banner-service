// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/storage/storage.go

// Package mock is a generated GoMock package.
package mock

import (
	model "banner-service/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIStorage is a mock of IStorage interface.
type MockIStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIStorageMockRecorder
}

// MockIStorageMockRecorder is the mock recorder for MockIStorage.
type MockIStorageMockRecorder struct {
	mock *MockIStorage
}

// NewMockIStorage creates a new mock instance.
func NewMockIStorage(ctrl *gomock.Controller) *MockIStorage {
	mock := &MockIStorage{ctrl: ctrl}
	mock.recorder = &MockIStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStorage) EXPECT() *MockIStorageMockRecorder {
	return m.recorder
}

// AddBannerToSlot mocks base method.
func (m *MockIStorage) AddBannerToSlot(slotId, bannerId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBannerToSlot", slotId, bannerId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBannerToSlot indicates an expected call of AddBannerToSlot.
func (mr *MockIStorageMockRecorder) AddBannerToSlot(slotId, bannerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBannerToSlot", reflect.TypeOf((*MockIStorage)(nil).AddBannerToSlot), slotId, bannerId)
}

// CloseDb mocks base method.
func (m *MockIStorage) CloseDb() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseDb")
}

// CloseDb indicates an expected call of CloseDb.
func (mr *MockIStorageMockRecorder) CloseDb() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseDb", reflect.TypeOf((*MockIStorage)(nil).CloseDb))
}

// FindBanner mocks base method.
func (m *MockIStorage) FindBanner(id uint) (*model.Banner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBanner", id)
	ret0, _ := ret[0].(*model.Banner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBanner indicates an expected call of FindBanner.
func (mr *MockIStorageMockRecorder) FindBanner(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBanner", reflect.TypeOf((*MockIStorage)(nil).FindBanner), id)
}

// FindBannersBySlot mocks base method.
func (m *MockIStorage) FindBannersBySlot(slotId uint) (model.Banners, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBannersBySlot", slotId)
	ret0, _ := ret[0].(model.Banners)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBannersBySlot indicates an expected call of FindBannersBySlot.
func (mr *MockIStorageMockRecorder) FindBannersBySlot(slotId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBannersBySlot", reflect.TypeOf((*MockIStorage)(nil).FindBannersBySlot), slotId)
}

// FindGroup mocks base method.
func (m *MockIStorage) FindGroup(id uint) (*model.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindGroup", id)
	ret0, _ := ret[0].(*model.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindGroup indicates an expected call of FindGroup.
func (mr *MockIStorageMockRecorder) FindGroup(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindGroup", reflect.TypeOf((*MockIStorage)(nil).FindGroup), id)
}

// FindOrCreateStat mocks base method.
func (m *MockIStorage) FindOrCreateStat(slotId, bannerId, groupId uint) (*model.Stat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrCreateStat", slotId, bannerId, groupId)
	ret0, _ := ret[0].(*model.Stat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrCreateStat indicates an expected call of FindOrCreateStat.
func (mr *MockIStorageMockRecorder) FindOrCreateStat(slotId, bannerId, groupId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrCreateStat", reflect.TypeOf((*MockIStorage)(nil).FindOrCreateStat), slotId, bannerId, groupId)
}

// FindSlot mocks base method.
func (m *MockIStorage) FindSlot(id uint) (*model.Slot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSlot", id)
	ret0, _ := ret[0].(*model.Slot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindSlot indicates an expected call of FindSlot.
func (mr *MockIStorageMockRecorder) FindSlot(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSlot", reflect.TypeOf((*MockIStorage)(nil).FindSlot), id)
}

// FindStats mocks base method.
func (m *MockIStorage) FindStats() (model.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindStats")
	ret0, _ := ret[0].(model.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindStats indicates an expected call of FindStats.
func (mr *MockIStorageMockRecorder) FindStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindStats", reflect.TypeOf((*MockIStorage)(nil).FindStats))
}

// RemoveBannerFromSlot mocks base method.
func (m *MockIStorage) RemoveBannerFromSlot(slotId, bannerId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBannerFromSlot", slotId, bannerId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveBannerFromSlot indicates an expected call of RemoveBannerFromSlot.
func (mr *MockIStorageMockRecorder) RemoveBannerFromSlot(slotId, bannerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBannerFromSlot", reflect.TypeOf((*MockIStorage)(nil).RemoveBannerFromSlot), slotId, bannerId)
}

// UpdateStat mocks base method.
func (m *MockIStorage) UpdateStat(statId, shows, hits uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStat", statId, shows, hits)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStat indicates an expected call of UpdateStat.
func (mr *MockIStorageMockRecorder) UpdateStat(statId, shows, hits interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStat", reflect.TypeOf((*MockIStorage)(nil).UpdateStat), statId, shows, hits)
}
