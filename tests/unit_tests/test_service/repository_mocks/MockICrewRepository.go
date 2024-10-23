package mock_repository_interfaces

import (
	models "PPO_BMSTU/internal/models"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

type MockICrewRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICrewRepositoryMockRecorder
}

// MockICrewRepositoryMockRecorder is the mock recorder for MockICrewRepository.
type MockICrewRepositoryMockRecorder struct {
	mock *MockICrewRepository
}

// NewMockICrewRepository creates a new mock instance.
func NewMockICrewRepository(ctrl *gomock.Controller) *MockICrewRepository {
	mock := &MockICrewRepository{ctrl: ctrl}
	mock.recorder = &MockICrewRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICrewRepository) EXPECT() *MockICrewRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockICrewRepository) Create(crew *models.Crew) (*models.Crew, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", crew)
	ret0, _ := ret[0].(*models.Crew)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockICrewRepositoryMockRecorder) Create(crew any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockICrewRepository)(nil).Create), crew)
}

// Delete mocks base method.
func (m *MockICrewRepository) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockICrewRepositoryMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockICrewRepository)(nil).Delete), id)
}

// GetCrewDataByID mocks base method.
func (m *MockICrewRepository) GetCrewDataByID(id uuid.UUID) (*models.Crew, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCrewDataByID", id)
	ret0, _ := ret[0].(*models.Crew)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCrewDataByID indicates an expected call of GetCrewDataByID.
func (mr *MockICrewRepositoryMockRecorder) GetCrewDataByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCrewDataByID", reflect.TypeOf((*MockICrewRepository)(nil).GetCrewDataByID), id)
}

// GetCrewDataBySailNumAndRatingID mocks base method.
func (m *MockICrewRepository) GetCrewDataBySailNumAndRatingID(sailNum int, ratingID uuid.UUID) (*models.Crew, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCrewDataBySailNumAndRatingID", sailNum, ratingID)
	ret0, _ := ret[0].(*models.Crew)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCrewDataBySailNumAndRatingID indicates an expected call of GetCrewDataBySailNumAndRatingID.
func (mr *MockICrewRepositoryMockRecorder) GetCrewDataBySailNumAndRatingID(sailNum any, ratingID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCrewDataBySailNumAndRatingID", reflect.TypeOf((*MockICrewRepository)(nil).GetCrewDataBySailNumAndRatingID), sailNum, ratingID)
}

// GetCrewsDataByRatingID mocks base method.
func (m *MockICrewRepository) GetCrewsDataByRatingID(id uuid.UUID) ([]models.Crew, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCrewsDataByRatingID", id)
	ret0, _ := ret[0].([]models.Crew)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCrewsDataByRatingID indicates an expected call of GetCrewsDataByRatingID.
func (mr *MockICrewRepositoryMockRecorder) GetCrewsDataByRatingID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCrewsDataByRatingID", reflect.TypeOf((*MockICrewRepository)(nil).GetCrewsDataByRatingID), id)
}

// GetCrewsDataByProtestID mocks base method.
func (m *MockICrewRepository) GetCrewsDataByProtestID(id uuid.UUID) ([]models.Crew, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCrewsDataByProtestID", id)
	ret0, _ := ret[0].([]models.Crew)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCrewsDataByProtestID indicates an expected call of GetCrewsDataByProtestID.
func (mr *MockICrewRepositoryMockRecorder) GetCrewsDataByProtestID(id uuid.UUID) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCrewsDataByProtestID", reflect.TypeOf((*MockICrewRepository)(nil).GetCrewsDataByProtestID), id)
}

// AttachParticipantToCrew mocks base method.
func (m *MockICrewRepository) AttachParticipantToCrew(crewID, participantID uuid.UUID, helmsman int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachParticipantToCrew", crewID, participantID, helmsman)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachParticipantToCrew indicates an expected call of AttachParticipantToCrew.
func (mr *MockICrewRepositoryMockRecorder) AttachParticipantToCrew(crewID, participantID, helmsman any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachParticipantToCrew", reflect.TypeOf((*MockICrewRepository)(nil).AttachParticipantToCrew), crewID, participantID, helmsman)
}

// DetachParticipantFromCrew mocks base method.
func (m *MockICrewRepository) DetachParticipantFromCrew(crewID, participantID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachParticipantFromCrew", crewID, participantID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachParticipantFromCrew indicates an expected call of DetachParticipantFromCrew.
func (mr *MockICrewRepositoryMockRecorder) DetachParticipantFromCrew(crewID, participantID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachParticipantFromCrew", reflect.TypeOf((*MockICrewRepository)(nil).DetachParticipantFromCrew), crewID, participantID)
}

// ReplaceParticipantStatusInCrew mocks base method.
func (m *MockICrewRepository) ReplaceParticipantStatusInCrew(crewID, participantID uuid.UUID, helmsman int, active int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplaceParticipantStatusInCrew", crewID, participantID, helmsman, active)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplaceParticipantStatusInCrew indicates an expected call of ReplaceParticipantStatusInCrew.
func (mr *MockICrewRepositoryMockRecorder) ReplaceParticipantStatusInCrew(crewID, participantID, helmsman, active any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceParticipantStatusInCrew", reflect.TypeOf((*MockICrewRepository)(nil).ReplaceParticipantStatusInCrew), crewID, participantID, helmsman, active)
}

// Update mocks base method.
func (m *MockICrewRepository) Update(crew *models.Crew) (*models.Crew, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", crew)
	ret0, _ := ret[0].(*models.Crew)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockICrewRepositoryMockRecorder) Update(crew any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockICrewRepository)(nil).Update), crew)
}
