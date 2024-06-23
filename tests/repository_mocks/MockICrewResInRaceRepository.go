package mock_repository_interfaces

import (
	"PPO_BMSTU/internal/models"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"reflect"
)

// MockICrewResInRaceRepository is a mock of ICrewResInRaceRepository interface.
type MockICrewResInRaceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICrewResInRaceRepositoryMockRecorder
}

// MockICrewResInRaceRepositoryMockRecorder is the mock recorder for MockICrewResInRaceRepository.
type MockICrewResInRaceRepositoryMockRecorder struct {
	mock *MockICrewResInRaceRepository
}

// NewMockICrewResInRaceRepository creates a new mock instance.
func NewMockICrewResInRaceRepository(ctrl *gomock.Controller) *MockICrewResInRaceRepository {
	mock := &MockICrewResInRaceRepository{ctrl: ctrl}
	mock.recorder = &MockICrewResInRaceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICrewResInRaceRepository) EXPECT() *MockICrewResInRaceRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockICrewResInRaceRepository) Create(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", crewResInRace)
	ret0, _ := ret[0].(*models.CrewResInRace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockICrewResInRaceRepositoryMockRecorder) Create(crewResInRace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockICrewResInRaceRepository)(nil).Create), crewResInRace)
}

// Delete mocks base method.
func (m *MockICrewResInRaceRepository) Delete(raceID uuid.UUID, crewID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", raceID, crewID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockICrewResInRaceRepositoryMockRecorder) Delete(raceID uuid.UUID, crewID uuid.UUID) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockICrewResInRaceRepository)(nil).Delete), raceID, crewID)
}

// GetCrewResInRaceDataByID mocks base method.
func (m *MockICrewResInRaceRepository) GetCrewResInRaceDataByID(id uuid.UUID) (*models.CrewResInRace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCrewResInRaceDataByID", id)
	ret0, _ := ret[0].(*models.CrewResInRace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCrewResInRaceDataByID indicates an expected call of GetCrewResInRaceDataByID.
func (mr *MockICrewResInRaceRepositoryMockRecorder) GetCrewResInRaceDataByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCrewResInRaceDataByID", reflect.TypeOf((*MockICrewResInRaceRepository)(nil).GetCrewResInRaceDataByID), id)
}

// Update mocks base method.
func (m *MockICrewResInRaceRepository) Update(crewResInRace *models.CrewResInRace) (*models.CrewResInRace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", crewResInRace)
	ret0, _ := ret[0].(*models.CrewResInRace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockICrewResInRaceRepositoryMockRecorder) Update(crewResInRace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockICrewResInRaceRepository)(nil).Update), crewResInRace)
}

// GetCrewResByRaceIDAndCrewID mocks base method.
func (m *MockICrewResInRaceRepository) GetCrewResByRaceIDAndCrewID(raceID uuid.UUID, crewID uuid.UUID) (*models.CrewResInRace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCrewResByRaceIDAndCrewID", raceID, crewID)
	ret0, _ := ret[0].(*models.CrewResInRace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCrewResByRaceIDAndCrewID indicates an expected call of GetCrewResByRaceIDAndCrewID.
func (mr *MockICrewResInRaceRepositoryMockRecorder) GetCrewResByRaceIDAndCrewID(raceID, crewID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCrewResByRaceIDAndCrewID", reflect.TypeOf((*MockICrewResInRaceRepository)(nil).GetCrewResByRaceIDAndCrewID), raceID, crewID)
}

// GetAllCrewResInRace mocks base method.
func (m *MockICrewResInRaceRepository) GetAllCrewResInRace(raceID uuid.UUID) ([]models.CrewResInRace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCrewResInRace", raceID)
	ret0, _ := ret[0].([]models.CrewResInRace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCrewResInRace indicates an expected call of GetAllCrewResInRace.
func (mr *MockICrewResInRaceRepositoryMockRecorder) GetAllCrewResInRace(raceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCrewResInRace", reflect.TypeOf((*MockICrewResInRaceRepository)(nil).GetAllCrewResInRace), raceID)
}
