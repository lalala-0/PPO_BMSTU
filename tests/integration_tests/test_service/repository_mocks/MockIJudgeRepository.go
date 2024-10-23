package mock_repository_interfaces

import (
	models "PPO_BMSTU/internal/models"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"reflect"
)

// MockIJudgeRepository is a mock of IJudgeRepository interface.
type MockIJudgeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIJudgeRepositoryMockRecorder
}

// MockIJudgeRepositoryMockRecorder is the mock recorder for MockIJudgeRepository.
type MockIJudgeRepositoryMockRecorder struct {
	mock *MockIJudgeRepository
}

// NewMockIJudgeRepository creates a new mock instance.
func NewMockIJudgeRepository(ctrl *gomock.Controller) *MockIJudgeRepository {
	mock := &MockIJudgeRepository{ctrl: ctrl}
	mock.recorder = &MockIJudgeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIJudgeRepository) EXPECT() *MockIJudgeRepositoryMockRecorder {
	return m.recorder
}

// CreateProfile mocks base method.
func (m *MockIJudgeRepository) CreateProfile(judge *models.Judge) (*models.Judge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfile", judge)
	ret0, _ := ret[0].(*models.Judge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProfile indicates an expected call of CreateProfile.
func (mr *MockIJudgeRepositoryMockRecorder) CreateProfile(judge any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfile", reflect.TypeOf((*MockIJudgeRepository)(nil).CreateProfile), judge)
}

// DeleteProfile mocks base method.
func (m *MockIJudgeRepository) DeleteProfile(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProfile", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProfile indicates an expected call of DeleteProfile.
func (mr *MockIJudgeRepositoryMockRecorder) DeleteProfile(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfile", reflect.TypeOf((*MockIJudgeRepository)(nil).DeleteProfile), id)
}

// GetAllJudges mocks base method.
func (m *MockIJudgeRepository) GetAllJudges() ([]models.Judge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllJudges")
	ret0, _ := ret[0].([]models.Judge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllJudges indicates an expected call of GetAllJudges.
func (mr *MockIJudgeRepositoryMockRecorder) GetAllJudges() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllJudges", reflect.TypeOf((*MockIJudgeRepository)(nil).GetAllJudges))
}

// GetJudgeDataByID mocks base method.
func (m *MockIJudgeRepository) GetJudgeDataByID(id uuid.UUID) (*models.Judge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJudgeDataByID", id)
	ret0, _ := ret[0].(*models.Judge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJudgeDataByID indicates an expected call of GetJudgeDataByID.
func (mr *MockIJudgeRepositoryMockRecorder) GetJudgeDataByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJudgeDataByID", reflect.TypeOf((*MockIJudgeRepository)(nil).GetJudgeDataByID), id)
}

// GetJudgesDataByRatingID mocks base method.
func (m *MockIJudgeRepository) GetJudgesDataByRatingID(id uuid.UUID) ([]models.Judge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJudgesDataByRatingID", id)
	ret0, _ := ret[0].([]models.Judge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJudgesDataByRatingID indicates an expected call of GetJudgesDataByRatingID.
func (mr *MockIJudgeRepositoryMockRecorder) GetJudgesDataByRatingID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJudgesDataByRatingID", reflect.TypeOf((*MockIJudgeRepository)(nil).GetJudgesDataByRatingID), id)
}

// GetJudgeDataByProtestID mocks base method.
func (m *MockIJudgeRepository) GetJudgeDataByProtestID(id uuid.UUID) (*models.Judge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJudgeDataByProtestID", id)
	ret0, _ := ret[0].(*models.Judge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJudgeDataByProtestID indicates an expected call of GetJudgeDataByProtestID.
func (mr *MockIJudgeRepositoryMockRecorder) GetJudgeDataByProtestID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJudgeDataByProtestID", reflect.TypeOf((*MockIJudgeRepository)(nil).GetJudgeDataByProtestID), id)
}

// GetJudgeDataByLogin mocks base method.
func (m *MockIJudgeRepository) GetJudgeDataByLogin(login string) (*models.Judge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJudgeDataByLogin", login)
	ret0, _ := ret[0].(*models.Judge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJudgeDataByLogin indicates an expected call of GetJudgeDataByLogin.
func (mr *MockIJudgeRepositoryMockRecorder) GetJudgeDataByLogin(login any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJudgeDataByLogin", reflect.TypeOf((*MockIJudgeRepository)(nil).GetJudgeDataByLogin), login)
}

// UpdateProfile mocks base method.
func (m *MockIJudgeRepository) UpdateProfile(judge *models.Judge) (*models.Judge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", judge)
	ret0, _ := ret[0].(*models.Judge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockIJudgeRepositoryMockRecorder) UpdateProfile(judge any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockIJudgeRepository)(nil).UpdateProfile), judge)
}
