package crew_service_tests

//var testCrewServiceGetCrewDataByID = []struct {
//	testName    string
//	inputData   struct{ crewID uuid.UUID }
//	prepare     func(fields *crewServiceFields)
//	checkOutput func(t *testing.T, crew *models.Crew, err error)
//}{
//	{
//		testName:  "get crew by id success test",
//		inputData: struct{ crewID uuid.UUID }{uuid.New()},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(builders.CrewMother.Default(), nil)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.NoError(t, err)
//			assert.NotNil(t, crew)
//		},
//	},
//	{
//		testName:  "crew not found",
//		inputData: struct{ crewID uuid.UUID }{uuid.New()},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Nil(t, crew)
//			assert.Equal(t, repository_errors.DoesNotExist, err)
//		},
//	},
//	{
//		testName:  "get crew by id error",
//		inputData: struct{ crewID uuid.UUID }{uuid.New()},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(nil, repository_errors.SelectError)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Nil(t, crew)
//			assert.Equal(t, repository_errors.SelectError, err)
//		},
//	},
//}
//
//func TestCrewServiceGetCrewDataByID(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	fields := initCrewServiceFields(ctrl)
//	crewService := initCrewService(fields)
//
//	for _, tt := range testCrewServiceGetCrewDataByID {
//		t.Run(tt.testName, func(t *testing.T) {
//			tt.prepare(fields)
//			crew, err := crewService.GetCrewDataByID(tt.inputData.crewID)
//			tt.checkOutput(t, crew, err)
//		})
//	}
//}
