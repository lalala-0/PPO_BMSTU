package crew_service_tests

//var testCrewServiceUpdateCrewByID = []struct {
//	testName  string
//	inputData struct {
//		id       uuid.UUID
//		ratingID uuid.UUID
//		sailNum  int
//		class    int
//	}
//	prepare     func(fields *crewServiceFields)
//	checkOutput func(t *testing.T, crew *models.Crew, err error)
//}{
//	{
//		testName: "Success",
//		inputData: struct {
//			id       uuid.UUID
//			ratingID uuid.UUID
//			sailNum  int
//			class    int
//		}{
//			uuid.New(),
//			uuid.New(),
//			89,
//			models.Cadet,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(
//				builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), 189, models.LaserRadial), nil)
//			fields.crewRepoMock.EXPECT().Update(gomock.Any()).Return(
//				builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), 89, models.LaserRadial), nil)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.NoError(t, err)
//		},
//	},
//	{
//		testName: "crew not found",
//		inputData: struct {
//			id       uuid.UUID
//			ratingID uuid.UUID
//			sailNum  int
//			class    int
//		}{
//			uuid.New(),
//			uuid.New(),
//			89,
//			models.Laser,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(nil, repository_errors.DoesNotExist)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, fmt.Errorf("SERVICE: GetCrewByID method failed"), err)
//		},
//	},
//	{
//		testName: "invalid input",
//		inputData: struct {
//			id       uuid.UUID
//			ratingID uuid.UUID
//			sailNum  int
//			class    int
//		}{
//			uuid.New(),
//			uuid.New(),
//			-89,
//			90,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().GetCrewDataByID(gomock.Any()).Return(
//				builders.CrewMother.CustomCrew(uuid.New(), uuid.New(), 189, models.LaserRadial),
//				nil)
//
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input"), err)
//		},
//	},
//}
//
//func TestCrewServiceUpdateCrewByID(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	fields := initCrewServiceFields(ctrl)
//	crewService := initCrewService(fields)
//
//	for _, tt := range testCrewServiceUpdateCrewByID {
//		t.Run(tt.testName, func(t *testing.T) {
//			tt.prepare(fields)
//			crew, err := crewService.UpdateCrewByID(tt.inputData.id, tt.inputData.ratingID, tt.inputData.class, tt.inputData.sailNum)
//			tt.checkOutput(t, crew, err)
//		})
//	}
//}
//
