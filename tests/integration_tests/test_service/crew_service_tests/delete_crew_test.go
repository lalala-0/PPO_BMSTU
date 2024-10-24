package crew_service_tests

//
//var testCrewServiceDelete = []struct {
//	testName  string
//	inputData struct {
//		crewID uuid.UUID
//	}
//	prepare     func(fields *crewServiceFields)
//	checkOutput func(t *testing.T, err error)
//}{
//	{
//		testName: "delete crew success test",
//		inputData: struct {
//			crewID uuid.UUID
//		}{
//			uuid.New(),
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(builders.CrewMother.Default(), nil)
//			fields.crewRepoMock.EXPECT().Delete(gomock.Any()).Return(nil)
//		},
//		checkOutput: func(t *testing.T, err error) {
//			assert.NoError(t, err)
//		},
//	},
//	{
//		testName: "crew not found",
//		inputData: struct {
//			crewID uuid.UUID
//		}{
//			uuid.New(),
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
//		},
//		checkOutput: func(t *testing.T, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, fmt.Errorf("SERVICE: GetCrewDataByID method failed"), err)
//		},
//	},
//	{
//		testName: "delete crew error",
//		inputData: struct {
//			crewID uuid.UUID
//		}{
//			uuid.New(),
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(builders.CrewMother.Default(), nil)
//			fields.crewRepoMock.EXPECT().Delete(gomock.Any()).Return(repository_errors.DeleteError)
//		},
//		checkOutput: func(t *testing.T, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, fmt.Errorf("SERVICE: Delete method failed"), err)
//		},
//	},
//}
//
//func TestCrewServiceDeleteCrew(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	fields := initCrewServiceFields(ctrl)
//	crewService := initCrewService(fields)
//
//	for _, tt := range testCrewServiceDelete {
//		t.Run(tt.testName, func(t *testing.T) {
//			tt.prepare(fields)
//			err := crewService.DeleteCrewByID(tt.inputData.crewID)
//			tt.checkOutput(t, err)
//		})
//	}
//}
//
