package crew_service_tests

//
//var testCrewServiceDetachParticipantFromCrew = []struct {
//	testName  string
//	inputData struct {
//		crewID        uuid.UUID
//		participantID uuid.UUID
//	}
//	prepare     func(fields *crewServiceFields)
//	checkOutput func(t *testing.T, crew *models.Crew, err error)
//}{
//	{
//		testName: "Success",
//		inputData: struct {
//			crewID        uuid.UUID
//			participantID uuid.UUID
//		}{
//			uuid.New(),
//			uuid.New(),
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().DetachParticipantFromCrew(gomock.Any(), gomock.Any()).Return(nil)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.NoError(t, err)
//		},
//	},
//	{
//		testName: "detach participant from crew error",
//		inputData: struct {
//			crewID        uuid.UUID
//			participantID uuid.UUID
//		}{
//			uuid.New(),
//			uuid.New(),
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().DetachParticipantFromCrew(gomock.Any(), gomock.Any()).Return(repository_errors.UpdateError)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, repository_errors.UpdateError, err)
//		},
//	},
//}
//
//func TestCrewServiceDetachParticipantFromCrew(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	fields := initCrewServiceFields(ctrl)
//	crewService := initCrewService(fields)
//
//	for _, tt := range testCrewServiceDetachParticipantFromCrew {
//		t.Run(tt.testName, func(t *testing.T) {
//			tt.prepare(fields)
//			err := crewService.DetachParticipantFromCrew(tt.inputData.participantID, tt.inputData.crewID)
//			tt.checkOutput(t, nil, err)
//		})
//	}
//}
