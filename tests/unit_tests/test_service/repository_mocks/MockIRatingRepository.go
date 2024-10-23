package mock_repository_interfaces

import (
	models "PPO_BMSTU/internal/models"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

type MockIRatingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRatingRepositoryMockRecorder
}

// MockIRatingRepositoryMockRecorder is the mock recorder for MockIRatingRepository.
type MockIRatingRepositoryMockRecorder struct {
	mock *MockIRatingRepository
}

// NewMockIRatingRepository creates a new mock instance.
func NewMockIRatingRepository(ctrl *gomock.Controller) *MockIRatingRepository {
	mock := &MockIRatingRepository{ctrl: ctrl}
	mock.recorder = &MockIRatingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRatingRepository) EXPECT() *MockIRatingRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIRatingRepository) Create(rating *models.Rating) (*models.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", rating)
	ret0, _ := ret[0].(*models.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIRatingRepositoryMockRecorder) Create(rating any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIRatingRepository)(nil).Create), rating)
}

// Delete mocks base method.
func (m *MockIRatingRepository) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIRatingRepositoryMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIRatingRepository)(nil).Delete), id)
}

// GetRatingDataByID mocks base method.
func (m *MockIRatingRepository) GetRatingDataByID(id uuid.UUID) (*models.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRatingDataByID", id)
	ret0, _ := ret[0].(*models.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRatingDataByID indicates an expected call of GetRatingDataByID.
func (mr *MockIRatingRepositoryMockRecorder) GetRatingDataByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRatingDataByID", reflect.TypeOf((*MockIRatingRepository)(nil).GetRatingDataByID), id)
}

// AttachJudgeToRating mocks base method.
func (m *MockIRatingRepository) AttachJudgeToRating(ratingID, judgeID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachJudgeToRating", ratingID, judgeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttachJudgeToRating indicates an expected call of AttachJudgeToRating.
func (mr *MockIRatingRepositoryMockRecorder) AttachJudgeToRating(ratingID, judgeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachJudgeToRating", reflect.TypeOf((*MockIRatingRepository)(nil).AttachJudgeToRating), ratingID, judgeID)
}

// DetachJudgeFromRating mocks base method.
func (m *MockIRatingRepository) DetachJudgeFromRating(ratingID, judgeID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachJudgeFromRating", ratingID, judgeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachJudgeFromRating indicates an expected call of DetachJudgeFromRating.
func (mr *MockIRatingRepositoryMockRecorder) DetachJudgeFromRating(ratingID, judgeID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachJudgeFromRating", reflect.TypeOf((*MockIRatingRepository)(nil).DetachJudgeFromRating), ratingID, judgeID)
}

// Update mocks base method.
func (m *MockIRatingRepository) Update(rating *models.Rating) (*models.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", rating)
	ret0, _ := ret[0].(*models.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIRatingRepositoryMockRecorder) Update(rating any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRatingRepository)(nil).Update), rating)
}

// GetAllRatings mocks base method.
func (m *MockIRatingRepository) GetAllRatings() ([]models.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRatings")
	ret0, _ := ret[0].([]models.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRatings indicates an expected call of GetAllRatings.
func (mr *MockIRatingRepositoryMockRecorder) GetAllRatings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRatings", reflect.TypeOf((*MockIRatingRepository)(nil).GetAllRatings))
}

// GetRatingTable mocks base method.
func (m *MockIRatingRepository) GetRatingTable(id uuid.UUID) ([]models.RatingTableLine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRatingTable", id)
	ret0, _ := ret[0].([]models.RatingTableLine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRatingTable indicates an expected call of GetRatingTable.
func (mr *MockIRatingRepositoryMockRecorder) GetRatingTable(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRatingTable", reflect.TypeOf((*MockIRatingRepository)(nil).GetRatingTable), id)
}
