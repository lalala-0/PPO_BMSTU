package mock_repository_interfaces

import (
	models "PPO_BMSTU/internal/models"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockIParticipantRepository is a mock of IParticipantRepository interface.
type MockIParticipantRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIParticipantRepositoryMockRecorder
}

// MockIParticipantRepositoryMockRecorder is the mock recorder for MockIParticipantRepository.
type MockIParticipantRepositoryMockRecorder struct {
	mock *MockIParticipantRepository
}

// NewMockIParticipantRepository creates a new mock instance.
func NewMockIParticipantRepository(ctrl *gomock.Controller) *MockIParticipantRepository {
	mock := &MockIParticipantRepository{ctrl: ctrl}
	mock.recorder = &MockIParticipantRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIParticipantRepository) EXPECT() *MockIParticipantRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIParticipantRepository) Create(participant *models.Participant) (*models.Participant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", participant)
	ret0, _ := ret[0].(*models.Participant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIParticipantRepositoryMockRecorder) Create(participant any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIParticipantRepository)(nil).Create), participant)
}

// Delete mocks base method.
func (m *MockIParticipantRepository) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIParticipantRepositoryMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIParticipantRepository)(nil).Delete), id)
}

// GetParticipantsDataByCrewID mocks base method.
func (m *MockIParticipantRepository) GetParticipantsDataByCrewID(id uuid.UUID) ([]models.Participant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParticipantsDataByCrewID", id)
	ret0, _ := ret[0].([]models.Participant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParticipantsDataByCrewID indicates an expected call of GetParticipantsDataByCrewID.
func (mr *MockIParticipantRepositoryMockRecorder) GetParticipantsDataByCrewID(id uuid.UUID) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParticipantsDataByCrewID", reflect.TypeOf((*MockIParticipantRepository)(nil).GetParticipantsDataByCrewID), id)
}

// GetAllParticipants mocks base method.
func (m *MockIParticipantRepository) GetAllParticipants() ([]models.Participant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllParticipants")
	ret0, _ := ret[0].([]models.Participant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllParticipants indicates an expected call of GetAllParticipants.
func (mr *MockIParticipantRepositoryMockRecorder) GetAllParticipants(id uuid.UUID) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllParticipants", reflect.TypeOf((*MockIParticipantRepository)(nil).GetAllParticipants))
}

// GetParticipantDataByID mocks base method.
func (m *MockIParticipantRepository) GetParticipantDataByID(id uuid.UUID) (*models.Participant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParticipantDataByID", id)
	ret0, _ := ret[0].(*models.Participant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParticipantDataByID indicates an expected call of GetParticipantDataByID.
func (mr *MockIParticipantRepositoryMockRecorder) GetParticipantDataByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParticipantDataByID", reflect.TypeOf((*MockIParticipantRepository)(nil).GetParticipantDataByID), id)
}

// Update mocks base method.
func (m *MockIParticipantRepository) Update(user *models.Participant) (*models.Participant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(*models.Participant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIParticipantRepositoryMockRecorder) Update(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIParticipantRepository)(nil).Update), user)
}

// GetParticipantsDataByProtestID mocks base method.
func (m *MockIParticipantRepository) GetParticipantsDataByProtestID(id uuid.UUID) ([]models.Participant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParticipantsDataByProtestID", id)
	ret0, _ := ret[0].([]models.Participant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParticipantsDataByProtestID indicates an expected call of GetParticipantsDataByProtestID.
func (mr *MockIParticipantRepositoryMockRecorder) GetParticipantsDataByProtestID(id uuid.UUID) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParticipantsDataByProtestID", reflect.TypeOf((*MockIParticipantRepository)(nil).GetParticipantsDataByCrewID), id)
}
