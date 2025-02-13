package mock_postgres

import (
	entity "demorestapi/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDataProvider is a mock of DataProvider interface.
type MockDataProvider struct {
	ctrl     *gomock.Controller
	recorder *MockDataProviderMockRecorder
}

// MockDataProviderMockRecorder is the mock recorder for MockDataProvider.
type MockDataProviderMockRecorder struct {
	mock *MockDataProvider
}

// NewMockDataProvider creates a new mock instance.
func NewMockDataProvider(ctrl *gomock.Controller) *MockDataProvider {
	mock := &MockDataProvider{ctrl: ctrl}
	mock.recorder = &MockDataProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataProvider) EXPECT() *MockDataProviderMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockDataProvider) AddUser(u *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", u)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockDataProviderMockRecorder) AddUser(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockDataProvider)(nil).AddUser), u)
}

// GetUser mocks base method.
func (m *MockDataProvider) GetUser(id string) *entity.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(*entity.User)
	return ret0
}

// GetUser indicates an expected call of GetUser.
func (mr *MockDataProviderMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockDataProvider)(nil).GetUser), id)
}

// UpdateUser mocks base method.
func (m *MockDataProvider) UpdateUser(u *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", u)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockDataProviderMockRecorder) UpdateUser(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockDataProvider)(nil).UpdateUser), u)
}
