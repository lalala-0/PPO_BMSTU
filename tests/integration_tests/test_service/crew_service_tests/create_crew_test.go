package crew_service_tests

import (
	"PPO_BMSTU/internal/models"
	"PPO_BMSTU/tests/unit_tests/builders"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

type inputDataType struct {
	id       uuid.UUID
	ratingID uuid.UUID
	sailNum  int
	class    int
}

var testCrewServiceAddNewCrew = []struct {
	testName    string
	inputData   inputDataType
	prepare     func(t *testing.T, suite *crewServiceTestSuite, data inputDataType)
	checkOutput func(t *testing.T, crew *models.Crew, err error)
}{
	{
		testName: "create crew success test",
		inputData: struct {
			id       uuid.UUID
			ratingID uuid.UUID
			sailNum  int
			class    int
		}{
			uuid.New(),
			uuid.New(),
			89,
			models.Laser,
		},
		prepare: func(t *testing.T, suite *crewServiceTestSuite, data inputDataType) {
			err := suite.initializer.ClearAll()
			assert.NoError(t, err)
			_, err = suite.initializer.CreateRating(builders.RatingMother.WithID(data.ratingID))
			assert.NoError(t, err)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.NoError(t, err)
			assert.NotNil(t, crew)
		},
	},
	{
		testName: "invalid input",
		inputData: struct {
			id       uuid.UUID
			ratingID uuid.UUID
			sailNum  int
			class    int
		}{
			uuid.New(),
			uuid.New(),
			-89,
			90,
		},
		prepare: func(t *testing.T, suite *crewServiceTestSuite, data inputDataType) {
			err := suite.initializer.ClearAll()
			assert.NoError(t, err)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Nil(t, crew)
			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input"), err)
		},
	},
	{
		testName: "crew creation error",
		inputData: struct {
			id       uuid.UUID
			ratingID uuid.UUID
			sailNum  int
			class    int
		}{
			uuid.New(),
			uuid.New(),
			89,
			models.Laser,
		},
		prepare: func(t *testing.T, suite *crewServiceTestSuite, data inputDataType) {
			err := suite.initializer.ClearAll()
			assert.NoError(t, err)
		},
		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
			assert.Error(t, err)
			assert.Nil(t, crew)
			assert.Equal(t, fmt.Errorf("SERVICE: Create method failed"), err)
		},
	},
}

func (suite *crewServiceTestSuite) TestCrewService_CreateCrew() {
	for _, tt := range testCrewServiceAddNewCrew {
		suite.T().Run(tt.testName, func(t *testing.T) {
			tt.prepare(t, suite, tt.inputData)
			crew, err := suite.service.AddNewCrew(tt.inputData.id, tt.inputData.ratingID, tt.inputData.class, tt.inputData.sailNum)
			tt.checkOutput(t, crew, err)
		})
	}
}
