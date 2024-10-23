package mock_repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"reflect"
)

// MockIRaceRepository is a mock of IRaceRepository interface.
type MockIRaceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRaceRepositoryMockRecorder
}

// MockIRaceRepositoryMockRecorder is the mock recorder for MockIRaceRepository.
type MockIRaceRepositoryMockRecorder struct {
	mock *MockIRaceRepository
}

// NewMockIRaceRepository creates a new mock instance.
func NewMockIRaceRepository(ctrl *gomock.Controller) *MockIRaceRepository {
	mock := &MockIRaceRepository{ctrl: ctrl}
	mock.recorder = &MockIRaceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRaceRepository) EXPECT() *MockIRaceRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIRaceRepository) Create(race *models.Race) (*models.Race, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", race)
	ret0, _ := ret[0].(*models.Race)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIRaceRepositoryMockRecorder) Create(race any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIRaceRepository)(nil).Create), race)
}

// Delete mocks base method.
func (m *MockIRaceRepository) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIRaceRepositoryMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIRaceRepository)(nil).Delete), id)
}

// GetRaceDataByID mocks base method.
func (m *MockIRaceRepository) GetRaceDataByID(id uuid.UUID) (*models.Race, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRaceDataByID", id)
	ret0, _ := ret[0].(*models.Race)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRaceDataByID indicates an expected call of GetRaceDataByID.
func (mr *MockIRaceRepositoryMockRecorder) GetRaceDataByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRaceDataByID", reflect.TypeOf((*MockIRaceRepository)(nil).GetRaceDataByID), id)
}

// GetRacesDataByRatingID mocks base method.
func (m *MockIRaceRepository) GetRacesDataByRatingID(ratingID uuid.UUID) ([]models.Race, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRacesDataByRatingID", ratingID)
	ret0, _ := ret[0].([]models.Race)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRacesDataByRatingID indicates an expected call of GetRacesDataByRatingID.
func (mr *MockIRaceRepositoryMockRecorder) GetRacesDataByRatingID(ratingID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRacesDataByRatingID", reflect.TypeOf((*MockIRaceRepository)(nil).GetRacesDataByRatingID), ratingID)
}

// Update mocks base method.
func (m *MockIRaceRepository) Update(race *models.Race) (*models.Race, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", race)
	ret0, _ := ret[0].(*models.Race)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIRaceRepositoryMockRecorder) Update(race any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRaceRepository)(nil).Update), race)
}
