package mock_repository_interfaces

import (
	models "PPO_BMSTU/internal/models"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockIProtestRepository is a mock of IProtestRepository interface.
type MockIProtestRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIProtestRepositoryMockRecorder
}

// MockIProtestRepositoryMockRecorder is the mock recorder for MockIProtestRepository.
type MockIProtestRepositoryMockRecorder struct {
	mock *MockIProtestRepository
}

// NewMockIProtestRepository creates a new mock instance.
func NewMockIProtestRepository(ctrl *gomock.Controller) *MockIProtestRepository {
	mock := &MockIProtestRepository{ctrl: ctrl}
	mock.recorder = &MockIProtestRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProtestRepository) EXPECT() *MockIProtestRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIProtestRepository) Create(protest *models.Protest) (*models.Protest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", protest)
	ret0, _ := ret[0].(*models.Protest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIProtestRepositoryMockRecorder) Create(protest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIProtestRepository)(nil).Create), protest)
}

// Delete mocks base method.
func (m *MockIProtestRepository) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIProtestRepositoryMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIProtestRepository)(nil).Delete), id)
}

// GetProtestDataByID mocks base method.
func (m *MockIProtestRepository) GetProtestDataByID(id uuid.UUID) (*models.Protest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProtestDataByID", id)
	ret0, _ := ret[0].(*models.Protest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProtestDataByID indicates an expected call of GetProtestDataByID.
func (mr *MockIProtestRepositoryMockRecorder) GetProtestDataByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProtestDataByID", reflect.TypeOf((*MockIProtestRepository)(nil).GetProtestDataByID), id)
}

// GetProtestsDataByRaceID mocks base method.
func (m *MockIProtestRepository) GetProtestsDataByRaceID(raceID uuid.UUID) ([]models.Protest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProtestsDataByRaceID", raceID)
	ret0, _ := ret[0].([]models.Protest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProtestsDataByRaceID indicates an expected call of GetProtestsDataByRaceID.
func (mr *MockIProtestRepositoryMockRecorder) GetProtestsDataByRaceID(raceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProtestsDataByRaceID", reflect.TypeOf((*MockIProtestRepository)(nil).GetProtestsDataByRaceID), raceID)
}

// GetProtestParticipantsIDByID mocks base method.
func (m *MockIProtestRepository) GetProtestParticipantsIDByID(id uuid.UUID) (map[int]uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProtestParticipantsIDByID", id)
	ret0, _ := ret[0].(map[int]uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProtestParticipantsIDByID indicates an expected call of GetProtestParticipantsIDByID.
func (mr *MockIProtestRepositoryMockRecorder) GetProtestParticipantsIDByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProtestParticipantsIDByID", reflect.TypeOf((*MockIProtestRepository)(nil).GetProtestParticipantsIDByID), id)
}

// AttachCrewToProtest mocks base method.
func (m *MockIProtestRepository) AttachCrewToProtest(protestID, crewID uuid.UUID, crewStatus int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachCrewToProtest", protestID, crewID, crewStatus)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachCrewToProtest indicates an expected call of AttachCrewToProtest.
func (mr *MockIProtestRepositoryMockRecorder) AttachCrewToProtest(protestID, crewID, crewStatus any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachCrewToProtest", reflect.TypeOf((*MockIProtestRepository)(nil).AttachCrewToProtest), protestID, crewID, crewStatus)
}

// DetachCrewFromProtest mocks base method.
func (m *MockIProtestRepository) DetachCrewFromProtest(protestID, crewID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachCrewFromProtest", protestID, crewID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachCrewFromProtest indicates an expected call of DetachCrewFromProtest.
func (mr *MockIProtestRepositoryMockRecorder) DetachCrewFromProtest(protestID, crewID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachCrewFromProtest", reflect.TypeOf((*MockIProtestRepository)(nil).DetachCrewFromProtest), protestID, crewID)
}

// Update mocks base method.
func (m *MockIProtestRepository) Update(protest *models.Protest) (*models.Protest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", protest)
	ret0, _ := ret[0].(*models.Protest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIProtestRepositoryMockRecorder) Update(protest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIProtestRepository)(nil).Update), protest)
}
