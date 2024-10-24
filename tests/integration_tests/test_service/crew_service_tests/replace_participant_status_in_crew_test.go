package crew_service_tests

//
//var testCrewServiceReplaceParticipantStatusInCrew = []struct {
//	testName  string
//	inputData struct {
//		participantID uuid.UUID
//		crewID        uuid.UUID
//		helmsman      int
//		active        int
//	}
//	prepare     func(fields *crewServiceFields)
//	checkOutput func(t *testing.T, crew *models.Crew, err error)
//}{
//	{
//		testName: "Success",
//		inputData: struct {
//			participantID uuid.UUID
//			crewID        uuid.UUID
//			helmsman      int
//			active        int
//		}{
//			uuid.New(),
//			uuid.New(),
//			1,
//			1,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().ReplaceParticipantStatusInCrew(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
//			//fields.crewRepoMock.EXPECT().ReplaceParticipantStatusInCrew(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
//
//			//fields.crewRepoMock.EXPECT().ReplaceParticipantStatusInCrew(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.NoError(t, err)
//		},
//	},
//	{
//		testName: "replace participant status in crew error",
//		inputData: struct {
//			participantID uuid.UUID
//			crewID        uuid.UUID
//			helmsman      int
//			active        int
//		}{
//			uuid.New(),
//			uuid.New(),
//			1,
//			1,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().ReplaceParticipantStatusInCrew(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(repository_errors.UpdateError)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, repository_errors.UpdateError, err)
//		},
//	},
//	{
//		testName: "input data error",
//		inputData: struct {
//			participantID uuid.UUID
//			crewID        uuid.UUID
//			helmsman      int
//			active        int
//		}{
//			uuid.New(),
//			uuid.New(),
//			1,
//			-1,
//		},
//		prepare: func(fields *crewServiceFields) {
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, fmt.Errorf("SERVICE: Invalid input"), err)
//		},
//	},
//}
//
//func TestCrewServiceReplaceParticipantStatusInCrew(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	fields := initCrewServiceFields(ctrl)
//	crewService := initCrewService(fields)
//
//	for _, tt := range testCrewServiceReplaceParticipantStatusInCrew {
//		t.Run(tt.testName, func(t *testing.T) {
//			tt.prepare(fields)
//			err := crewService.ReplaceParticipantStatusInCrew(tt.inputData.participantID, tt.inputData.crewID, tt.inputData.helmsman, tt.inputData.active)
//			tt.checkOutput(t, nil, err)
//		})
//	}
//}
//
